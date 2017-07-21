package interface_generator

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
		zap.String("Type", "Reader Interface"),
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

func (t *TableReaderGenerator) ShortName() string {
	return strings.ToLower(string(t.table.Meta().Name()[0]))
}

func (t *TableReaderGenerator) Table() corev1.TableReader {
	return t.table
}
func (t *TableReaderGenerator) StartInterface() (err error) {
	tmpl, err := template.New("TableReaderGenerator").Funcs(t.funcMap).Parse(`
{{GoDoc (print (TableName .) "Reader") .Meta.Documentation}}
type {{TableName .}}Reader interface {
`)
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

func (t *TableReaderGenerator) EndInterface() (err error) {
	_, err = t.definitionsBuffer.Write([]byte("\n" + "}\n"))
	return
}

func (t *TableReaderGenerator) CreateAccessors() (err error) {
	var i int32
	var field corev1.FieldReader

	//create Accessors
	for i = 0; i < t.table.Fields().Len(); i++ {
		field, err = t.table.Fields().Get(i)
		fieldNumber :=  uint16(i)
		if err != nil {
			t.zap.Error(err.Error())
			return
		}
		switch field.Type() {
		case "String":
			err = t.StringAccessor(field, fieldNumber)
			if err != nil {
				return
			}
		case "Vector":
			switch field.Items() {
			case "String":
				err = t.VectorStringAccessor(field, fieldNumber)
				if err != nil {
					return
				}
			default:
				pid := corev1Native.NewIDReader([]byte(t.table.Meta().Application()+"/"+field.Items()), false)
				if pid != nil {
					if pid.Kind() == "Table" {
						err = t.VectorTableAccessor(field, fieldNumber, utils.TableNameFromID(pid))
						if err != nil {
							return
						}
					}
				}
			}
		default:
			pid := corev1Native.NewIDReader([]byte(t.table.Meta().Application()+"/"+field.Type()), false)
			if pid != nil {
				if pid.Kind() == "Table" {
					err = t.TableAccessor(field, fieldNumber, pid.ID())
					if err != nil {
						return
					}

				}
			}

		}
	}
	return
}

func (t *TableReaderGenerator) FlushBuffers(wr io.Writer) (err error) {
	_, err = t.definitionsBuffer.WriteTo(wr)
	if err != nil {
		return err
	}
	_, err = t.functionsBuffer.WriteTo(wr)
	if err != nil {
		return err
	}
	return
}

func (t *TableReaderGenerator) Generate(wr io.Writer) (err error) {

	//add imports
	wr.Write([]byte(`

`))
	err = t.StartInterface()
	if err != nil {
		return err
	}

	err = t.CreateAccessors()
	if err != nil {
		return err
	}


	err = t.CreateVector()
	if err != nil {
		return err
	}

	err = t.EndInterface()
	if err != nil {
		return err
	}

	err = t.FlushBuffers(wr)
	if err != nil {
		return err
	}

	t.zap.Info(fmt.Sprintf("Interface for Table %s Created successfully", utils.TableName(t.table)))
	return
}

//Todo:
//default
//resource
func (t *TableReaderGenerator) StringAccessor(freader corev1.FieldReader, fn uint16) (err error) {
	tmpl, err := template.New("StringAccessor").
		Funcs(t.funcMap).Parse(`
    {{GoDoc .Field.Name .Field.Documentation}}
    {{Capitalize .Field.Name}}() string
`)
	if err != nil {
		return
	}

	err = tmpl.Execute(t.definitionsBuffer, map[string]interface{}{"Table": t.table, "Field": freader, "FieldNumber": fn})
	if err != nil {
		return
	}
	return
}

func (t *TableReaderGenerator) VectorStringAccessor(freader corev1.FieldReader, fn uint16) (err error) {
	tmpl, err := template.New("StringAccessor").
		Funcs(t.funcMap).Parse(`
    {{GoDoc .Field.Name .Field.Documentation}}
    {{Capitalize .Field.Name}}() hybrids.VectorStringReader
`)
	if err != nil {
		return
	}

	err = tmpl.Execute(t.definitionsBuffer, map[string]interface{}{"Table": t.table, "Field": freader, "FieldNumber": fn})
	if err != nil {
		return
	}

	return

}

func (t *TableReaderGenerator) VectorTableAccessor(freader corev1.FieldReader, fn uint16, tableName string) (err error) {
	tmpl, err := template.New("StringAccessor").
		Funcs(t.funcMap).Parse(`
    {{GoDoc .Field.Name .Field.Documentation}}
    {{Capitalize .Field.Name}}() {{.TypeTableName}}Reader
`)
	if err != nil {
		return
	}

	err = tmpl.Execute(t.definitionsBuffer, map[string]interface{}{"Table": t.table,
		"Field":                                                          freader,
		"FieldNumber":                                                    fn,
		"TypeTableName":                                                  "Vector" + tableName,
		"PackageName":                                                    t.interfacePackageShort,
	})
	if err != nil {
		return
	}

	return
}

func (t *TableReaderGenerator) TableAccessor(freader corev1.FieldReader, fn uint16, tableName string) (err error) {

	tmpl, err := template.New("StringAccessor").
		Funcs(t.funcMap).Parse(`
    {{GoDoc .Field.Name .Field.Documentation}}
    {{Capitalize .Field.Name}}() {{.TypeTableName}}Reader
`)
	if err != nil {
		return
	}

	err = tmpl.Execute(t.definitionsBuffer, map[string]interface{}{"Table": t.table,
		"Field":                                                          freader,
		"FieldNumber":                                                    fn,
		"TypeTableName":                                                  tableName,
		"PackageName":                                                    t.interfacePackageShort,
	})
	if err != nil {
		return
	}

	return

}

func (t *TableReaderGenerator) CreateVector() (err error) {

	tmpl, err := template.New("TableReaderGenerator::GenerateVector").
		Funcs(t.funcMap).Parse(`
type Vector{{TableName .Table}}Reader interface {
     Len() int
     Get(i int) (item {{TableName .Table}}Reader, err error)
}
`)
	if err != nil {
		return
	}

	err = tmpl.Execute(t.functionsBuffer, map[string]interface{}{
		"Table":       t.table,
	})
	if err != nil {
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
