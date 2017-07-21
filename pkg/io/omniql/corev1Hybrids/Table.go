package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//TableReader ...
type TableReader struct {
    _table hybrids.TableReader
    _resource hybrids.ResourceReader
    _rid corev1.ResourceIDReader
    meta *MetadataReader
}

//RID get resource id
func (t *TableReader) RID() corev1.ResourceIDReader {
	return t._rid
}

//Meta ...
func (t *TableReader) Meta() corev1.MetadataReader {

	if t.meta != nil {
		return t.meta
	}

	return NewMetadataReader(t._table.Table(1))
}

func NewTableReader(t hybrids.TableReader) corev1.TableReader{
	if t==nil{
		return nil
	}
	return &TableReader{_table:t}
}
type VectorTableReader struct {
    _vectorHybrid    hybrids.VectorTableReader
    _vectorAllocated [] *TableReader
}

func (vt *VectorTableReader) Len() (size int) {

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

func (vt *VectorTableReader) Get(i int) (item corev1.TableReader, err error) {
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

func NewVectorTableReader(v hybrids.VectorTableReader) corev1.VectorTableReader {
    if v == nil {
        return nil
    }
    return &VectorTableReader{_vectorHybrid: v}
}
