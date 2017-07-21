package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//EnumerationReader ...
type EnumerationReader struct {
    _table hybrids.TableReader
    _resource hybrids.ResourceReader
    _rid corev1.ResourceIDReader
    meta *MetadataReader
    items *VectorEnumerationItemReader
    groups *VectorEnumerationGroupReader
}

//RID get resource id
func (e *EnumerationReader) RID() corev1.ResourceIDReader {
	return e._rid
}

//Meta ...
func (e *EnumerationReader) Meta() corev1.MetadataReader {

	if e.meta != nil {
		return e.meta
	}

	return NewMetadataReader(e._table.Table(1))
}

//Items ...
func (e *EnumerationReader) Items() corev1.VectorEnumerationItemReader {

	if e.items != nil {
		return e.items
	}

	return NewVectorEnumerationItemReader(e._table.VectorTable(3))
}
	
//Groups ...
func (e *EnumerationReader) Groups() corev1.VectorEnumerationGroupReader {

	if e.groups != nil {
		return e.groups
	}

	return NewVectorEnumerationGroupReader(e._table.VectorTable(4))
}
	
func NewEnumerationReader(t hybrids.TableReader) corev1.EnumerationReader{
	if t==nil{
		return nil
	}
	return &EnumerationReader{_table:t}
}
type VectorEnumerationReader struct {
    _vectorHybrid    hybrids.VectorTableReader
    _vectorAllocated [] *EnumerationReader
}

func (ve *VectorEnumerationReader) Len() (size int) {

    if ve._vectorAllocated != nil {
        size = len(ve._vectorAllocated)
        return
    }

    if ve._vectorHybrid != nil {
        size = ve._vectorHybrid.Len()
        return
    }

    return
}

func (ve *VectorEnumerationReader) Get(i int) (item corev1.EnumerationReader, err error) {
    var table hybrids.TableReader

    if ve._vectorAllocated != nil {
        if i < 0 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(ve._vectorAllocated)}
            return
        }

        if i > len(ve._vectorAllocated)-1 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(ve._vectorAllocated)}
            return
        }

        item = ve._vectorAllocated[i]
        return
    }

    if ve._vectorHybrid != nil {
        table, err = ve._vectorHybrid.Get(i)
        item = NewEnumerationReader(table)
        return
    }

    err = &hybrids.VectorInvalidIndexError{Index: i, Len: 0}
    return
}

func NewVectorEnumerationReader(v hybrids.VectorTableReader) corev1.VectorEnumerationReader {
    if v == nil {
        return nil
    }
    return &VectorEnumerationReader{_vectorHybrid: v}
}
