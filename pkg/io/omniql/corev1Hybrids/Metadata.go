package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//MetadataReader meta data
type MetadataReader struct {
    _table hybrids.TableReader
    documentation corev1.DocumentationReader
}

//Application ...
func (m *MetadataReader) Application() (value string) {
	value, _ = m._table.String(0)
	return
}

//Name ...
func (m *MetadataReader) Name() (value string) {
	value, _ = m._table.String(1)
	return
}

//Kind ...
func (m *MetadataReader) Kind() (value string) {
	value, _ = m._table.String(2)
	return
}

//Parent ...
func (m *MetadataReader) Parent() (value string) {
	value, _ = m._table.String(3)
	return
}

//Documentation ...
func (m *MetadataReader) Documentation() corev1.DocumentationReader {

	if m.documentation != nil {
		return m.documentation
	}

	return NewDocumentationReader(m._table.Table(4))
}

func NewMetadataReader(t hybrids.TableReader) corev1.MetadataReader{
	if t==nil{
		return nil
	}
	return &MetadataReader{_table:t}
}