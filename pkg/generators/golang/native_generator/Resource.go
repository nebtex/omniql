package native_generator

import (
	"io"
	"github.com/nebtex/omniql/pkg/io/omniql/corev1"
	"go.uber.org/zap"
	"fmt"
	"github.com/nebtex/omniql/pkg/utils"
	"text/template"
	"github.com/nebtex/omniql/pkg/generators/golang"
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
	err = r.trg.gt.StructAddField("RID", "[]byte")
	if err != nil {
		return err
	}


	//add imports
	err = r.trg.StartStruct()
	if err != nil {
		return err
	}



	err = r.CreateRIDAccessor()
	if err != nil {
		return err
	}

	err = r.trg.CreateAccessors(1)
	if err != nil {
		return err
	}

	err = r.trg.CreateAllocator()
	if err != nil {
		return err
	}

	err = r.trg.CreateVector()
	if err != nil {
		return err
	}


	err = r.trg.EndStruct()
	if err != nil {
		return err
	}

	err = r.trg.FlushBuffers(wr)
	if err != nil {
		return err
	}

	r.zap.Info(fmt.Sprintf("Resource %s Created successfully", utils.TableName(r.resource)))
	return

}

func (r *ResourceReaderGenerator) CreateRIDAccessor() (err error) {
	tmpl, err := template.New("ResourceReaderGenerator::CreateRIDAccessor").
		Funcs(golang.DefaultTemplateFunctions).Parse(`
//RID get resource id
func ({{ShortName (TableName .Table)}} *{{TableName .Table}}Reader) RID() {{.PackageName}}.ResourceIDReader {
	return {{ShortName (TableName .Table)}}._rid
}
`)
	if err != nil {
		return
	}

	err = tmpl.Execute(r.trg.functionsBuffer, map[string]interface{}{
		"Table":       r.resource,
		"PackageName": r.trg.interfacePackageShort},
	)
	if err != nil {
		return
	}

	return

}
