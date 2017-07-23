package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//MetadataReader ...
type MetadataReader struct {
    _table hybrids.TableReader
    documentation *DocumentationReader
}

//Application ...
func (m *MetadataReader) Application() (value string) {
	value, _ = m._table.String(0)
	return
}

//Name ...
func (m *MetadataReader) Name() (value string) {
	value, _ = m._table.String(1)
	return
}

//Kind ...
func (m *MetadataReader) Kind() (value string) {
	value, _ = m._table.String(2)
	return
}

//Parent ...
func (m *MetadataReader) Parent() (value string) {
	value, _ = m._table.String(3)
	return
}

//Documentation ...
func (m *MetadataReader) Documentation() corev1.DocumentationReader {

	if m.documentation != nil {
		return m.documentation
	}

	return NewDocumentationReader(m._table.Table(4))
}

func NewMetadataReader(t hybrids.TableReader) corev1.MetadataReader{
	if t==nil{
		return nil
	}
	return &MetadataReader{_table:t}
}
type VectorMetadataReader struct {
    _vectorHybrid    hybrids.VectorTableReader
    _vectorAllocated [] *MetadataReader
}

func (vm *VectorMetadataReader) Len() (size int) {

    if vm._vectorAllocated != nil {
        size = len(vm._vectorAllocated)
        return
    }

    if vm._vectorHybrid != nil {
        size = vm._vectorHybrid.Len()
        return
    }

    return
}

func (vm *VectorMetadataReader) Get(i int) (item corev1.MetadataReader, err error) {
    var table hybrids.TableReader

    if vm._vectorAllocated != nil {
        if i < 0 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vm._vectorAllocated)}
            return
        }

        if i > len(vm._vectorAllocated)-1 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vm._vectorAllocated)}
            return
        }

        item = vm._vectorAllocated[i]
        return
    }

    if vm._vectorHybrid != nil {
        table, err = vm._vectorHybrid.Get(i)
        item = NewEnumerationReader(table)
        return
    }

    err = &hybrids.VectorInvalidIndexError{Index: i, Len: 0}
    return
}

func NewVectorMetadataReader(v hybrids.VectorTableReader) corev1.VectorMetadataReader {
    if v == nil {
        return nil
    }
    return &VectorMetadataReader{_vectorHybrid: v}
}
