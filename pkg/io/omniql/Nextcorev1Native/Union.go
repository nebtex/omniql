package Nextcorev1Native

//Union ...
type Union struct {

    RID []byte `json:"rid"`
    Metadata *Metadata `json:"metadata"`
}

//Union ...
type Union struct {

    Metadata *Metadata `json:"metadata"`
}

//UnionReader ...
type UnionReader struct {
    _union *Union
    metadata *MetadataReader
}

//RID get resource id
func (u *UnionReader) RID() corev1.ResourceIDReader {
	return u._rid
}

//Metadata ...
func (u *UnionReader) Metadata() corev1.MetadataReader {

	if u.metadata != nil {
		return u.metadata
	}

	return nil
}

//NewUnionReader ...
func NewUnionReader(u *UnionReader) corev1.UnionReader{
	if u!=nil{
		return &UnionReader{_union:u}
	}
	return nil
}

type VectorUnionReader struct {
    _vector  []*UnionReader
}

func (vu *VectorUnionReader) Len() (size int) {
    size = len(vu._vector)
    return
}

func (vu *VectorUnionReader) Get(i int) (item corev1.UnionReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vu._vector)}
		return
	}

	if i > len(vu._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vu._vector)}
		return
	}

	item = vu._vector[i]
	return


}


func NewVectorUnionReader(v hybrids.VectorTableReader) corev1.VectorUnionReader {
    if v == nil {
        return nil
    }
    return &VectorUnionReader{_vectorHybrid: v}
}
