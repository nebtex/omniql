package Nextcorev1Native
import(
    "github.com/nebtex/hybrids/golang/hybrids"
    "github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1"
)
//EnumerationItem ...
type EnumerationItem struct {

    Name string `json:"name"`
    Documentation *Documentation `json:"documentation"`
}

//EnumerationItemReader ...
type EnumerationItemReader struct {
    _enumerationitem *EnumerationItem
    documentation *DocumentationReader
}

//Name ...
func (ei *EnumerationItemReader) Name() (value string) {
	value = ei._enumerationitem.Name
	return
}

//Documentation ...
func (ei *EnumerationItemReader) Documentation() (dr Nextcorev1.DocumentationReader, err error) {

	if ei.documentation != nil {
		dr  =  ei.documentation
	}

	return
}

//NewEnumerationItemReader ...
func NewEnumerationItemReader(ei *EnumerationItem) *EnumerationItemReader{
	if ei!=nil{
		return &EnumerationItemReader{
		                                   _enumerationitem:ei,
documentation: NewDocumentationReader(i.Documentation),
		                                   }
	}
	return nil
}

//VectorEnumerationItemReader ...
type VectorEnumerationItemReader struct {
    _vector  []*EnumerationItemReader
}

//Len Returns the current size of this vector
func (vei *VectorEnumerationItemReader) Len() (size int) {
    size = len(vei._vector)
    return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
func (vei *VectorEnumerationItemReader) Get(i int) (item Nextcorev1.EnumerationItemReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vei._vector)}
		return
	}

	if i > len(vei._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vei._vector)}
		return
	}

	item = vei._vector[i]
	return


}

//NewVectorEnumerationItemReader ...
func NewVectorEnumerationItemReader(vei []*EnumerationItem) (veir *VectorEnumerationItemReader) {
    veir = &VectorEnumerationItemReader{}
	veir._vector = make([]*EnumerationItemReader, len(vei))

	for i := 0; i < len(vei); i++ {
		veir._vector[i] = NewEnumerationItemReader(vei[i])
	}
	return
}
