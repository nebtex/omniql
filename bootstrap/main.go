package main

import (
	"io/ioutil"
	"path/filepath"
	"github.com/ghodss/yaml"
	"github.com/mitchellh/mapstructure"
	"os"
	"github.com/jjeffery/errors"
	"strings"
	"fmt"
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1Native"
	"github.com/nebtex/omnibuff/pkg/generators/golang/hybrids_generator"
	"go.uber.org/zap"
	"github.com/nebtex/omnibuff/pkg/utils"
	"github.com/nebtex/omnibuff/pkg/generators/golang/interface_generator"
	"github.com/nebtex/omnibuff/pkg/generators/golang/native_generator"
)

type YamlFile struct {
	Api   string `json:"api"`
	RID string `json:"rid"`
	Spec  interface{} `json:"spec"`
}

type Application struct {
	Name      string `json:"name"`
	Resources []*corev1Native.Resource `json:"resources"`
	Tables    []*corev1Native.Table `json:"tables"`
	//Unions       []*UnionSpec `json:"tables"`
	Enumerations []*corev1Native.Enumeration  `json:"enumerations"`
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
		types := strings.Split(yf.RID, "/")
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

		case "Enumeration":
			enum := &corev1Native.Enumeration{}
			err = mapstructure.Decode(yf.Spec, enum)
			if err != nil {
				return errors.Wrap(err, path)
			}
			app.Enumerations = append(app.Enumerations, enum)
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

func WriteNative(app *Application, baseDir string, packageName string, logger *zap.Logger) {

	for _, table := range app.Tables {

		trg := native_generator.NewTableReaderGenerator(corev1Native.NewTableReader(table), "github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1", logger)
		//create file
		f, err := os.Create(baseDir + utils.TableName(corev1Native.NewTableReader(table)) + ".go")
		CheckPanic(err)
		_, err = f.Write([]byte("package " + packageName + "\n"))
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

		trg := native_generator.NewResourceReaderGenerator(corev1Native.NewResourceReader(resource), "github.com/nebtex/omnibuff/pkg/io/omniql/corev1", logger)
		//create file
		f, err := os.Create(baseDir + resource.Meta.Name + ".go")
		CheckPanic(err)
		_, err = f.Write([]byte("package " + packageName + "\n"))
		CheckPanic(err)

		err = trg.Generate(f)
		CheckPanic(err)

		f.Close()
	}
}

func WriteInterface(app *Application, baseDir string, packageName string, logger *zap.Logger) {

	for _, table := range app.Tables {

		trg := interface_generator.NewTableReaderGenerator(corev1Native.NewTableReader(table), "github.com/nebtex/omnibuff/pkg/io/omniql/corev1", logger)
		//create file
		f, err := os.Create(baseDir + utils.TableName(corev1Native.NewTableReader(table)) + ".go")
		CheckPanic(err)
		_, err = f.Write([]byte("package " + packageName + "\n"))
		CheckPanic(err)

		err = trg.Generate(f)
		CheckPanic(err)

		//close
		f.Close()
	}

	for _, resource := range app.Resources {

		trg := interface_generator.NewResourceReaderGenerator(corev1Native.NewResourceReader(resource), "github.com/nebtex/omnibuff/pkg/io/omniql/corev1", logger)
		//create file
		f, err := os.Create(baseDir + resource.Meta.Name + ".go")
		CheckPanic(err)
		_, err = f.Write([]byte("package " + packageName + "\n"))
		CheckPanic(err)

		err = trg.Generate(f)
		CheckPanic(err)

		f.Close()
	}

	for _, enumeration := range app.Enumerations {

		trg := interface_generator.NewEnumerationGenerator(corev1Native.NewEnumerationReader(enumeration), logger)
		//create file
		f, err := os.Create(baseDir + enumeration.Meta.Name + ".go")
		CheckPanic(err)
		_, err = f.Write([]byte("package " + packageName + "\n"))
		CheckPanic(err)

		err = trg.Generate(f)
		CheckPanic(err)

		f.Close()
	}
}

func main() {
	app, err := Load("/home/cristian/nebtex/go/src/github.com/nebtex/omnibuff/reflection/omniql/omnibuf")
	CheckPanic(err)
/*
	//generate application types enum
	appEnum := `
api: io.omniql.core.v1
oqlid: io.omniql.core.v1/Enumeration/BasicTypes
spec:
  meta:
    application: io.omniql.core.v1
    kind: Enumeration
    name: ApplicationType
  kind: UnsignedShort
  items:`
	for _, r := range app.Resources {
		appEnum += "\n  - name: " + r.Meta.Name
	}
	for _, t := range app.Tables {
		appEnum += "\n  - name: " +   utils.TableName(corev1Native.NewTableReader(t))
	}
	ioutil.WriteFile("/home/cristian/nebtex/go/src/github.com/nebtex/omnibuff/reflection/omniql/omnibuf/enumeration/ApplicationType.yml", []byte(appEnum), os.ModePerm)
	app, _ = Load("/home/cristian/nebtex/go/src/github.com/nebtex/omnibuff/reflection/omniql/omnibuf")*/
	baseDir := "/home/cristian/nebtex/go/src/github.com/nebtex/omnibuff/pkg/io/omniql/corev1Hybrids/"

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	undo := zap.RedirectStdLog(logger)
	defer undo()

	//WriteHybrids(app, baseDir, logger)

	baseDir = "/home/cristian/nebtex/go/src/github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1/"
	WriteInterface(app, baseDir, "Nextcorev1", logger)

	//baseDir = "/home/cristian/nebtex/go/src/github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1Native/"
	//WriteNative(app, baseDir, "Nextcorev1Native", logger)

}
