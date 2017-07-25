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
	metadata *Metadata
	doc  *DocumentationReader
}

func (m *MetadataReader) Application() string {
	return m.metadata.Application
}
func (m *MetadataReader) Name() string {
	return m.metadata.Name
}
func (m *MetadataReader) Kind() string {
	return m.metadata.Kind
}

func (m *MetadataReader) Parent() string {
	return m.metadata.Parent
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
	mr := &MetadataReader{metadata: m, doc: NewDocumentationReader(m.Documentation)}
	return mr

}
