package Nextcorev1Native
import(
    "github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1"
    "github.com/nebtex/hybrids/golang/hybrids"
)
//Enumeration ...
type Enumeration struct {

    RID []byte `json:"rid"`
    Metadata *Metadata `json:"metadata"`
    Kind string `json:"kind"`
    Items []*EnumerationItem `json:"items"`
    Groups []*EnumerationGroup `json:"groups"`
}

//EnumerationReader ...
type EnumerationReader struct {
    _enumeration *Enumeration
    metadata *MetadataReader
    items *VectorEnumerationItemReader
    groups *VectorEnumerationGroupReader
}

//RID get resource id
func (e *EnumerationReader) RID() Nextcorev1.ResourceIDReader {
	return e._rid
}



//Metadata ...
func (e *EnumerationReader) Metadata() (mr Nextcorev1.MetadataReader, err error) {

	if e.metadata != nil {
		mr  =  e.metadata
	}

	return
}

//Kind ...
func (e *EnumerationReader) Kind() (value Nextcorev1.BasicType) {
	value = Nextcorev1.FromStringToBasicType(e._enumeration.Kind)

	if !value.IsScalar(){
		value = Nextcorev1.BasicTypeNone
	}

	return
}

//Items ...
func (e *EnumerationReader) Items() Nextcorev1.VectorEnumerationItemReader {

	if e.items != nil {
		return e.items
	}

	return nil
}
	
//Groups ...
func (e *EnumerationReader) Groups() Nextcorev1.VectorEnumerationGroupReader {

	if e.groups != nil {
		return e.groups
	}

	return nil
}
	
//NewEnumerationReader ...
func NewEnumerationReader(e *Enumeration) *EnumerationReader{
	if e!=nil{
		return &EnumerationReader{
		                                   _enumeration:e,
metadata: NewMetadataReader(e.Metadata),
		                                   }
	}
	return nil
}



//VectorEnumerationReader ...
type VectorEnumerationReader struct {
    _vector  []*EnumerationReader
}

//Len Returns the current size of this vector
func (ve *VectorEnumerationReader) Len() (size int) {
    size = len(ve._vector)
    return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
func (ve *VectorEnumerationReader) Get(i int) (item Nextcorev1.EnumerationReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(ve._vector)}
		return
	}

	if i > len(ve._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(ve._vector)}
		return
	}

	item = ve._vector[i]
	return


}

//NewVectorEnumerationReader ...
func NewVectorEnumerationReader(ve []*Enumeration) (ver *VectorEnumerationReader) {
    ver = &VectorEnumerationReader{}
	ver._vector = make([]*EnumerationReader, len(ve))

	for i := 0; i < len(ve); i++ {
		ver._vector[i] = NewEnumerationReader(ve[i])
	}
	return
}
