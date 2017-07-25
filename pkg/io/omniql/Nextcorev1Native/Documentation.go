package Nextcorev1Native

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
func NewDocumentationReader(d *DocumentationReader) Nextcorev1.DocumentationReader{
	if d!=nil{
		return &DocumentationReader{_documentation:d}
	}
	return nil
}

type VectorDocumentationReader struct {
    _vector  []*DocumentationReader
}

func (vd *VectorDocumentationReader) Len() (size int) {
    size = len(vd._vector)
    return
}

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


func NewVectorDocumentationReader(v hybrids.VectorTableReader) Nextcorev1.VectorDocumentationReader {
    if v == nil {
        return nil
    }
    return &VectorDocumentationReader{_vectorHybrid: v}
}
