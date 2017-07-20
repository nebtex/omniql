package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//FieldReader field type
type FieldReader struct {
    _table hybrids.TableReader
    documentation corev1.DocumentationReader
}

//Name ...
func (f *FieldReader) Name() (value string) {
	value, _ = f._table.String(0)
	return
}

//Type ...
func (f *FieldReader) Type() (value string) {
	value, _ = f._table.String(1)
	return
}

//Documentation ...
func (f *FieldReader) Documentation() corev1.DocumentationReader {

	if f.documentation != nil {
		return f.documentation
	}

	return NewDocumentationReader(f._table.Table(4))
}

//Default String representation of the default value
func (f *FieldReader) Default() (value string) {
	value, _ = f._table.String(5)
	return
}

func NewFieldReader(t hybrids.TableReader) corev1.FieldReader{
	if t==nil{
		return nil
	}
	return &FieldReader{_table:t}
}
type VectorFieldReader struct {
    _vectorHybrid    hybrids.VectorTableReader
    _vectorAllocated [] corev1.FieldReader
}

func (vf *VectorFieldReader) Len() (size int) {

    if vf._vectorAllocated != nil {
        size = len(vf._vectorAllocated)
        return
    }

    if vf._vectorHybrid != nil {
        size = vf._vectorHybrid.Len()
        return
    }

    return
}

func (vf *VectorFieldReader) Get(i int) (item corev1.FieldReader, err error) {
    var table hybrids.TableReader

    if vf._vectorAllocated != nil {
        if i < 0 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vf._vectorAllocated)}
            return
        }

        if i > len(vf._vectorAllocated)-1 {
            err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vf._vectorAllocated)}
            return
        }

        item = vf._vectorAllocated[i]
        return
    }

    if vf._vectorHybrid != nil {
        table, err = vf._vectorHybrid.Get(i)
        item = NewEnumerationReader(table)
        return
    }

    err = &hybrids.VectorInvalidIndexError{Index: i, Len: 0}
    return
}

func NewVectorFieldReader(v hybrids.VectorTableReader) corev1.VectorFieldReader {
    if v == nil {
        return nil
    }
    return &VectorFieldReader{_vectorHybrid: v}
}
