package corev1Hybrids

import "github.com/nebtex/hybrids/golang/hybrids"


type MetadataReader struct {
	_table hybrids.TableReader
	doc DocumentationReader
}

func (m *MetadataReader) Application() (value string) {
	value, _ = m._table.String(0)
	return
}


func (m *MetadataReader) Name() (value string) {
	value, _ = m._table.String(1)
	return
}


func (m *MetadataReader) Kind() (value string) {
	value, _ = m._table.String(2)
	return
}


func (m *MetadataReader) Parent() (value string) {
	value, _ = m._table.String(3)
	return
}
