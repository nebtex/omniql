package native_generator

import (
	"text/template"
	"io"
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
	"github.com/nebtex/omnibuff/pkg/utils"
	"go.uber.org/zap"
	"fmt"
	"strings"
	"github.com/nebtex/omnibuff/pkg/generators/golang"
	"bytes"
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1Native"
)

type GoTypeGenerator struct {
	table        corev1.TableReader
	zap          *zap.Logger
	structBuffer *bytes.Buffer
	fields [][]string
}

func NewGoTypeGenerator(table corev1.TableReader, logger *zap.Logger) *GoTypeGenerator {
	zap := logger.With(zap.String("TableName", table.Metadata().Name()),
		zap.String("Type", "Reader implementation"),
		zap.String("Application", table.Metadata().Application()),
	)

	t := &GoTypeGenerator{table: table, zap: zap}
	t.structBuffer = bytes.NewBuffer(nil)

	return t
}

func (g *GoTypeGenerator) StartStruct() (err error) {
	tmpl, err := template.New("TableReaderGenerator").Funcs(golang.DefaultTemplateFunctions).Parse(`
{{GoDoc (print (TableName .)) .Metadata.Documentation}}
type {{TableName .}} struct {
`)
	if err != nil {
		return
	}
	err = tmpl.Execute(g.structBuffer, g.table)
	return
}

func (g *GoTypeGenerator) StructAddField(fn string, ft string) (err error) {
	g.fields = append(g.fields, []string{fn, ft})
	return
}

func (g *GoTypeGenerator) EndStruct() (err error) {
	_, err = g.structBuffer.Write([]byte("\n" + "}\n"))
	return
}

func (g *GoTypeGenerator) Generate(wr io.Writer) (err error) {

	err = g.StartStruct()
	if err != nil {
		return err
	}

	err = g.CreateAccessors()
	if err != nil {
		return err
	}
	for _,field:=range g.fields {
		_, err = g.structBuffer.Write([]byte("\n    " + field[0] + " " + field[1] + fmt.Sprintf(" `json:\"%s\"`", strings.ToLower(field[0]))))
		if err != nil {
			return err
		}
	}

	err = g.EndStruct()
	if err != nil {
		return err
	}
	_, err = g.structBuffer.WriteTo(wr)

	if err != nil {
		return err
	}

	g.zap.Info(fmt.Sprintf("Table %s Created successfully", utils.TableName(g.table)))
	return
}

func (t *GoTypeGenerator) CreateAccessors() (err error) {
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
			err = t.StringAccessor(field)
			if err != nil {
				return
			}
		case "Vector":
			switch field.Items() {
			case "String":
				err = t.VectorStringAccessor(field)
				if err != nil {
					return
				}
			default:
				pid := corev1Native.NewIDReader([]byte(t.table.Metadata().Application()+"/"+field.Items()), false)
				if pid != nil {
					if pid.Kind() == "Table" {
						err = t.VectorTableAccessor(field, utils.TableNameFromID(pid))
						if err != nil {
							return
						}
					}
				}
			}
		default:
			pid := corev1Native.NewIDReader([]byte(t.table.Metadata().Application()+"/"+field.Type()), false)
			if pid != nil {
				if pid.Kind() == "Table" {
					err = t.TableAccessor(field, pid.ID())
					if err != nil {
						return
					}

				}
			}

		}
	}
	return
}

func (t *GoTypeGenerator) StringAccessor(freader corev1.FieldReader) (err error) {
	err = t.StructAddField(strings.Title(freader.Name()), "string")
	return
}

func (g *GoTypeGenerator) VectorStringAccessor(freader corev1.FieldReader) (err error) {
	err = g.StructAddField(strings.Title(freader.Name()), "[]string")
	return
}

func (g *GoTypeGenerator) VectorTableAccessor(freader corev1.FieldReader, tableName string) (err error) {
	err = g.StructAddField(strings.Title(freader.Name()), "[]*"+tableName)
	return
}

func (g *GoTypeGenerator) TableAccessor(freader corev1.FieldReader, tableName string) (err error) {

	err = g.StructAddField(strings.Title(freader.Name()), "*"+tableName)
	return

}
