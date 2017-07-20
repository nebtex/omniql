package hybrids_generator

import (
	"io"
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
	"go.uber.org/zap"
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
	err = r.trg.Generate(wr)
	return

}
