package hybrids_generator

import (
	"text/template"
	"io"
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
	"github.com/nebtex/omnibuff/pkg/utils"
	"go.uber.org/zap"
	"fmt"
	"strings"
)

type TableReaderGenerator struct {
	table   corev1.TableReader
	zap     *zap.Logger
	funcMap map[string]interface{}
}

func NewTableReaderGenerator(table corev1.TableReader, logger *zap.Logger) *TableReaderGenerator {
	table = table
	zap := logger.With(zap.String("TableName", table.Meta().Name()),
		zap.String("Type", "Reader implementation"),
		zap.String("Application", table.Meta().Application()),
	)

	t := &TableReaderGenerator{table: table, zap: zap}
	t.funcMap = map[string]interface{}{
		"TableName":  utils.TableName,
		"ShortName":  t.ShortName,
		"Capitalize": strings.Title,
		"FieldDoc": func(f corev1.FieldReader) (value string) {
			if f.Documentation() != nil {
				return f.Documentation().Short()
			}
			return
		}}
	return t
}

func ScalarFieldReaderAccessor(freader corev1.FieldReader, wr io.Writer) (err error) {

	tmpl, err := template.New("ScalarFieldReaderAccessor").Parse(`
func (t *TableSpec) {{.Field.Name}}() (value {{ToNativeType .Field.Type}}, ok bool) {
    value, ok := t.table.{{.Field.Type}}()
    if !ok{
		value = {{.Field.Default}}f
    }
}`)
	if err != nil {
		return
	}
	err = tmpl.Execute(wr, nil)
	return
}

func (t *TableReaderGenerator) ShortName() string {
	return strings.ToLower(string(t.table.Meta().Name()[0]))
}

func (t *TableReaderGenerator) Table() corev1.TableReader {
	return t.table
}

func (t *TableReaderGenerator) Generate(wr io.Writer) (err error) {
	var i int32
	var field corev1.FieldReader

	//add imports
	wr.Write([]byte(`
import "github.com/nebtex/hybrids/golang/hybrids"

`))

	// create struct
	tmpl, err := template.New("TableReaderGenerator").Funcs(t.funcMap).Parse(`
type {{TableName .}}Reader struct {
	_table hybrids.TableReader
}`)
	if err != nil {
		t.zap.Error(err.Error())
		return
	}

	err = tmpl.Execute(wr, t.Table())

	if err != nil {
		t.zap.Error(err.Error())
		return
	}
	fmt.Println(t.table.Fields().Len())
	//create Accessors
	for i = 0; i < t.table.Fields().Len(); i++ {
		field, err = t.table.Fields().Get(i)
		if err != nil {
			t.zap.Error(err.Error())
			return
		}
		switch field.Type() {
		case "String":
			err = t.StringAccessor(field, uint16(i), wr)
			if err != nil {
				return err
			}
		}
	}
	t.zap.Info(fmt.Sprintf("Table %s Created successfully", utils.TableName(t.table)))
	return
}
//Todo:
//default
//documentation
//table reader
//vector table reader
//resource
func (t *TableReaderGenerator) StringAccessor(freader corev1.FieldReader, fn uint16, wr io.Writer) (err error) {
	tmpl, err := template.New("StringAccessor").
		Funcs(t.funcMap).Parse(`
{{FieldDoc .Field}}
func ({{ShortName}} *{{TableName .Table}}Reader) {{Capitalize .Field.Name}}() (value string) {
	value, _ = {{ShortName}}._table.String({{.FieldNumber}})
	return
}
`)
	if err != nil {
		t.zap.Error("StringAccessor: template definition error", zap.String("err", err.Error()))
		return
	}

	err = tmpl.Execute(wr, map[string]interface{}{"Table": t.table, "Field": freader, "FieldNumber": fn})
	if err != nil {
		t.zap.Error("StringAccessor: template render error", zap.String("err", err.Error()))
		return
	}

	return

}

func (t *TableReaderGenerator) TableAccessor(freader corev1.FieldReader, fn uint16, wr io.Writer) (err error) {
	tmpl, err := template.New("StringAccessor").
		Funcs(t.funcMap).Parse(`
{{FieldDoc .Field}}
func ({{ShortName}} *{{TableName .Table}}Reader) {{Capitalize .Field.Name}}() {{TableName .Table}}Reader {
	return {{ShortName}}._table.Table({{.FieldNumber}})

}
`)
	if err != nil {
		t.zap.Error("StringAccessor: template definition error", zap.String("err", err.Error()))
		return
	}

	err = tmpl.Execute(wr, map[string]interface{}{"Table": t.table, "Field": freader, "FieldNumber": fn})
	if err != nil {
		t.zap.Error("StringAccessor: template render error", zap.String("err", err.Error()))
		return
	}

	return

}

//vectorscalar
//struct
//string vector string
//table vector table
//union vector union
//resource
