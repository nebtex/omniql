package native_generator

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
	"github.com/nebtex/omnibuff/pkg/generators/golang"
)

//TODO: create initializators
type TableReaderGenerator struct {
	table                 corev1.TableReader
	interfacePackage      string
	interfacePackageShort string
	zap                   *zap.Logger
	funcMap               map[string]interface{}
	packagesBuffer        *bytes.Buffer
	structBuffer          *bytes.Buffer
	implementationBuffer  *bytes.Buffer
	functionsBuffer       *bytes.Buffer
	imports               *golang.Imports
}

func NewTableReaderGenerator(table corev1.TableReader, ip string, logger *zap.Logger) *TableReaderGenerator {
	table = table
	zap := logger.With(zap.String("TableName", table.Metadata().Name()),
		zap.String("Type", "Reader implementation"),
		zap.String("Application", table.Metadata().Application()),
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
	t.implementationBuffer = bytes.NewBuffer(nil)
	t.structBuffer = bytes.NewBuffer(nil)
	t.interfacePackage = ip
	items := strings.Split(ip, "/")
	t.interfacePackageShort = items[len(items)-1]
	t.imports = golang.NewImports()

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

	tmpl, err := template.New("TableReaderGenerator::CreateAllocator").Funcs(golang.DefaultTemplateFunctions).Parse(`
//New{{TableName .Table}}Reader ...
func New{{TableName .Table}}Reader({{ShortName .Table.Metadata.Name}} *{{TableName .Table}}Reader) {{.PackageName}}.{{TableName .Table}}Reader{
	if {{ShortName .Table.Metadata.Name}}!=nil{
		return &{{TableName .Table}}Reader{_{{ToLower .Table.Metadata.Name}}:{{ShortName .Table.Metadata.Name}}}
	}
	return nil
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
	return strings.ToLower(string(t.table.Metadata().Name()[0]))
}

func (t *TableReaderGenerator) Table() corev1.TableReader {
	return t.table
}

func (t *TableReaderGenerator) StartStruct() (err error) {
	tmpl, err := template.New("TableReaderGenerator::Native::StartStruct").Funcs(t.funcMap).Parse(`
{{GoDoc (print (TableName .) "Reader") .Metadata.Documentation}}
type {{TableName .}}Reader struct {
    _{{ToLower (TableName .)}} *{{TableName .}}`)
	if err != nil {
		return
	}
	err = tmpl.Execute(t.implementationBuffer, t.Table())
	return
}

func (t *TableReaderGenerator) StructAddField(fn string, ft string) (err error) {
	_, err = t.implementationBuffer.Write([]byte("\n    " + fn + " " + ft))
	return
}

func (t *TableReaderGenerator) EndStruct() (err error) {
	_, err = t.implementationBuffer.Write([]byte("\n" + "}\n"))
	return
}

func (t *TableReaderGenerator) CreateAccessors(offset uint16) (err error) {
	var i int32
	var field corev1.FieldReader

	//create Accessors
	for i = 0; i < t.table.Fields().Len(); i++ {
		field, err = t.table.Fields().Get(i)
		fieldNumber := offset + uint16(i)
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
				pid := corev1Native.NewIDReader([]byte(t.table.Metadata().Application()+"/"+field.Items()), false)
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
			pid := corev1Native.NewIDReader([]byte(t.table.Metadata().Application()+"/"+field.Type()), false)
			if pid != nil {
				if pid.Kind() == "Table" {
					err = t.TableAccessor(field, fieldNumber, pid.ID())
					if err != nil {
						return
					}

				}
				if pid.Kind() == "EnumerationGroup" {
					err = t.EnumerationAccessor(field, fieldNumber, pid.Parent().ID(), pid.ID())
					if err != nil {
						return
					}
				}
				if pid.Kind() == "Enumeration" {
					err = t.EnumerationAccessor(field, fieldNumber, pid.ID(), "")
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
	err = t.imports.Write(wr)
	if err != nil {
		return err
	}

	gt := NewGoTypeGenerator(t.table, t.zap)
	err = gt.Generate(wr)
	if err != nil {
		return err
	}

	_, err = t.structBuffer.WriteTo(wr)
	if err != nil {
		return err
	}
	_, err = t.implementationBuffer.WriteTo(wr)
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

	err = t.StartStruct()
	if err != nil {
		return err
	}

	err = t.CreateAccessors(0)
	if err != nil {
		return err
	}

	err = t.CreateAllocator()
	if err != nil {
		return err
	}

	err = t.CreateVector()
	if err != nil {
		return err
	}

	err = t.EndStruct()
	if err != nil {
		return err
	}
	err = t.FlushBuffers(wr)
	if err != nil {
		return err
	}

	t.zap.Info(fmt.Sprintf("Table %s Created successfully", utils.TableName(t.table)))
	return
}

//Todo:
//default
//resource
func (t *TableReaderGenerator) StringAccessor(freader corev1.FieldReader, fn uint16) (err error) {

	tmpl, err := template.New("StringAccessor").
		Funcs(t.funcMap).Parse(`
{{GoDoc .Field.Name .Field.Documentation}}
func ({{ShortName}} *{{TableName .Table}}Reader) {{Capitalize .Field.Name}}() (value string) {
	value = {{ShortName}}._{{ToLower (TableName .Table)}}.{{Capitalize .Field.Name}}
	return
}
`)
	if err != nil {
		return
	}

	err = tmpl.Execute(t.functionsBuffer, map[string]interface{}{"Table": t.table, "Field": freader, "FieldNumber": fn})
	if err != nil {
		return
	}

	return

}

func (t *TableReaderGenerator) EnumerationAccessor(freader corev1.FieldReader, fn uint16, enumeration string, enumerationGroup string) (err error) {
	t.imports.AddImport(t.interfacePackage)

	tmpl, err := template.New("TableReaderGenerator::Native::Enumeration").
		Funcs(golang.DefaultTemplateFunctions).Parse(`
{{GoDoc .Field.Name .Field.Documentation}}
func ({{ShortName .Table.Metadata.Name}} *{{TableName .Table}}Reader) {{Capitalize .Field.Name}}() (value {{.PackageName}}.{{.Enumeration}}) {
	value = {{.PackageName}}.FromStringTo{{.Enumeration}}({{ShortName .Table.Metadata.Name}}._{{ToSnakeCase (TableName .Table)}}.{{Capitalize .Field.Name}})

	{{if .EnumerationGroup}}if !value.Is{{.EnumerationGroup}}(){
		value = {{.PackageName}}.{{.Enumeration}}None
	}{{else}}if !value.IsValid(){
		value = {{.PackageName}}.{{.Enumeration}}None
	}{{end}}

	return
}
`)
	if err != nil {
		return
	}

	err = tmpl.Execute(t.functionsBuffer, map[string]interface{}{
		"Table":            t.table,
		"Field":            freader,
		"FieldNumber":      fn,
		"Enumeration":      enumeration,
		"EnumerationGroup": enumerationGroup,
		"PackageName":      t.interfacePackageShort})
	if err != nil {
		return
	}

	return
}

func (t *TableReaderGenerator) VectorStringAccessor(freader corev1.FieldReader, fn uint16) (err error) {

	err = t.StructAddField(strings.ToLower(freader.Name()), "*native.VectorStringReader")
	if err != nil {
		return
	}
	tmpl, err := template.New("TableReaderGenerator::Native::VectorStringAccessor").
		Funcs(t.funcMap).Parse(`
{{GoDoc .Field.Name .Field.Documentation}}
func ({{ShortName}} *{{TableName .Table}}Reader) {{Capitalize .Field.Name}}() hybrids.VectorStringReader {
	if {{ShortName}}.{{ToLower .Field.Name}} != nil {
		return {{ShortName}}.{{ToLower .Field.Name}}
	}
	return nil
}
`)
	if err != nil {
		return
	}

	err = tmpl.Execute(t.functionsBuffer, map[string]interface{}{"Table": t.table, "Field": freader, "FieldNumber": fn})
	if err != nil {
		return
	}

	return

}

func (t *TableReaderGenerator) VectorTableAccessor(freader corev1.FieldReader, fn uint16, tableName string) (err error) {

	err = t.StructAddField(strings.ToLower(freader.Name()), "*Vector"+tableName+"Reader")
	if err != nil {
		return
	}

	tmpl, err := template.New("StringAccessor").
		Funcs(t.funcMap).Parse(`
{{GoDoc .Field.Name .Field.Documentation}}
func ({{ShortName}} *{{TableName .Table}}Reader) {{Capitalize .Field.Name}}() {{.PackageName}}.{{.TypeTableName}}Reader {

	if {{ShortName}}.{{ToLower .Field.Name}} != nil {
		return {{ShortName}}.{{ToLower .Field.Name}}
	}

	return nil
}
	`)
	if err != nil {
		return
	}

	err = tmpl.Execute(t.functionsBuffer, map[string]interface{}{"Table": t.table,
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

	err = t.StructAddField(strings.ToLower(freader.Name()), "*"+tableName+"Reader")
	if err != nil {
		return err
	}

	tmpl, err := template.New("TableReaderGenerator::Native::TableAccessor").
		Funcs(t.funcMap).Parse(`
{{GoDoc .Field.Name .Field.Documentation}}
func ({{ShortName}} *{{TableName .Table}}Reader) {{Capitalize .Field.Name}}() {{.PackageName}}.{{.TypeTableName}}Reader {

	if {{ShortName}}.{{ToLower .Field.Name}} != nil {
		return {{ShortName}}.{{ToLower .Field.Name}}
	}

	return nil
}
`)
	if err != nil {
		return
	}

	err = tmpl.Execute(t.functionsBuffer, map[string]interface{}{"Table": t.table,
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

type Vector{{TableName .Table}}Reader struct {
    _vector  []*{{TableName .Table}}Reader
}

func (v{{ShortName}} *Vector{{TableName .Table}}Reader) Len() (size int) {
    size = len(v{{ShortName}}._vector)
    return
}

func (v{{ShortName}} *Vector{{TableName .Table}}Reader) Get(i int) (item {{.PackageName}}.{{TableName .Table}}Reader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(v{{ShortName}}._vector)}
		return
	}

	if i > len(v{{ShortName}}._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(v{{ShortName}}._vector)}
		return
	}

	item = v{{ShortName}}._vector[i]
	return


}


func NewVector{{TableName .Table}}Reader(v hybrids.VectorTableReader) {{.PackageName}}.Vector{{TableName .Table}}Reader {
    if v == nil {
        return nil
    }
    return &Vector{{TableName .Table}}Reader{_vectorHybrid: v}
}
`)

	if err != nil {
		return
	}

	err = tmpl.Execute(t.functionsBuffer, map[string]interface{}{
		"Table":       t.table,
		"PackageName": t.interfacePackageShort,
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
