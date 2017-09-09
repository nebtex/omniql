package native_generator

import (
	"text/template"
	"io"
	"github.com/nebtex/omniql/pkg/io/omniql/corev1"
	"github.com/nebtex/omniql/pkg/utils"
	"go.uber.org/zap"
	"fmt"
	"strings"
	"github.com/nebtex/omniql/pkg/io/omniql/corev1Native"
	"bytes"
	"github.com/nebtex/omniql/pkg/generators/golang"
)

type TableReaderGenerator struct {
	table                 corev1.TableReader
	interfacePackage      string
	interfacePackageShort string
	zap                   *zap.Logger
	packagesBuffer        *bytes.Buffer
	structBuffer          *bytes.Buffer
	implementationBuffer  *bytes.Buffer
	functionsBuffer       *bytes.Buffer
	imports               *golang.Imports
	gt                    *GoTypeGenerator
}

func NewTableReaderGenerator(table corev1.TableReader, ip string, logger *zap.Logger) *TableReaderGenerator {
	table = table
	zap := logger.With(zap.String("TableName", table.Metadata().Name()),
		zap.String("Type", "Reader implementation"),
		zap.String("Application", table.Metadata().Application()),
	)

	t := &TableReaderGenerator{table: table, zap: zap}
	t.functionsBuffer = bytes.NewBuffer(nil)
	t.implementationBuffer = bytes.NewBuffer(nil)
	t.structBuffer = bytes.NewBuffer(nil)
	t.interfacePackage = ip
	items := strings.Split(ip, "/")
	t.interfacePackageShort = items[len(items)-1]
	t.imports = golang.NewImports()
	t.gt = NewGoTypeGenerator(t.table, t.zap)

	return t
}

func (t *TableReaderGenerator) CreateAllocator() (err error) {

	tmpl, err := template.New("TableReaderGenerator::CreateAllocator").
		Funcs(golang.DefaultTemplateFunctions).Funcs(map[string]interface{}{
		"DeepInit": func(reader corev1.TableReader) (value string, err error) {
			var field corev1.FieldReader
			fields := reader.Fields()
			for i := 0; i < fields.Len(); i++ {
				field, err = fields.Get(i)
				if err != nil {
					return
				}
				pid := corev1Native.NewIDReader([]byte(t.table.Metadata().Application()+"/"+field.Type()), false)
				if pid == nil {
					continue
				}
				if pid.Kind() == "Table" {
					value += fmt.Sprintf("\n%s: %s,", strings.ToLower(field.Name()), fmt.Sprintf("New%sReader(%s.%s)", utils.TableNameFromID(pid), golang.ShortName(utils.TableName(reader)), strings.Title(field.Name())))

				}

			}
			return
		},
	}).Parse(`
//New{{TableName .Table}}Reader ...
func New{{TableName .Table}}Reader({{ShortName (TableName .Table)}} *{{TableName .Table}}) *{{TableName .Table}}Reader{
	if {{ShortName  (TableName .Table)}}!=nil{
		return &{{TableName .Table}}Reader{
		                                   _{{ToLower (TableName .Table)}}:{{ShortName (TableName .Table)}},{{DeepInit .Table}}
		                                   }
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
	tmpl, err := template.New("TableReaderGenerator::Native::StartStruct").Funcs(golang.DefaultTemplateFunctions).Parse(`
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
	var field corev1.FieldReader

	//create Accessors
	for i := 0; i < t.table.Fields().Len(); i++ {
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

	err = t.gt.Generate(wr)
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

	tmpl, err := template.New("TableReaderGenerator::Native::StringAccessor").
		Funcs(golang.DefaultTemplateFunctions).Parse(`
{{GoDoc .Field.Name .Field.Documentation}}
func ({{ShortName (TableName .Table)}} *{{TableName .Table}}Reader) {{Capitalize .Field.Name}}() (value string) {
	value = {{ShortName (TableName .Table)}}._{{ToLower (TableName .Table)}}.{{Capitalize .Field.Name}}
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
		Funcs(golang.DefaultTemplateFunctions).Parse(`
{{GoDoc .Field.Name .Field.Documentation}}
func ({{ShortName (TableName .Table)}} *{{TableName .Table}}Reader) {{Capitalize .Field.Name}}() hybrids.VectorStringReader {
	if {{ShortName (TableName .Table)}}.{{ToLower .Field.Name}} != nil {
		return {{ShortName (TableName .Table)}}.{{ToLower .Field.Name}}
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

	tmpl, err := template.New("TableReaderGenerator::Native::VectorTableAccessor").
		Funcs(golang.DefaultTemplateFunctions).Parse(`
{{GoDoc .Field.Name .Field.Documentation}}
func ({{ShortName (TableName .Table)}} *{{TableName .Table}}Reader) {{Capitalize .Field.Name}}() {{.PackageName}}.{{.TypeTableName}}Reader {

	if {{ShortName (TableName .Table)}}.{{ToLower .Field.Name}} != nil {
		return {{ShortName (TableName .Table)}}.{{ToLower .Field.Name}}
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
		Funcs(golang.DefaultTemplateFunctions).Parse(`
{{$table_name:=(TableName .Table)}}
{{$short_name:=(ShortName $table_name)}}
{{GoDoc .Field.Name .Field.Documentation}}
func ({{$short_name}} *{{$table_name}}Reader) {{Capitalize .Field.Name}}() ({{ShortName (print .TypeTableName "Reader")}} {{.PackageName}}.{{.TypeTableName}}Reader, err error) {

	if {{$short_name}}.{{ToLower .Field.Name}} != nil {
		{{ShortName (print .TypeTableName "Reader")}}  =  {{$short_name}}.{{ToLower .Field.Name}}
	}

	return
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
	t.imports.AddImport("github.com/nebtex/hybrids/golang/hybrids")
	t.imports.AddImport(t.interfacePackage)

	tmpl, err := template.New("TableReaderGenerator::GenerateVector").
		Funcs(golang.DefaultTemplateFunctions).Parse(`
{{$table_name:=(TableName .Table)}}
{{$short_name:=(ShortName $table_name)}}

//Vector{{$table_name}}Reader ...
type Vector{{$table_name}}Reader struct {
    _vector  []*{{$table_name}}Reader
}

//Len Returns the current size of this vector
func (v{{$short_name}} *Vector{{$table_name}}Reader) Len() (size int) {
    size = len(v{{$short_name}}._vector)
    return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
func (v{{$short_name}} *Vector{{$table_name}}Reader) Get(i int) (item {{.PackageName}}.{{$table_name}}Reader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(v{{$short_name}}._vector)}
		return
	}

	if i > len(v{{$short_name}}._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(v{{$short_name}}._vector)}
		return
	}

	item = v{{$short_name}}._vector[i]
	return


}

//NewVector{{$table_name}}Reader ...
func NewVector{{$table_name}}Reader(v{{$short_name}} []*{{$table_name}}) (v{{ShortName (print $table_name "Reader")}} *Vector{{$table_name}}Reader) {
    v{{ShortName (print $table_name "Reader")}} = &Vector{{$table_name}}Reader{}
	v{{ShortName (print $table_name "Reader")}}._vector = make([]*{{$table_name}}Reader, len(v{{$short_name}}))

	for i := 0; i < len(v{{$short_name}}); i++ {
		v{{ShortName (print $table_name "Reader")}}._vector[i] = New{{$table_name}}Reader(v{{$short_name}}[i])
	}
	return
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
