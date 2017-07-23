package corev1Native

import (
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
)

type Metadata struct {
	Application   string `json:"application"`
	Name          string `json:"name"`
	Kind          string `json:"kind"`
	Parent        string `json:"parent"`
	Resource      string `json:"resource"`
	Documentation *Documentation `json:"documentation"`
}

type MetadataReader struct {
	meta *Metadata
	doc  *DocumentationReader
}

func (m *MetadataReader) Application() string {
	return m.meta.Application
}
func (m *MetadataReader) Name() string {
	return m.meta.Name
}
func (m *MetadataReader) Kind() string {
	return m.meta.Kind
}

func (m *MetadataReader) Parent() string {
	return m.meta.Parent
}
func (m *MetadataReader) Documentation() corev1.DocumentationReader {
	if m.doc != nil {
		return m.doc
	}
	return nil
}

func NewMetadataReader(m *Metadata) *MetadataReader {
	if m == nil {
		return nil
	}
	mr := &MetadataReader{meta: m, doc: NewDocumentationReader(m.Documentation)}
	return mr

}
