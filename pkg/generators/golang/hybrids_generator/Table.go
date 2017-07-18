package hybrids_generator

import (
	"text/template"
	"io"
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
	"go.uber.org/zap"
)

type TableReaderGenerator struct {
	table corev1.TableReader
	zap   *zap.Logger
}

func NewTableReaderGenerator(table corev1.TableReader, logger *zap.Logger) *TableReaderGenerator {
	table = table
	zap := logger.With(zap.String("TableName", table.Meta().Name()),
		zap.String("Type", "Reader implementation"),
		zap.String("Application", table.Meta().Application()),
	)
	return &TableReaderGenerator{table: table, zap: zap}
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
	return string(t.table.Meta().Name()[0])
}

func (t *TableReaderGenerator) Table() corev1.TableReader {
	return t.table
}

func (t *TableReaderGenerator) Generate(wr io.Writer) (err error) {
	var i int32
	var field corev1.FieldReader
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
	return
}

//default
//documentation
func (t *TableReaderGenerator) StringAccessor(freader corev1.FieldReader, fn uint16, wr io.Writer) (err error) {
	tmpl, err := template.New("StringAccessor").Funcs(map[string]interface{}{"ShortName": t.ShortName()}).Parse(`
func ({{ShortName}} *{{.Table.Meta.Name}}) {{.Field.Name}}() (value string) {
	value, _ = {{ShortName}}.table.String({{.FieldNumber}})
	return
}`)
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
