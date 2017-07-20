package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//UnionTableReader ...
type UnionTableReader struct {
    _table hybrids.TableReader
}

func NewUnionTableReader(t hybrids.TableReader) corev1.UnionTableReader{
	if t==nil{
		return nil
	}
	return &UnionTableReader{_table:t}
}