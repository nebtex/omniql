package Nextcorev1Native
import(
    "github.com/nebtex/hybrids/golang/hybrids"
    "github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1"
)
//EnumerationGroup ...
type EnumerationGroup struct {

    Name string `json:"name"`
    Documentation *Documentation `json:"documentation"`
    Items []string `json:"items"`
}

//EnumerationGroupReader ...
type EnumerationGroupReader struct {
    _enumerationgroup *EnumerationGroup
    documentation *DocumentationReader
    items *native.VectorStringReader
}

//Name ...
func (eg *EnumerationGroupReader) Name() (value string) {
	value = eg._enumerationgroup.Name
	return
}

//Documentation ...
func (eg *EnumerationGroupReader) Documentation() (dr Nextcorev1.DocumentationReader, err error) {

	if eg.documentation != nil {
		dr  =  eg.documentation
	}

	return
}

//Items ...
func (eg *EnumerationGroupReader) Items() hybrids.VectorStringReader {
	if eg.items != nil {
		return eg.items
	}
	return nil
}

//NewEnumerationGroupReader ...
func NewEnumerationGroupReader(eg *EnumerationGroup) *EnumerationGroupReader{
	if eg!=nil{
		return &EnumerationGroupReader{
		                                   _enumerationgroup:eg,
documentation: NewDocumentationReader(g.Documentation),
		                                   }
	}
	return nil
}

//VectorEnumerationGroupReader ...
type VectorEnumerationGroupReader struct {
    _vector  []*EnumerationGroupReader
}

//Len Returns the current size of this vector
func (veg *VectorEnumerationGroupReader) Len() (size int) {
    size = len(veg._vector)
    return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
func (veg *VectorEnumerationGroupReader) Get(i int) (item Nextcorev1.EnumerationGroupReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(veg._vector)}
		return
	}

	if i > len(veg._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(veg._vector)}
		return
	}

	item = veg._vector[i]
	return


}

//NewVectorEnumerationGroupReader ...
func NewVectorEnumerationGroupReader(veg []*EnumerationGroup) (vegr *VectorEnumerationGroupReader) {
    vegr = &VectorEnumerationGroupReader{}
	vegr._vector = make([]*EnumerationGroupReader, len(veg))

	for i := 0; i < len(veg); i++ {
		vegr._vector[i] = NewEnumerationGroupReader(veg[i])
	}
	return
}
