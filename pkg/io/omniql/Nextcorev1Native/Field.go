package Nextcorev1Native

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
	"github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1"
)



//Field field type
type Field struct {

    Name string `json:"name"`
    Type string `json:"type"`
    Documentation *Documentation `json:"documentation"`
    Default string `json:"default"`
}

//FieldReader field type
type FieldReader struct {
    _field *Field
    documentation *DocumentationReader
}

//Name ...
func (f *FieldReader) Name() (value string) {
	value = f._field.Name
	return
}

//Type ...
func (f *FieldReader) Type() (value string) {
	value = f._field.Type
	return
}

//Documentation ...
func (f *FieldReader) Documentation() Nextcorev1.DocumentationReader {

	if f.documentation != nil {
		return f.documentation
	}

	return nil
}

//Default String representation of the default value
func (f *FieldReader) Default() (value string) {
	value = f._field.Default
	return
}

func NewFieldReader(t hybrids.TableReader) Nextcorev1.FieldReader{
	if t==nil{
		return nil
	}
	return &FieldReader{_table:t}
}

type VectorFieldReader struct {
    _vector  []*FieldReader
}

func (vf *VectorFieldReader) Len() (size int) {
    size = len(vf._vector)
    return
}

func (vf *VectorFieldReader) Get(i int) (item Nextcorev1.FieldReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vf._vector)}
		return
	}

	if i > len(vf._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vf._vector)}
		return
	}

	item = vf._vector[i]
	return


}

func NewVectorFieldReader(v hybrids.VectorTableReader) Nextcorev1.VectorFieldReader {
    if v == nil {
        return nil
    }
    return &VectorFieldReader{_vectorHybrid: v}
}