package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//UnionResourceReader ...
type UnionResourceReader struct {
    _table hybrids.TableReader
}

func NewUnionResourceReader(t hybrids.TableReader) corev1.UnionResourceReader{
	if t==nil{
		return nil
	}
	return &UnionResourceReader{_table:t}
}