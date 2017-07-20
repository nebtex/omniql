package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//EnumerationItemReader ...
type EnumerationItemReader struct {
    _table hybrids.TableReader
}

//Name ...
func (e *EnumerationItemReader) Name() (value string) {
	value, _ = e._table.String(0)
	return
}

func NewEnumerationItemReader(t hybrids.TableReader) corev1.EnumerationItemReader{
	if t==nil{
		return nil
	}
	return &EnumerationItemReader{_table:t}
}