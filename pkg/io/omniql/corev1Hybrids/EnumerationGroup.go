package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//EnumerationGroupReader ...
type EnumerationGroupReader struct {
    _table hybrids.TableReader
    documentation *DocumentationReader
}

//Name ...
func (g *EnumerationGroupReader) Name() (value string) {
	value, _ = g._table.String(0)
	return
}

//Documentation ...
func (g *EnumerationGroupReader) Documentation() corev1.DocumentationReader {

	if g.documentation != nil {
		return g.documentation
	}

	return NewDocumentationReader(g._table.Table(1))
}

//Items ...
func (g *EnumerationGroupReader) Items() hybrids.VectorStringReader {
    return g._table.VectorString(2)
}

func NewEnumerationGroupReader(t hybrids.TableReader) corev1.EnumerationGroupReader{
	if t==nil{
		return nil
	}
	return &EnumerationGroupReader{_table:t}
}
type VectorEnumerationGroupReader struct {
    _vectorHybrid    hybrids.VectorTableReader
    _vectorAllocated [] *EnumerationGroupReader
}

func (vg *VectorEnumerationGroupReader) Len() (size int) {

    if vg._vectorAllocated != nil {
        size = len(vg._vectorAllocated)
        return
    }

    if vg._vectorHybrid != nil {
        size = vg._vectorHybrid.Len()
        return
    }

    return
}

func (vg *VectorEnumerationGroupReader) Get(i int) (item corev1.EnumerationGroupReader, err error) {
    var table hybrids.TableReader

    if vg._vectorAllocated != nil {
        if i < 0 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vg._vectorAllocated)}
            return
        }

        if i > len(vg._vectorAllocated)-1 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vg._vectorAllocated)}
            return
        }

        item = vg._vectorAllocated[i]
        return
    }

    if vg._vectorHybrid != nil {
        table, err = vg._vectorHybrid.Get(i)
        item = NewEnumerationReader(table)
        return
    }

    err = &hybrids.VectorInvalidIndexError{Index: i, Len: 0}
    return
}

func NewVectorEnumerationGroupReader(v hybrids.VectorTableReader) corev1.VectorEnumerationGroupReader {
    if v == nil {
        return nil
    }
    return &VectorEnumerationGroupReader{_vectorHybrid: v}
}
