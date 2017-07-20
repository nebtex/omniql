package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//InterfaceReader ...
type InterfaceReader struct {
    _table hybrids.TableReader
    meta corev1.MetadataReader
}

//Meta ...
func (i *InterfaceReader) Meta() corev1.MetadataReader {

	if i.meta != nil {
		return i.meta
	}

	return NewMetadataReader(i._table.Table(0))
}

func NewInterfaceReader(t hybrids.TableReader) corev1.InterfaceReader{
	if t==nil{
		return nil
	}
	return &InterfaceReader{_table:t}
}