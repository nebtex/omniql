package main

import (
	"io/ioutil"
	"path/filepath"
	"github.com/ghodss/yaml"
	"github.com/mitchellh/mapstructure"
	"os"
	"github.com/jjeffery/errors"
	"text/template"
	"strings"
	"fmt"
	"io"
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1Native"
	"github.com/nebtex/omnibuff/pkg/generators/golang/hybrids_generator"
	"go.uber.org/zap"
	"github.com/nebtex/omnibuff/pkg/utils"
	"github.com/nebtex/omnibuff/pkg/generators/golang/interface_generator"
)

var OM map[string]*Meta

func init() {
	OM = map[string]*Meta{}

}

type Documentation struct {
	Short string `json:"short"`
	Long  string `json:"long"`
}

type Meta struct {
	Application   string `json:"application"`
	Name          string `json:"name"`
	Kind          string `json:"kind"`
	Parent        string `json:"parent"`
	Resource      string `json:"resource"`
	Documentation *Documentation `json:"documentation"`
}

type Field struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	Items         string `json:"items"`
	Documentation *Documentation `json:"documentation"`
	Required      bool `json:"required"`
	Default       string `json:"default"`
}

type TableSpec struct {
	Meta   *Meta `json:"meta"`
	Fields []*corev1Native.Field `json:"fields"`
}

func (t *TableSpec) ToFlatBuffer(wr io.Writer) (error) {
	funcMap := template.FuncMap{
		"tableName": func(meta *Meta) string {
			if meta.Parent == "" {
				return meta.Name
			}
			types := strings.Split(meta.Parent, "/")

			return types[1] + meta.Name
		},
		"fieldName": func(str string) string {
			return strings.Title(str)
		},
		"toFbsSchema": func(field *Field) string {
			if strings.Contains(field.Type, "/") {
				types := strings.Split(field.Type, "/")
				if len(types) == 2 {
					return fmt.Sprintf("%s", types[1])
				}

				panic("Schema[Table] not found: " + field.Type)
			}
			switch field.Type {
			case "Integer":
				return "int"
			case "String":
				return "string"
			case "Boolean":
				return "bool"
			case "Vector":
				types := strings.Split(field.Items, "/")
				if len(types) == 1 {
					return fmt.Sprintf("[%s]", strings.ToLower(types[0]))
				}
				if len(types) == 2 {
					return fmt.Sprintf("[%s]", types[1])

				}

				panic("Vector[table] not found: " + field.Items)

			default:
				panic("Schema[table] not found: " + field.Type)
			}
		},
	}

	tmpl, err := template.New("table").Funcs(funcMap).Parse(`
table {{ tableName .Meta }} {
{{range .Fields}}    {{fieldName .Name}}:{{toFbsSchema .}};
{{end}}}
`)
	err = tmpl.Execute(wr, t)

	if err != nil {
		return err
	}

	return nil
}

func (t *TableSpec) ToStreamInterface(wr io.Writer) (error) {
	funcMap := template.FuncMap{
		"tableName": func(meta *Meta) string {

			if meta.Parent == "" {
				return meta.Name
			}
			types := strings.Split(meta.Parent, "/")
			return types[1] + meta.Name
		},
		"fieldName": func(str string) string {

			return strings.Title(str)
		},
		"toStreamSchema": func(field *Field) string {

			if strings.Contains(field.Type, "/") {
				types := strings.Split(field.Type, "/")
				if len(types) == 2 {
					return fmt.Sprintf("(%s, error)", types[1])
				}

				panic("Schema[Table] not found: " + field.Type)
			}
			switch field.Type {
			case "Integer":
				return "int"
			case "String":
				return "string"
			case "Boolean":
				return "bool"
			case "Vector":
				types := strings.Split(field.Items, "/")
				if len(types) == 1 {
					return fmt.Sprintf("Vector%s", strings.Title(types[0]))
				}
				if len(types) == 2 {
					return fmt.Sprintf("(Vector%sStreamReader, error)", types[1])

				}

				panic("Vector[table] not found: " + field.Items)

			default:
				panic("Schema[table] not found: " + field.Type)
			}
		},
	}

	tmpl, err := template.New("table").Funcs(funcMap).Parse(`
type  {{ tableName .Meta }}StreamReader interface{
{{range .Fields}}    {{fieldName .Name}}() {{toStreamSchema .}}
{{end}}}

}
`)
	err = tmpl.Execute(wr, t)

	if err != nil {
		return err
	}

	return nil
}

func (t *ResourceSpec) ToFlatBuffer(wr io.Writer) (error) {
	funcMap := template.FuncMap{
		"tableName": func(meta *Meta) string {
			if meta.Parent == "" {
				return meta.Name
			}
			types := strings.Split(meta.Parent, "/")

			return types[1] + meta.Name
		},
		"fieldName": func(str string) string {
			return strings.Title(str)
		},
		"toFbsSchema": func(field *Field) string {
			if strings.Contains(field.Type, "/") {
				types := strings.Split(field.Type, "/")
				if len(types) == 2 {
					return fmt.Sprintf("%s", types[1])
				}
				if len(types) == 4 {
					if types[0] == "Enumeration" {
						return types[1]
					}

				}

				panic("Schema[Resource] not found: " + field.Type)
			}
			switch field.Type {
			case "Integer":
				return "int"
			case "String":
				return "string"
			case "Boolean":
				return "bool"
			case "Vector":
				types := strings.Split(field.Items, "/")
				if len(types) == 1 {
					return fmt.Sprintf("[%s]", strings.ToLower(types[0]))
				}
				if len(types) == 2 {
					return fmt.Sprintf("[%s]", types[1])

				}

				panic("Vector[Resource] not found: " + field.Items)

			default:
				panic("Schema[Resource] not found: " + field.Type)
			}
		},
	}

	tmpl, err := template.New("table").Funcs(funcMap).Parse(`
table {{ tableName .Meta }} {
{{range .Fields}}    {{fieldName .Name}}:{{toFbsSchema .}};
{{end}}}
`)
	err = tmpl.Execute(wr, t)

	if err != nil {
		return err
	}

	return nil
}

func (t *ResourceSpec) ToStreamInterface(wr io.Writer) (error) {
	funcMap := template.FuncMap{
		"tableName": func(meta *Meta) string {

			if meta.Parent == "" {
				return meta.Name
			}
			types := strings.Split(meta.Parent, "/")
			return types[1] + meta.Name
		},
		"fieldName": func(str string) string {

			return strings.Title(str)
		},
		"toStreamSchema": func(field *Field) string {

			if strings.Contains(field.Type, "/") {
				types := strings.Split(field.Type, "/")
				if len(types) == 2 {
					return fmt.Sprintf("(%s, error)", types[1])
				}
				if len(types) == 4 {
					return ""
				}

				panic("Schema[Table] not found: " + field.Type)
			}
			switch field.Type {
			case "Integer":
				return "int"
			case "String":
				return "string"
			case "Boolean":
				return "bool"
			case "Vector":
				types := strings.Split(field.Items, "/")
				if len(types) == 1 {
					return fmt.Sprintf("Vector%s", strings.Title(types[0]))
				}
				if len(types) == 2 {
					return fmt.Sprintf("(Vector%sStreamReader, error)", types[1])

				}

				panic("Vector[table] not found: " + field.Items)

			default:
				panic("Schema[table] not found: " + field.Type)
			}
		},
	}

	tmpl, err := template.New("table").Funcs(funcMap).Parse(`
type  {{ tableName .Meta }}StreamReader interface{
May be greater than the caccessoruntouchedurrent vector
    {{fieldName .Name}}() {{toStreamSchema .}}
{{end}}}


`)
	err = tmpl.Execute(wr, t)

	if err != nil {
		return err
	}

	return nil
}

type UnionResource struct {
	Items []string `json:"Items"`
}

type UnionSpec struct {
	Meta *Meta `json:"meta"`
	Type map[string]interface{} `json:"type"`
}

type ResourceSpec struct {
	Meta   *Meta `json:"meta"`
	Fields []*corev1Native.Field `json:"fields"`
}

type YamlFile struct {
	Api   string `json:"api"`
	OqlID string `json:"oqlid"`
	Spec  interface{} `json:"spec"`
}

type Application struct {
	Name      string `json:"name"`
	Resources []*corev1Native.Resource `json:"resources"`
	Tables    []*corev1Native.Table `json:"tables"`
	Unions    []*UnionSpec `json:"tables"`
}

func Load(path string) (app *Application, err error) {
	app = &Application{}
	err = filepath.Walk(path, func(subPath string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if info.IsDir() {
			return nil
		}
		data, err := ioutil.ReadFile(subPath)
		if err != nil {
			return err
		}
		yf := &YamlFile{}
		yaml.Unmarshal(data, yf)
		types := strings.Split(yf.OqlID, "/")
		switch types[len(types)-2] {
		case "Table":
			table := &corev1Native.Table{}
			err = mapstructure.Decode(yf.Spec, table)
			if err != nil {
				return errors.Wrap(err, path)
			}
			app.Tables = append(app.Tables, table)

		case "Resource":
			resource := &corev1Native.Resource{}
			err = mapstructure.Decode(yf.Spec, resource)
			if err != nil {
				return errors.Wrap(err, path)
			}
			app.Resources = append(app.Resources, resource)
		}
		return nil

	})
	if err != nil {
		return
	}

	return app, nil

}
func CheckPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func WriteHybrids(app *Application, baseDir string, logger *zap.Logger) {

	for _, table := range app.Tables {

		trg := hybrids_generator.NewTableReaderGenerator(corev1Native.NewTableReader(table), "github.com/nebtex/omnibuff/pkg/io/omniql/corev1", logger)
		//create file
		f, err := os.Create(baseDir + utils.TableName(corev1Native.NewTableReader(table)) + ".go")
		CheckPanic(err)
		_, err = f.Write([]byte("package corev1Hybrids\n"))
		CheckPanic(err)

		err = trg.Generate(f)
		CheckPanic(err)

		//write struct with hybrids
		//write fields
		for _, f := range table.Fields {
			field := corev1Native.NewFieldReader(f)
			fmt.Println(field.Name(), table.Meta.Name, field.Type())

			if field.Type() == "String" {

			}

		}
		//close
		f.Close()
	}

	for _, resource := range app.Resources {

		trg := hybrids_generator.NewResourceReaderGenerator(corev1Native.NewResourceReader(resource), "github.com/nebtex/omnibuff/pkg/io/omniql/corev1", logger)
		//create file
		f, err := os.Create(baseDir + resource.Meta.Name + ".go")
		CheckPanic(err)
		_, err = f.Write([]byte("package corev1Hybrids\n"))
		CheckPanic(err)

		err = trg.Generate(f)
		CheckPanic(err)

		f.Close()
	}
}


func WriteInterface(app *Application, baseDir string, logger *zap.Logger) {

	for _, table := range app.Tables {

		trg := interface_generator.NewTableReaderGenerator(corev1Native.NewTableReader(table), "github.com/nebtex/omnibuff/pkg/io/omniql/corev1", logger)
		//create file
		f, err := os.Create(baseDir + utils.TableName(corev1Native.NewTableReader(table)) + ".go")
		CheckPanic(err)
		_, err = f.Write([]byte("package corev1Hybrids\n"))
		CheckPanic(err)

		err = trg.Generate(f)
		CheckPanic(err)

		//write struct with hybrids
		//write fields
		for _, f := range table.Fields {
			field := corev1Native.NewFieldReader(f)
			fmt.Println(field.Name(), table.Meta.Name, field.Type())

			if field.Type() == "String" {

			}

		}
		//close
		f.Close()
	}

	for _, resource := range app.Resources {

		trg := interface_generator.NewResourceReaderGenerator(corev1Native.NewResourceReader(resource), "github.com/nebtex/omnibuff/pkg/io/omniql/corev1", logger)
		//create file
		f, err := os.Create(baseDir + resource.Meta.Name + ".go")
		CheckPanic(err)
		_, err = f.Write([]byte("package corev1Hybrids\n"))
		CheckPanic(err)

		err = trg.Generate(f)
		CheckPanic(err)

		f.Close()
	}
}
func main() {
	app, _ := Load("/home/cristian/nebtex/go/src/github.com/nebtex/omnibuff/reflection/omniql/omnibuf")
	baseDir := "/home/cristian/nebtex/go/src/github.com/nebtex/omnibuff/pkg/io/omniql/corev1Hybrids/"

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	undo := zap.RedirectStdLog(logger)
	defer undo()

	WriteHybrids(app, baseDir, logger)

	baseDir = "/home/cristian/nebtex/go/src/github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1/"
	WriteInterface(app, baseDir, logger)

}
