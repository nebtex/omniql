package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//UnionReader ...
type UnionReader struct {
    _table hybrids.TableReader
    meta corev1.MetadataReader
}

//Meta ...
func (u *UnionReader) Meta() corev1.MetadataReader {

	if u.meta != nil {
		return u.meta
	}

	return NewMetadataReader(u._table.Table(0))
}

func NewUnionReader(t hybrids.TableReader) corev1.UnionReader{
	if t==nil{
		return nil
	}
	return &UnionReader{_table:t}
}