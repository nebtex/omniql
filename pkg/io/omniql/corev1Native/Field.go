package corev1Native

import "github.com/nebtex/omniql/pkg/io/omniql/corev1"

type Field struct {
	RID           []byte `json:"rid"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	Items         string `json:"items"`
	Documentation *Documentation `json:"documentation"`
	Required      bool `json:"required"`
	Default       string `json:"default"`
}

type FieldReader struct {
	field         *Field
	rid           *IDReader
	documentation *DocumentationReader
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
	return nil
}

func (fr *FieldReader) Required() (value bool) {

	value = fr.field.Required
	return
}

func NewFieldReader(f *Field) corev1.FieldReader {
	return &FieldReader{
		field:         f,
		rid:           NewIDReader(f.RID, true),
		documentation: NewDocumentationReader(f.Documentation),
	}
}

type VectorFieldReader struct {
	nativeFields []*Field
	fields       []corev1.FieldReader
}

func (vf *VectorFieldReader) Len() int {
	if vf.nativeFields != nil {

		return int(len(vf.nativeFields))
	}
	return int(len(vf.fields))

}

func (vf *VectorFieldReader) Get(i int) (item corev1.FieldReader, err error) {
	if vf.nativeFields != nil {
		item = NewFieldReader(vf.nativeFields[i])
		return

	}
	return
}
