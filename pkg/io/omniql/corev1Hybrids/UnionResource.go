package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//UnionResourceReader ...
type UnionResourceReader struct {
    _table hybrids.TableReader
}

func NewUnionResourceReader(t hybrids.TableReader) corev1.UnionResourceReader{
	if t==nil{
		return nil
	}
	return &UnionResourceReader{_table:t}
}
type VectorUnionResourceReader struct {
    _vectorHybrid    hybrids.VectorTableReader
    _vectorAllocated [] *UnionResourceReader
}

func (vr *VectorUnionResourceReader) Len() (size int) {

    if vr._vectorAllocated != nil {
        size = len(vr._vectorAllocated)
        return
    }

    if vr._vectorHybrid != nil {
        size = vr._vectorHybrid.Len()
        return
    }

    return
}

func (vr *VectorUnionResourceReader) Get(i int) (item corev1.UnionResourceReader, err error) {
    var table hybrids.TableReader

    if vr._vectorAllocated != nil {
        if i < 0 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vr._vectorAllocated)}
            return
        }

        if i > len(vr._vectorAllocated)-1 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vr._vectorAllocated)}
            return
        }

        item = vr._vectorAllocated[i]
        return
    }

    if vr._vectorHybrid != nil {
        table, err = vr._vectorHybrid.Get(i)
        item = NewEnumerationReader(table)
        return
    }

    err = &hybrids.VectorInvalidIndexError{Index: i, Len: 0}
    return
}

func NewVectorUnionResourceReader(v hybrids.VectorTableReader) corev1.VectorUnionResourceReader {
    if v == nil {
        return nil
    }
    return &VectorUnionResourceReader{_vectorHybrid: v}
}
