package corev1Native

import "github.com/nebtex/omniql/pkg/io/omniql/corev1"

type Table struct {
	Metadata *Metadata `json:"meta"`
	Fields   []*Field  `json:"fields"`
}

type TableReader struct {
	table    *Table
	metadata corev1.MetadataReader
	doc      corev1.DocumentationReader
}

func (t *TableReader) RID() corev1.ResourceIDReader {
	return NewIDReader(t.table.RID, true)
}

func (t *TableReader) Metadata() corev1.MetadataReader {
	return &MetadataReader{metadata: t.table.Metadata}

}

func (t *TableReader) Fields() corev1.VectorFieldReader {
	return &VectorFieldReader{nativeFields: t.table.Fields}

}

func NewTableReader(t *Table) corev1.TableReader {
	return &TableReader{table: t}

}
