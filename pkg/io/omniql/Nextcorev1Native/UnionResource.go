package Nextcorev1Native

//UnionResource ...
type UnionResource struct {

}

//UnionResourceReader ...
type UnionResourceReader struct {
    _unionresource *UnionResource
}

//NewUnionResourceReader ...
func NewUnionResourceReader(r *UnionResourceReader) Nextcorev1.UnionResourceReader{
	if r!=nil{
		return &UnionResourceReader{_resource:r}
	}
	return nil
}

type VectorUnionResourceReader struct {
    _vector  []*UnionResourceReader
}

func (vr *VectorUnionResourceReader) Len() (size int) {
    size = len(vr._vector)
    return
}

func (vr *VectorUnionResourceReader) Get(i int) (item Nextcorev1.UnionResourceReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vr._vector)}
		return
	}

	if i > len(vr._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vr._vector)}
		return
	}

	item = vr._vector[i]
	return


}


func NewVectorUnionResourceReader(v hybrids.VectorTableReader) Nextcorev1.VectorUnionResourceReader {
    if v == nil {
        return nil
    }
    return &VectorUnionResourceReader{_vectorHybrid: v}
}
