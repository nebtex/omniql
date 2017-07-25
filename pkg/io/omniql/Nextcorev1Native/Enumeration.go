package Nextcorev1Native

//Enumeration ...
type Enumeration struct {

    RID []byte `json:"rid"`
    Metadata *Metadata `json:"metadata"`
    Kind string `json:"kind"`
    Items []*EnumerationItem `json:"items"`
    Groups []*EnumerationGroup `json:"groups"`
}
import(
    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
)
//Enumeration ...
type Enumeration struct {

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
func (e *EnumerationReader) RID() corev1.ResourceIDReader {
	return e._rid
}

//Metadata ...
func (e *EnumerationReader) Metadata() corev1.MetadataReader {

	if e.metadata != nil {
		return e.metadata
	}

	return nil
}

//Kind ...
func (e *EnumerationReader) Kind() (value corev1.BasicType) {
	value = corev1.FromStringToBasicType(e._enumeration.Kind)

	if !value.IsScalars(){
		value = corev1.BasicTypeNone
	}

	return
}

//Items ...
func (e *EnumerationReader) Items() corev1.VectorEnumerationItemReader {

	if e.items != nil {
		return e.items
	}

	return nil
}
	
//Groups ...
func (e *EnumerationReader) Groups() corev1.VectorEnumerationGroupReader {

	if e.groups != nil {
		return e.groups
	}

	return nil
}
	
//NewEnumerationReader ...
func NewEnumerationReader(e *EnumerationReader) corev1.EnumerationReader{
	if e!=nil{
		return &EnumerationReader{_enumeration:e}
	}
	return nil
}

type VectorEnumerationReader struct {
    _vector  []*EnumerationReader
}

func (ve *VectorEnumerationReader) Len() (size int) {
    size = len(ve._vector)
    return
}

func (ve *VectorEnumerationReader) Get(i int) (item corev1.EnumerationReader, err error) {

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


func NewVectorEnumerationReader(v hybrids.VectorTableReader) corev1.VectorEnumerationReader {
    if v == nil {
        return nil
    }
    return &VectorEnumerationReader{_vectorHybrid: v}
}
