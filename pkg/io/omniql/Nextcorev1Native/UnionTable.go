package Nextcorev1Native

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//UnionTable ...
type UnionTable struct {

}

//UnionTableReader ...
type UnionTableReader struct {
    _uniontable *UnionTable
}

func NewUnionTableReader(t hybrids.TableReader) Nextcorev1.UnionTableReader{
	if t==nil{
		return nil
	}
	return &UnionTableReader{_table:t}
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
