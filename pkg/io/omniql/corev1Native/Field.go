package corev1Native

import "github.com/nebtex/omnibuff/pkg/io/omniql/corev1"

type Field struct {
	RID           []byte  `json:"rid"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	Items         string `json:"items"`
	Documentation *Documentation `json:"documentation"`
	Required      bool `json:"required"`
	Default       string `json:"default"`
}

type FieldReader struct {
	field         *Field
	rid           corev1.ResourceIDReader
	documentation corev1.DocumentationReader
}

func (fr *FieldReader) Name() (value string) {

	value = fr.field.Name
	return
}

func (fr *FieldReader) RID() corev1.ResourceIDReader {

	return fr.rid
}

func (fr *FieldReader) Type() (value string) {

	value = fr.field.Type
	return
}

func (fr *FieldReader) Items() (value string) {

	value = fr.field.Items
	return
}

func (fr *FieldReader) Default() (value string) {

	value = fr.field.Default
	return
}

func (fr *FieldReader) Documentation() corev1.DocumentationReader {
	if fr.documentation != nil {
		return fr.documentation
	}
	return NewDocumentationReader(fr.field.Documentation)
}

func (fr *FieldReader) Required() (value bool) {

	value = fr.field.Required
	return
}

func NewFieldReader(f *Field) corev1.FieldReader {
	return &FieldReader{field: f, rid: NewIDReader(f.RID, true)}
}

func NewDeepFieldReader(f *Field) corev1.FieldReader {
	return &FieldReader{field: f, rid: NewIDReader(f.RID, true), documentation: NewDeepDocumentationReader(f.Documentation)}
}

type VectorFieldReader struct {
	nativeFields []*Field
	fields       []corev1.FieldReader
}

func (vf *VectorFieldReader) Len() int32 {
	return int32(len(vf.fields))
}

func (vf *VectorFieldReader) Get(i int32) (item corev1.FieldReader, err error) {
	if vf.nativeFields != nil {
		item = NewFieldReader(vf.nativeFields[i])
		return

	}
	return
}
