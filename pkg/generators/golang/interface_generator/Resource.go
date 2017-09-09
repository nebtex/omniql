package interface_generator

import (
	"io"
	"github.com/nebtex/omniql/pkg/io/omniql/corev1"
	"go.uber.org/zap"
	"fmt"
	"github.com/nebtex/omniql/pkg/utils"
	"text/template"
)

type ResourceReaderGenerator struct {
	zap      *zap.Logger
	trg      *TableReaderGenerator
	resource corev1.ResourceReader
}

func NewResourceReaderGenerator(resource corev1.ResourceReader, ip string, logger *zap.Logger) (r *ResourceReaderGenerator) {
	zap := logger.With(zap.String("ResourceName", resource.Metadata().Name()),
		zap.String("Type", "Reader implementation"),
		zap.String("Application", resource.Metadata().Application()),
	)
	r = &ResourceReaderGenerator{zap: zap, resource: resource}
	r.trg = NewTableReaderGenerator(resource, ip, zap)
	return
}
func (r *ResourceReaderGenerator) Generate(wr io.Writer) (err error) {
	//add imports
	wr.Write([]byte(`

`))
	err = r.trg.StartInterface()
	if err != nil {
		return err
	}

	err = r.CreateRIDAccessor()
	if err != nil {
		return err
	}

	err = r.trg.CreateAccessors()
	if err != nil {
		return err
	}

	err = r.trg.CreateVector()
	if err != nil {
		return err
	}
	err = r.trg.EndInterface()
	if err != nil {
		return err
	}

	err = r.trg.FlushBuffers(wr)
	if err != nil {
		return err
	}

	r.zap.Info(fmt.Sprintf("Interface for Resource %s Created successfully", utils.TableName(r.resource)))
	return

}

func (r *ResourceReaderGenerator) CreateRIDAccessor() (err error) {
	tmpl, err := template.New("ResourceReaderGenerator::CreateRIDAccessor").
		Funcs(r.trg.funcMap).Parse(`
    //RID get resource id
    RID() ResourceIDReader
`)
	if err != nil {
		return
	}

	err = tmpl.Execute(r.trg.definitionsBuffer, map[string]interface{}{
		"Table":       r.resource,
		"PackageName": r.trg.interfacePackageShort},
	)
	if err != nil {
		return
	}

	return

}
