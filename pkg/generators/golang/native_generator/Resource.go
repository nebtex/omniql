package native_generator

import (
	"io"
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
	"go.uber.org/zap"
	"fmt"
	"github.com/nebtex/omnibuff/pkg/utils"
	"text/template"
)

type ResourceReaderGenerator struct {
	zap      *zap.Logger
	trg      *TableReaderGenerator
	resource corev1.ResourceReader
}

func NewResourceReaderGenerator(resource corev1.ResourceReader, ip string, logger *zap.Logger) (r *ResourceReaderGenerator) {
	zap := logger.With(zap.String("ResourceName", resource.Meta().Name()),
		zap.String("Type", "Reader implementation"),
		zap.String("Application", resource.Meta().Application()),
	)
	r = &ResourceReaderGenerator{zap: zap, resource: resource}
	r.trg = NewTableReaderGenerator(resource, ip, zap)
	return
}
func (r *ResourceReaderGenerator) Generate(wr io.Writer) (err error) {
	//add imports
	wr.Write([]byte(`
import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")


`))
	err = r.trg.StartStruct()
	if err != nil {
		return err
	}

	err = r.trg.StartImplementation()
	if err != nil {
		return err
	}


	err = r.trg.StructAddField("RID", "[]byte")
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

	err = r.trg.EndImplementation()
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
		Funcs(r.trg.funcMap).Parse(`
//RID get resource id
func ({{ShortName}} *{{TableName .Table}}Reader) RID() {{.PackageName}}.ResourceIDReader {
	return {{ShortName}}._rid
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
