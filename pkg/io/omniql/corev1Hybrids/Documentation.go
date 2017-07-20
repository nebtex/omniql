package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//DocumentationReader documentation type
type DocumentationReader struct {
    _table hybrids.TableReader
}

//Short ...
func (d *DocumentationReader) Short() (value string) {
	value, _ = d._table.String(0)
	return
}

//Long ...
func (d *DocumentationReader) Long() (value string) {
	value, _ = d._table.String(1)
	return
}

func NewDocumentationReader(t hybrids.TableReader) corev1.DocumentationReader{
	if t==nil{
		return nil
	}
	return &DocumentationReader{_table:t}
}