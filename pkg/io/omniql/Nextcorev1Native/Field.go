package Nextcorev1Native
import(
    "github.com/nebtex/hybrids/golang/hybrids"
    "github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1"
)
//Field ...
type Field struct {

    Name string `json:"name"`
    Type string `json:"type"`
    Documentation *Documentation `json:"documentation"`
    Default string `json:"default"`
}

//FieldReader ...
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
func (f *FieldReader) Documentation() (dr Nextcorev1.DocumentationReader, err error) {

	if f.documentation != nil {
		dr  =  f.documentation
	}

	return
}

//Default String representation of the default value
func (f *FieldReader) Default() (value string) {
	value = f._field.Default
	return
}

//NewFieldReader ...
func NewFieldReader(f *Field) *FieldReader{
	if f!=nil{
		return &FieldReader{
		                                   _field:f,
documentation: NewDocumentationReader(f.Documentation),
		                                   }
	}
	return nil
}



//VectorFieldReader ...
type VectorFieldReader struct {
    _vector  []*FieldReader
}

//Len Returns the current size of this vector
func (vf *VectorFieldReader) Len() (size int) {
    size = len(vf._vector)
    return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
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

//NewVectorFieldReader ...
func NewVectorFieldReader(vf []*Field) (vfr *VectorFieldReader) {
    vfr = &VectorFieldReader{}
	vfr._vector = make([]*FieldReader, len(vf))

	for i := 0; i < len(vf); i++ {
		vfr._vector[i] = NewFieldReader(vf[i])
	}
	return
}
