package Nextcorev1Native

//UnionTable ...
type UnionTable struct {

}

//UnionTableReader ...
type UnionTableReader struct {
    _uniontable *UnionTable
}

//NewUnionTableReader ...
func NewUnionTableReader(t *UnionTableReader) Nextcorev1.UnionTableReader{
	if t!=nil{
		return &UnionTableReader{_table:t}
	}
	return nil
}

type VectorUnionTableReader struct {
    _vector  []*UnionTableReader
}

func (vt *VectorUnionTableReader) Len() (size int) {
    size = len(vt._vector)
    return
}

func (vt *VectorUnionTableReader) Get(i int) (item Nextcorev1.UnionTableReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vt._vector)}
		return
	}

	if i > len(vt._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vt._vector)}
		return
	}

	item = vt._vector[i]
	return


}


func NewVectorUnionTableReader(v hybrids.VectorTableReader) Nextcorev1.VectorUnionTableReader {
    if v == nil {
        return nil
    }
    return &VectorUnionTableReader{_vectorHybrid: v}
}
