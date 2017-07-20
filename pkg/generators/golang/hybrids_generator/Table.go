package hybrids_generator

import (
	"text/template"
	"io"
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
	"github.com/nebtex/omnibuff/pkg/utils"
	"go.uber.org/zap"
	"fmt"
	"strings"
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1Native"
	"bytes"
)

type TableReaderGenerator struct {
	table                 corev1.TableReader
	interfacePackage      string
	interfacePackageShort string
	zap                   *zap.Logger
	funcMap               map[string]interface{}
	packagesBuffer        *bytes.Buffer
	definitionsBuffer     *bytes.Buffer
	functionsBuffer       *bytes.Buffer
}

func NewTableReaderGenerator(table corev1.TableReader, ip string, logger *zap.Logger) *TableReaderGenerator {
	table = table
	zap := logger.With(zap.String("TableName", table.Meta().Name()),
		zap.String("Type", "Reader implementation"),
		zap.String("Application", table.Meta().Application()),
	)

	t := &TableReaderGenerator{table: table, zap: zap}
	t.funcMap = map[string]interface{}{
		"TableName":      utils.TableName,
		"GetPackageName": utils.GolangGetPackageName,
		"ShortName":      t.ShortName,
		"ToLower":        strings.ToLower,
		"Capitalize":     strings.Title,
		"GoDoc": func(name string, d corev1.DocumentationReader) (value string) {
			if d != nil {
				if d.Short() == "" && d.Long() == "" {
					return "//" + strings.Title(name) + " ..."
				}

				if d.Short() != "" && d.Long() != "" {
					return strings.Title(name) + " " + d.Short() + "\n" + d.Long()
				}
				if d.Short() != "" {
					return utils.CommentGolang(strings.Title(name) + " " + d.Short())
				}
				if d.Long() != "" {
					return utils.CommentGolang(strings.Title(name) + " " + d.Long())
				}
			}
			return "//" + strings.Title(name) + " ..."
		}}
	t.functionsBuffer = bytes.NewBuffer(nil)
	t.definitionsBuffer = bytes.NewBuffer(nil)
	t.interfacePackage = ip
	items := strings.Split(ip, "/")
	t.interfacePackageShort = items[len(items)-1]
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
func (t *TableReaderGenerator) CreateAllocator() (err error) {

	tmpl, err := template.New("TableReaderGenerator::CreateAllocator").Funcs(t.funcMap).Parse(`
func New{{TableName .Table}}Reader(t hybrids.TableReader) {{.PackageName}}.{{TableName .Table}}Reader{
	if t==nil{
		return nil
	}
	return &{{TableName .Table}}Reader{_table:t}
}`)
	if err != nil {
		return
	}

	err = tmpl.Execute(t.functionsBuffer, map[string]interface{}{
		"Table":       t.Table(),
		"PackageName": t.interfacePackageShort, })
	return

}
func (t *TableReaderGenerator) ShortName() string {
	return strings.ToLower(string(t.table.Meta().Name()[0]))
}

func (t *TableReaderGenerator) Table() corev1.TableReader {
	return t.table
}
func (t *TableReaderGenerator) StartStruct() (err error) {
	tmpl, err := template.New("TableReaderGenerator").Funcs(t.funcMap).Parse(`
{{GoDoc (print (TableName .) "Reader") .Meta.Documentation}}
type {{TableName .}}Reader struct {
    _table hybrids.TableReader`)
	if err != nil {
		return
	}
	err = tmpl.Execute(t.definitionsBuffer, t.Table())
	return
}

func (t *TableReaderGenerator) StructAddField(fn string, ft string) (err error) {
	_, err = t.definitionsBuffer.Write([]byte("\n    " + fn + " " + ft))
	return
}

func (t *TableReaderGenerator) EndStruct() (err error) {
	_, err = t.definitionsBuffer.Write([]byte("\n" + "}\n"))
	return
}

func (t *TableReaderGenerator) Generate(wr io.Writer) (err error) {
	var i int32
	var field corev1.FieldReader

	//add imports
	wr.Write([]byte(`
import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")


`))
	err = t.StartStruct()
	if err != nil {
		return err
	}

	//create Accessors
	for i = 0; i < t.table.Fields().Len(); i++ {
		field, err = t.table.Fields().Get(i)
		if err != nil {
			t.zap.Error(err.Error())
			return
		}
		switch field.Type() {
		case "String":
			err = t.StringAccessor(field, uint16(i))
			if err != nil {
				return err
			}
		default:
			pid := corev1Native.NewIDReader([]byte(t.table.Meta().Application()+"/"+field.Type()), false)
			if pid != nil {
				if pid.Kind() == "Table" {
					err = t.TableAccessor(field, uint16(i), pid.ID())
					if err != nil {
						return err
					}

				}
			}

		}
	}
	err = t.EndStruct()
	if err != nil {
		return err
	}


	_, err = t.definitionsBuffer.WriteTo(wr)
	if err != nil {
		return err
	}


	err= t.CreateAllocator()
	if err != nil {
		return err
	}

	_, err = t.functionsBuffer.WriteTo(wr)
	if err != nil {
		return err
	}
	t.zap.Info(fmt.Sprintf("Table %s Created successfully", utils.TableName(t.table)))
	return
}

//Todo:
//default
//table reader
//vector table reader
//resource
func (t *TableReaderGenerator) StringAccessor(freader corev1.FieldReader, fn uint16) (err error) {
	tmpl, err := template.New("StringAccessor").
		Funcs(t.funcMap).Parse(`
{{GoDoc .Field.Name .Field.Documentation}}
func ({{ShortName}} *{{TableName .Table}}Reader) {{Capitalize .Field.Name}}() (value string) {
	value, _ = {{ShortName}}._table.String({{.FieldNumber}})
	return
}
`)
	if err != nil {
		t.zap.Error("StringAccessor: template definition error", zap.String("err", err.Error()))
		return
	}

	err = tmpl.Execute(t.functionsBuffer, map[string]interface{}{"Table": t.table, "Field": freader, "FieldNumber": fn})
	if err != nil {
		t.zap.Error("StringAccessor: template render error", zap.String("err", err.Error()))
		return
	}

	return

}

func (t *TableReaderGenerator) TableAccessor(freader corev1.FieldReader, fn uint16, tableName string) (err error) {

	err = t.StructAddField(strings.ToLower(freader.Name()), t.interfacePackageShort+"."+tableName+"Reader")
	if err != nil {
		return err
	}

	tmpl, err := template.New("StringAccessor").
		Funcs(t.funcMap).Parse(`
{{GoDoc .Field.Name .Field.Documentation}}
func ({{ShortName}} *{{TableName .Table}}Reader) {{Capitalize .Field.Name}}() {{.PackageName}}.{{.TypeTableName}}Reader {

	if {{ShortName}}.{{ToLower .Field.Name}} != nil {
		return {{ShortName}}.{{ToLower .Field.Name}}
	}

	return New{{.TypeTableName}}Reader({{ShortName}}._table.Table({{.FieldNumber}}))
}
`)
	if err != nil {
		t.zap.Error("TableAccessor: template definition error", zap.String("err", err.Error()))
		return
	}

	err = tmpl.Execute(t.functionsBuffer, map[string]interface{}{"Table": t.table,
		"Field":                                                          freader,
		"FieldNumber":                                                    fn,
		"TypeTableName":                                                  tableName,
		"PackageName":                                                    t.interfacePackageShort,
	})
	if err != nil {
		t.zap.Error("TableAccessor: template render error", zap.String("err", err.Error()))
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
