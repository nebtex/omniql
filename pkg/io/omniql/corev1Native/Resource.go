package corev1Native

import "github.com/nebtex/omnibuff/pkg/io/omniql/corev1"

type Resource struct {
	RID    []byte  `json:"rid"`
	Meta   *Metadata `json:"meta"`
	Fields []*Field `json:"fields"`
}

type ResourceReader struct {
	resource *Resource
	meta     corev1.MetadataReader
	doc      corev1.DocumentationReader
}

func (t *ResourceReader) RID() corev1.ResourceIDReader {
	return NewIDReader(t.resource.RID, true)
}

func (t *ResourceReader) Meta() corev1.MetadataReader {
	return &MetadataReader{meta: t.resource.Meta}

}

func (t *ResourceReader) Fields() corev1.VectorFieldReader {
	return &VectorFieldReader{nativeFields: t.resource.Fields}

}

func NewResourceReader(r *Resource) corev1.ResourceReader {
	return &ResourceReader{resource: r}

}
