package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//EnumerationReader ...
type EnumerationReader struct {
    _table hybrids.TableReader
    meta corev1.MetadataReader
}

//Meta ...
func (e *EnumerationReader) Meta() corev1.MetadataReader {

	if e.meta != nil {
		return e.meta
	}

	return NewMetadataReader(e._table.Table(0))
}

func NewEnumerationReader(t hybrids.TableReader) corev1.EnumerationReader{
	if t==nil{
		return nil
	}
	return &EnumerationReader{_table:t}
}