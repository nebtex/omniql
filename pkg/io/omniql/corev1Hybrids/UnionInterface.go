package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//UnionInterfaceReader ...
type UnionInterfaceReader struct {
    _table hybrids.TableReader
}

func NewUnionInterfaceReader(t hybrids.TableReader) corev1.UnionInterfaceReader{
	if t==nil{
		return nil
	}
	return &UnionInterfaceReader{_table:t}
}