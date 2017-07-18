package corev1Native

import "github.com/nebtex/omnibuff/pkg/io/omniql/corev1"

type Table struct {
	RID    []byte  `json:"rid"`
	Meta   *Metadata `json:"meta"`
	Fields []*Field `json:"fields"`
}

type TableReader struct {
	table *Table
	meta  corev1.MetadataReader
	doc   corev1.DocumentationReader
}

func (t *TableReader) RID() corev1.ResourceIDReader {
	return NewIDReader(t.table.RID, true)
}

func (t *TableReader) Meta() corev1.MetadataReader {
	return &MetadataReader{meta: t.table.Meta}

}

func (t *TableReader) Fields() corev1.VectorFieldReader {
	return &VectorFieldReader{nativeFields: t.table.Fields}

}

func NewTableReader(t *Table) corev1.TableReader {
	return &TableReader{table: t}

}
