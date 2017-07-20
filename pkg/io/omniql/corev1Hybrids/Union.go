package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//UnionReader ...
type UnionReader struct {
    _table hybrids.TableReader
    _resource hybrids.ResourceReader
    _rid corev1.ResourceIDReader
    meta corev1.MetadataReader
}

//RID get resource id
func (u *UnionReader) RID() corev1.ResourceIDReader {
	return u._rid
}

//Meta ...
func (u *UnionReader) Meta() corev1.MetadataReader {

	if u.meta != nil {
		return u.meta
	}

	return NewMetadataReader(u._table.Table(1))
}

func NewUnionReader(t hybrids.TableReader) corev1.UnionReader{
	if t==nil{
		return nil
	}
	return &UnionReader{_table:t}
}
type VectorUnionReader struct {
    _vectorHybrid    hybrids.VectorTableReader
    _vectorAllocated [] corev1.UnionReader
}

func (vu *VectorUnionReader) Len() (size int) {

    if vu._vectorAllocated != nil {
        size = len(vu._vectorAllocated)
        return
    }

    if vu._vectorHybrid != nil {
        size = vu._vectorHybrid.Len()
        return
    }

    return
}

func (vu *VectorUnionReader) Get(i int) (item corev1.UnionReader, err error) {
    var table hybrids.TableReader

    if vu._vectorAllocated != nil {
        if i < 0 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vu._vectorAllocated)}
            return
        }

        if i > len(vu._vectorAllocated)-1 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vu._vectorAllocated)}
            return
        }

        item = vu._vectorAllocated[i]
        return
    }

    if vu._vectorHybrid != nil {
        table, err = vu._vectorHybrid.Get(i)
        item = NewEnumerationReader(table)
        return
    }

    err = &hybrids.VectorInvalidIndexError{Index: i, Len: 0}
    return
}

func NewVectorUnionReader(v hybrids.VectorTableReader) corev1.VectorUnionReader {
    if v == nil {
        return nil
    }
    return &VectorUnionReader{_vectorHybrid: v}
}
