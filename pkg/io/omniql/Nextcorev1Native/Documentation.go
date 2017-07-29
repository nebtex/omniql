package Nextcorev1Native
import(
    "github.com/nebtex/hybrids/golang/hybrids"
    "github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1"
)
//Documentation ...
type Documentation struct {

    Short string `json:"short"`
    Long string `json:"long"`
}

//DocumentationReader ...
type DocumentationReader struct {
    _documentation *Documentation
}

//Short ...
func (d *DocumentationReader) Short() (value string) {
	value = d._documentation.Short
	return
}

//Long ...
func (d *DocumentationReader) Long() (value string) {
	value = d._documentation.Long
	return
}

//NewDocumentationReader ...
func NewDocumentationReader(d *Documentation) *DocumentationReader{
	if d!=nil{
		return &DocumentationReader{
		                                   _documentation:d,
		                                   }
	}
	return nil
}



//VectorDocumentationReader ...
type VectorDocumentationReader struct {
    _vector  []*DocumentationReader
}

//Len Returns the current size of this vector
func (vd *VectorDocumentationReader) Len() (size int) {
    size = len(vd._vector)
    return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
func (vd *VectorDocumentationReader) Get(i int) (item Nextcorev1.DocumentationReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vd._vector)}
		return
	}

	if i > len(vd._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vd._vector)}
		return
	}

	item = vd._vector[i]
	return


}

//NewVectorDocumentationReader ...
func NewVectorDocumentationReader(vd []*Documentation) (vdr *VectorDocumentationReader) {
    vdr = &VectorDocumentationReader{}
	vdr._vector = make([]*DocumentationReader, len(vd))

	for i := 0; i < len(vd); i++ {
		vdr._vector[i] = NewDocumentationReader(vd[i])
	}
	return
}
