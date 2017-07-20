package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//TableReader ...
type TableReader struct {
    _table hybrids.TableReader
    meta corev1.MetadataReader
}

//Meta ...
func (t *TableReader) Meta() corev1.MetadataReader {

	if t.meta != nil {
		return t.meta
	}

	return NewMetadataReader(t._table.Table(0))
}

func NewTableReader(t hybrids.TableReader) corev1.TableReader{
	if t==nil{
		return nil
	}
	return &TableReader{_table:t}
}