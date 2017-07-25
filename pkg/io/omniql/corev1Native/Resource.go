package corev1Native

import "github.com/nebtex/omnibuff/pkg/io/omniql/corev1"

type Resource struct {
	RID    []byte    `json:"rid"`
	Metadata   *Metadata `json:"meta"`
	Fields []*Field  `json:"fields"`
}

type ResourceReader struct {
	resource *Resource
	metadata     corev1.MetadataReader
	doc      corev1.DocumentationReader
}

func (t *ResourceReader) RID() corev1.ResourceIDReader {
	return NewIDReader(t.resource.RID, true)
}

func (t *ResourceReader) Metadata() corev1.MetadataReader {
	return &MetadataReader{metadata: t.resource.Metadata}

}

func (t *ResourceReader) Fields() corev1.VectorFieldReader {
	return &VectorFieldReader{nativeFields: t.resource.Fields}

}

func NewResourceReader(r *Resource) corev1.ResourceReader {
	return &ResourceReader{resource: r}

}
