package Nextcorev1Native

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//Enumeration ...
type Enumeration struct {

    RID []byte `json:"rid"`
    Meta *Metadata `json:"meta"`
    Items []*EnumerationItem `json:"items"`
    Groups []*EnumerationGroup `json:"groups"`
}

//EnumerationReader ...
type EnumerationReader struct {
    _enumeration *Enumeration
    meta *MetadataReader
    items *VectorEnumerationItemReader
    groups *VectorEnumerationGroupReader
}

//RID get resource id
func (e *EnumerationReader) RID() corev1.ResourceIDReader {
	return e._rid
}

//Meta ...
func (e *EnumerationReader) Meta() corev1.MetadataReader {

	if e.meta != nil {
		return e.meta
	}

	return nil
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
	
func NewEnumerationReader(t hybrids.TableReader) corev1.EnumerationReader{
	if t==nil{
		return nil
	}
	return &EnumerationReader{_table:t}
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
