package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//ResourceReader ...
type ResourceReader struct {
    _table hybrids.TableReader
    _resource hybrids.ResourceReader
    _rid corev1.ResourceIDReader
    meta *MetadataReader
    fields *VectorFieldReader
}

//RID get resource id
func (r *ResourceReader) RID() corev1.ResourceIDReader {
	return r._rid
}

//Meta ...
func (r *ResourceReader) Meta() corev1.MetadataReader {

	if r.meta != nil {
		return r.meta
	}

	return NewMetadataReader(r._table.Table(1))
}

//Fields ...
func (r *ResourceReader) Fields() corev1.VectorFieldReader {

	if r.fields != nil {
		return r.fields
	}

	return NewVectorFieldReader(r._table.VectorTable(2))
}
	
func NewResourceReader(t hybrids.TableReader) corev1.ResourceReader{
	if t==nil{
		return nil
	}
	return &ResourceReader{_table:t}
}
type VectorResourceReader struct {
    _vectorHybrid    hybrids.VectorTableReader
    _vectorAllocated [] *ResourceReader
}

func (vr *VectorResourceReader) Len() (size int) {

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

func (vr *VectorResourceReader) Get(i int) (item corev1.ResourceReader, err error) {
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

func NewVectorResourceReader(v hybrids.VectorTableReader) corev1.VectorResourceReader {
    if v == nil {
        return nil
    }
    return &VectorResourceReader{_vectorHybrid: v}
}
