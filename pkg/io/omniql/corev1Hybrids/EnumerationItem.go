package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//EnumerationItemReader ...
type EnumerationItemReader struct {
    _table hybrids.TableReader
    documentation corev1.DocumentationReader
}

//Name ...
func (i *EnumerationItemReader) Name() (value string) {
	value, _ = i._table.String(0)
	return
}

//Documentation ...
func (i *EnumerationItemReader) Documentation() corev1.DocumentationReader {

	if i.documentation != nil {
		return i.documentation
	}

	return NewDocumentationReader(i._table.Table(1))
}

func NewEnumerationItemReader(t hybrids.TableReader) corev1.EnumerationItemReader{
	if t==nil{
		return nil
	}
	return &EnumerationItemReader{_table:t}
}
type VectorEnumerationItemReader struct {
    _vectorHybrid    hybrids.VectorTableReader
    _vectorAllocated [] corev1.EnumerationItemReader
}

func (vi *VectorEnumerationItemReader) Len() (size int) {

    if vi._vectorAllocated != nil {
        size = len(vi._vectorAllocated)
        return
    }

    if vi._vectorHybrid != nil {
        size = vi._vectorHybrid.Len()
        return
    }

    return
}

func (vi *VectorEnumerationItemReader) Get(i int) (item corev1.EnumerationItemReader, err error) {
    var table hybrids.TableReader

    if vi._vectorAllocated != nil {
        if i < 0 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vi._vectorAllocated)}
            return
        }

        if i > len(vi._vectorAllocated)-1 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vi._vectorAllocated)}
            return
        }

        item = vi._vectorAllocated[i]
        return
    }

    if vi._vectorHybrid != nil {
        table, err = vi._vectorHybrid.Get(i)
        item = NewEnumerationReader(table)
        return
    }

    err = &hybrids.VectorInvalidIndexError{Index: i, Len: 0}
    return
}

func NewVectorEnumerationItemReader(v hybrids.VectorTableReader) corev1.VectorEnumerationItemReader {
    if v == nil {
        return nil
    }
    return &VectorEnumerationItemReader{_vectorHybrid: v}
}
