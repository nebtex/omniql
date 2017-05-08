package schema

import (
	"io/ioutil"
	"path/filepath"
	"github.com/ghodss/yaml"
	"github.com/mitchellh/mapstructure"
	"os"
	"github.com/jjeffery/errors"
	"text/template"
)

var OM map[string]*Meta

func init()  {
	OM = map[string]*Meta{}

}
type Documentation struct {
	Short string `json:"short"`
	Long  string `json:"long"`
}

type Meta struct {
	Application   string `json:"application"`
	Name          string `json:"name"`
	Resource      string `json:"resource"`
	Documentation *Documentation `json:"documentation"`
}

type Field struct {
	Name          string `json:"name"`
	Schema        string `json:"schema"`
	Items         string `json:"items"`
	Documentation *Documentation `json:"documentation"`
	Required      bool `json:"required"`
}

type TableSpec struct {
	Meta   *Meta `json:"meta"`
	Fields []*Field `json:"fields"`
}

func (t *TableSpec) ToFlatBuffer() (string, error) {
	funcMap := template.FuncMap{
		"tableName": func(meta *Meta) string {
			if meta.Resource == "" {
				return meta.Name
			}
			return meta.Resource + meta.Name
		},
		"toFbsSchema": func(field *Field) string {
			switch field.Schema {
			case "io.omniql.core.v1/Integer":
				return "int"
			case "io.omniql.core.v1/String":
				return "string"
			case "io.omniql.core.v1/Boolean":
				return "bool"
			case "io.omniql.core.v1/Vector":
				switch field.Items {
				case "io.omniql.core.v1/String":
					return "[string]"
				default:
					panic("Vector not found: " + field.Items)
				}
			default:
				panic("Schema not found: " + field.Schema)
			}
		},
	}

	tmpl, err := template.New("table").Funcs(funcMap).Parse(`
table {{ tableName .Meta }} {
{{range .Fields}}    {{.Name}}:{{toFbsSchema .}};
{{end}}}
`)
	err = tmpl.Execute(os.Stdout, t)

	if err != nil {
		return "", err
	}

	return "", nil
}

type UnionSpec struct {
	Meta    *Meta `json:"meta"`
	Schemas []string `json:"schemas"`
}

type ResourceSpec struct {
	Meta   *Meta `json:"meta"`
	Fields []*Field `json:"fields"`
}

type YamlFile struct {
	Api  string `json:"api"`
	Kind string `json:"kind"`
	Spec interface{} `json:"spec"`
}

type Application struct {
	Name      string `json:"name"`
	Resources []*ResourceSpec `json:"resources"`
	Tables    []*TableSpec `json:"tables"`
	Unions    []*UnionSpec `json:"tables"`
}

func LoadResource(path string) (app *Application, err error) {
	app = &Application{}
	// load resource
	data, err := ioutil.ReadFile(filepath.Join(path, "index.yml"))
	if err != nil {
		return
	}
	yf := &YamlFile{}
	yaml.Unmarshal(data, yf)
	app.Name = yf.Api
	resource := &ResourceSpec{}
	err = mapstructure.Decode(yf.Spec, resource)
	if err != nil {
		return
	}

	//load tables
	err = filepath.Walk(filepath.Join(path, "tables"), func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		yf := &YamlFile{}
		yaml.Unmarshal(data, yf)
		table := &TableSpec{}
		err = mapstructure.Decode(yf.Spec, table)
		if err != nil {
			return errors.Wrap(err, path)
		}
		app.Tables = append(app.Tables, table)
		return nil
	})

	//load unions
	err = filepath.Walk(filepath.Join(path, "unions"), func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		yf := &YamlFile{}
		yaml.Unmarshal(data, yf)
		union := &UnionSpec{}
		err = mapstructure.Decode(yf.Spec, union)
		if err != nil {
			return errors.Wrap(err, path)
		}
		app.Unions = append(app.Unions, union)
		return nil
	})

	if err != nil {
		return
	}

	// load union

	return

}
