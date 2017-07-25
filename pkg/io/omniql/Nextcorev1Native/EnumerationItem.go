package Nextcorev1Native

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
func (i *EnumerationItemReader) Name() (value string) {
	value = i._enumerationitem.Name
	return
}

//Documentation ...
func (i *EnumerationItemReader) Documentation() Nextcorev1.DocumentationReader {

	if i.documentation != nil {
		return i.documentation
	}

	return nil
}

//NewEnumerationItemReader ...
func NewEnumerationItemReader(i *EnumerationItemReader) Nextcorev1.EnumerationItemReader{
	if i!=nil{
		return &EnumerationItemReader{_item:i}
	}
	return nil
}

type VectorEnumerationItemReader struct {
    _vector  []*EnumerationItemReader
}

func (vi *VectorEnumerationItemReader) Len() (size int) {
    size = len(vi._vector)
    return
}

func (vi *VectorEnumerationItemReader) Get(i int) (item Nextcorev1.EnumerationItemReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vi._vector)}
		return
	}

	if i > len(vi._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vi._vector)}
		return
	}

	item = vi._vector[i]
	return


}


func NewVectorEnumerationItemReader(v hybrids.VectorTableReader) Nextcorev1.VectorEnumerationItemReader {
    if v == nil {
        return nil
    }
    return &VectorEnumerationItemReader{_vectorHybrid: v}
}
