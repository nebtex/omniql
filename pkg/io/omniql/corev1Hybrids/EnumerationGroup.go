package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//EnumerationGroupReader allow to group enumerations
type EnumerationGroupReader struct {
    _table hybrids.TableReader
    documentation corev1.DocumentationReader
}

//Documentation ...
func (e *EnumerationGroupReader) Documentation() corev1.DocumentationReader {

	if e.documentation != nil {
		return e.documentation
	}

	return NewDocumentationReader(e._table.Table(0))
}

func NewEnumerationGroupReader(t hybrids.TableReader) corev1.EnumerationGroupReader{
	if t==nil{
		return nil
	}
	return &EnumerationGroupReader{_table:t}
}