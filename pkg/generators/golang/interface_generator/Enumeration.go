package interface_generator

import (
	"text/template"
	"io"
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
	"github.com/nebtex/omnibuff/pkg/utils"
	"github.com/nebtex/omnibuff/pkg/generators/golang"
	"go.uber.org/zap"
	"fmt"
	"bytes"
	"io/ioutil"
)

type EnumerationGenerator struct {
	enum    corev1.EnumerationReader
	zap     *zap.Logger
	imports *golang.Imports
	hybridsInterfacePackage string
	funcMap map[string]interface{}
}

func RenderEnumItems(e corev1.EnumerationReader) (value string, err error) {
	var item corev1.EnumerationItemReader
	for i := 0; i < e.Items().Len(); i++ {
		item, err = e.Items().Get(i)
		if err != nil {
			return
		}
		value = value + "\n    " + golang.GenerateDocs(item.Name(), item.Documentation())
		value = value + "\n    " + e.Metadata().Name() + item.Name() + " " + e.Metadata().Name() + fmt.Sprintf(" = %d", i+1)
	}
	return
}

func RenderGroupMapsDeclaration(e corev1.EnumerationReader) (value string, err error) {
	var group corev1.EnumerationGroupReader
	for i := 0; i < e.Groups().Len(); i++ {
		group, err = e.Groups().Get(i)
		if err != nil {
			return
		}
		value = value + "\nvar " + utils.ToSnakeCase(e.Metadata().Name()) + "_" + utils.ToSnakeCase(group.Name()) + fmt.Sprintf("_map map[%s]bool", e.Metadata().Name())
	}
	return
}

func RenderEnumerationMainMaps(e corev1.EnumerationReader) (value string, err error) {
	var item corev1.EnumerationItemReader
	value = value + "\n    " + utils.ToSnakeCase(e.Metadata().Name()) + fmt.Sprintf("_map = map[%s]string{}", e.Metadata().Name())
	value = value + "\n    " + utils.ToSnakeCase(e.Metadata().Name()) + fmt.Sprintf("_reverse_map = map[string]%s{}", e.Metadata().Name())

	for i := 0; i < e.Items().Len(); i++ {
		item, err = e.Items().Get(i)
		if err != nil {
			return
		}
		value = value + "\n    " + utils.ToSnakeCase(e.Metadata().Name()) + fmt.Sprintf(`_map[%s] = "%s"`, e.Metadata().Name()+item.Name(), item.Name())
		value = value + "\n    " + utils.ToSnakeCase(e.Metadata().Name()) + fmt.Sprintf(`_reverse_map["%s"] = %s`, item.Name(), e.Metadata().Name()+item.Name())
	}
	return
}

func RenderGroupMaps(e corev1.EnumerationReader) (value string, err error) {
	var group corev1.EnumerationGroupReader
	var item string

	for i := 0; i < e.Groups().Len(); i++ {
		group, err = e.Groups().Get(i)
		if err != nil {
			return
		}
		value = value + "\n    " + utils.ToSnakeCase(e.Metadata().Name()) + "_" + utils.ToSnakeCase(group.Name()) + "_map =  " + fmt.Sprintf("map[%s]bool{}", e.Metadata().Name())

		for j := 0; j < group.Items().Len(); j++ {
			item, err = group.Items().Get(j)
			if err != nil {
				return
			}
			value = value + "\n    " + utils.ToSnakeCase(e.Metadata().Name()) + "_" + utils.ToSnakeCase(group.Name()) + fmt.Sprintf("_map[%s] = true", e.Metadata().Name()+item)
		}
	}
	return
}

func RenderGroupValidators(e corev1.EnumerationReader) (value string, err error) {
	var group corev1.EnumerationGroupReader
	var tmpl *template.Template
	buffer := bytes.NewBuffer(nil)

	for i := 0; i < e.Groups().Len(); i++ {
		group, err = e.Groups().Get(i)
		if err != nil {
			return
		}
		tmpl, err = template.New("RenderGroupValidators").Funcs(golang.DefaultTemplateFunctions).Parse(`
{{GoDoc  (print "Is" .Group.Name) .Group.Documentation}}
func ({{ShortName .Enumeration.Metadata.Name}} {{.Enumeration.Metadata.Name}}) Is{{.Group.Name}}() (result bool) {
    _,result = {{ToSnakeCase .Enumeration.Metadata.Name}}_{{ToSnakeCase .Group.Name}}_map[{{ShortName .Enumeration.Metadata.Name}}]
    return

}`)
		if err != nil {
			return
		}

		tmpl.Execute(buffer, map[string]interface{}{
			"Group":       group,
			"Enumeration": e,
		})
	}
	b, err := ioutil.ReadAll(buffer)
	if err != nil {
		return
	}

	value = string(b)
	return
}

func NewEnumerationGenerator(enum corev1.EnumerationReader, logger *zap.Logger) *EnumerationGenerator {
	zap := logger.With(zap.String("EnumerationName", enum.Metadata().Name()),
		zap.String("Type", "Enumeration"),
		zap.String("Application", enum.Metadata().Application()),
	)

	t := &EnumerationGenerator{enum: enum, zap: zap}
	t.funcMap = map[string]interface{}{
		"RenderEnumItems":            RenderEnumItems,
		"RenderGroupMapsDeclaration": RenderGroupMapsDeclaration,
		"RenderEnumerationMainMaps":  RenderEnumerationMainMaps,
		"RenderGroupMaps":            RenderGroupMaps,
		"RenderGroupValidators":      RenderGroupValidators,

	}
	t.hybridsInterfacePackage = "github.com/nebtex/hybrids/golang/hybrids"

	t.imports = golang.NewImports()

	return t
}

func (e *EnumerationGenerator) Generate(wr io.Writer) (err error) {
	e.imports.AddImport(e.hybridsInterfacePackage)
	e.imports.Write(wr)
	tmpl, err := template.New("EnumerationGenerator").Funcs(golang.DefaultTemplateFunctions).
		Funcs(e.funcMap).Parse(`
{{GoDoc .Enumeration.Metadata.Name .Enumeration.Metadata.Documentation}}
type {{.Enumeration.Metadata.Name}} {{ScalarToGolangType .Enumeration.Kind}}

const (
    //{{.Enumeration.Metadata.Name}}None ...
    {{.Enumeration.Metadata.Name}}None {{.Enumeration.Metadata.Name}} = 0{{RenderEnumItems .Enumeration}}
)

var {{ToSnakeCase .Enumeration.Metadata.Name}}_map map[{{.Enumeration.Metadata.Name}}]string
var {{ToSnakeCase .Enumeration.Metadata.Name}}_reverse_map map[string]{{.Enumeration.Metadata.Name}}
{{RenderGroupMapsDeclaration .Enumeration}}

func init(){
	//init maps
{{RenderEnumerationMainMaps .Enumeration}}

{{RenderGroupMaps .Enumeration}}
}

//String stringer
func ({{ShortName .Enumeration.Metadata.Name}} {{.Enumeration.Metadata.Name}}) String() (value string) {
	value = {{ToSnakeCase .Enumeration.Metadata.Name}}_map[{{ShortName .Enumeration.Metadata.Name}}]
	return
}

//IsValid check if the variable has a valid enumeration value
func ({{ShortName .Enumeration.Metadata.Name}} {{.Enumeration.Metadata.Name}}) IsValid() (result bool) {
	_, result = {{ToSnakeCase .Enumeration.Metadata.Name}}_map[{{ShortName .Enumeration.Metadata.Name}}]
	return
}

{{RenderGroupValidators .Enumeration}}

//{{.Enumeration.Metadata.Name}}FromString convert a string to its {{.Enumeration.Metadata.Name}} representation
func FromStringTo{{.Enumeration.Metadata.Name}}(str string) (value {{.Enumeration.Metadata.Name}}) {
    var ok bool
    value, ok = {{ToSnakeCase .Enumeration.Metadata.Name}}_reverse_map[str]
    if !ok{
        value = {{.Enumeration.Metadata.Name}}None
	}
	return
}

//Vector{{.Enumeration.Metadata.Name}} ...
type Vector{{.Enumeration.Metadata.Name}} interface {

	// Returns the current size of this vector
	Len() int

	// Get the item in the position i, if i < Len(),
	// if item does not exist should return the default value for the underlying data type
	// when i > Len() should return an VectorInvalidIndexError
	Get(i int) ({{.Enumeration.Metadata.Name}}, error)
}

type vector_{{ToSnakeCase .Enumeration.Metadata.Name}} struct {
	_vector []{{.Enumeration.Metadata.Name}}
}

//Len Returns the current size of this vector
func (v{{ShortName .Enumeration.Metadata.Name}} *vector_{{ToSnakeCase .Enumeration.Metadata.Name}}) Len() (size int) {
	size = len(v{{ShortName .Enumeration.Metadata.Name}}._vector)
	return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
func (v{{ShortName .Enumeration.Metadata.Name}} *vector_{{ToSnakeCase .Enumeration.Metadata.Name}}) Get(i int) (item {{.Enumeration.Metadata.Name}}, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(v{{ShortName .Enumeration.Metadata.Name}}._vector)}
		return
	}

	if i > len(v{{ShortName .Enumeration.Metadata.Name}}._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(v{{ShortName .Enumeration.Metadata.Name}}._vector)}
		return
	}

	item = v{{ShortName .Enumeration.Metadata.Name}}._vector[i]
	return

}

//NewVector{{.Enumeration.Metadata.Name}} ...
func NewVector{{.Enumeration.Metadata.Name}}(v []{{.Enumeration.Metadata.Name}}) Vector{{.Enumeration.Metadata.Name}} {
	return &vector_{{ToSnakeCase .Enumeration.Metadata.Name}}{_vector: v}
}
`)
	if err != nil {
		return err
	}

	tmpl.Execute(wr, map[string]interface{}{
		"Enumeration": e.enum,
	})

	e.zap.Info(fmt.Sprintf("Enumeration %s Created successfully", e.enum.Metadata().Name()))
	return
}
