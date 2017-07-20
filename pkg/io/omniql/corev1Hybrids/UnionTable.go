package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//UnionTableReader ...
type UnionTableReader struct {
    _table hybrids.TableReader
}

func NewUnionTableReader(t hybrids.TableReader) corev1.UnionTableReader{
	if t==nil{
		return nil
	}
	return &UnionTableReader{_table:t}
}
type VectorUnionTableReader struct {
    _vectorHybrid    hybrids.VectorTableReader
    _vectorAllocated [] corev1.UnionTableReader
}

func (vt *VectorUnionTableReader) Len() (size int) {

    if vt._vectorAllocated != nil {
        size = len(vt._vectorAllocated)
        return
    }

    if vt._vectorHybrid != nil {
        size = vt._vectorHybrid.Len()
        return
    }

    return
}

func (vt *VectorUnionTableReader) Get(i int) (item corev1.UnionTableReader, err error) {
    var table hybrids.TableReader

    if vt._vectorAllocated != nil {
        if i < 0 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vt._vectorAllocated)}
            return
        }

        if i > len(vt._vectorAllocated)-1 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vt._vectorAllocated)}
            return
        }

        item = vt._vectorAllocated[i]
        return
    }

    if vt._vectorHybrid != nil {
        table, err = vt._vectorHybrid.Get(i)
        item = NewEnumerationReader(table)
        return
    }

    err = &hybrids.VectorInvalidIndexError{Index: i, Len: 0}
    return
}

func NewVectorUnionTableReader(v hybrids.VectorTableReader) corev1.VectorUnionTableReader {
    if v == nil {
        return nil
    }
    return &VectorUnionTableReader{_vectorHybrid: v}
}
