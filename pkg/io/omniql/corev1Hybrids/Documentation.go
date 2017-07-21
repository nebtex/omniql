package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//DocumentationReader documentation type
type DocumentationReader struct {
    _table hybrids.TableReader
}

//Short ...
func (d *DocumentationReader) Short() (value string) {
	value, _ = d._table.String(0)
	return
}

//Long ...
func (d *DocumentationReader) Long() (value string) {
	value, _ = d._table.String(1)
	return
}

func NewDocumentationReader(t hybrids.TableReader) corev1.DocumentationReader{
	if t==nil{
		return nil
	}
	return &DocumentationReader{_table:t}
}
type VectorDocumentationReader struct {
    _vectorHybrid    hybrids.VectorTableReader
    _vectorAllocated [] *DocumentationReader
}

func (vd *VectorDocumentationReader) Len() (size int) {

    if vd._vectorAllocated != nil {
        size = len(vd._vectorAllocated)
        return
    }

    if vd._vectorHybrid != nil {
        size = vd._vectorHybrid.Len()
        return
    }

    return
}

func (vd *VectorDocumentationReader) Get(i int) (item corev1.DocumentationReader, err error) {
    var table hybrids.TableReader

    if vd._vectorAllocated != nil {
        if i < 0 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vd._vectorAllocated)}
            return
        }

        if i > len(vd._vectorAllocated)-1 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vd._vectorAllocated)}
            return
        }

        item = vd._vectorAllocated[i]
        return
    }

    if vd._vectorHybrid != nil {
        table, err = vd._vectorHybrid.Get(i)
        item = NewEnumerationReader(table)
        return
    }

    err = &hybrids.VectorInvalidIndexError{Index: i, Len: 0}
    return
}

func NewVectorDocumentationReader(v hybrids.VectorTableReader) corev1.VectorDocumentationReader {
    if v == nil {
        return nil
    }
    return &VectorDocumentationReader{_vectorHybrid: v}
}
