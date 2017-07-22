package Nextcorev1Native

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//UnionResource ...
type UnionResource struct {

}

//UnionResourceReader ...
type UnionResourceReader struct {
    _unionresource *UnionResource
}

func NewUnionResourceReader(t hybrids.TableReader) Nextcorev1.UnionResourceReader{
	if t==nil{
		return nil
	}
	return &UnionResourceReader{_table:t}
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
