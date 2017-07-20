package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//ResourceReader Resource type
type ResourceReader struct {
    _table hybrids.TableReader
    meta corev1.MetadataReader
}

//Meta ...
func (r *ResourceReader) Meta() corev1.MetadataReader {

	if r.meta != nil {
		return r.meta
	}

	return NewMetadataReader(r._table.Table(0))
}

func NewResourceReader(t hybrids.TableReader) corev1.ResourceReader{
	if t==nil{
		return nil
	}
	return &ResourceReader{_table:t}
}