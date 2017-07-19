package corev1Hybrids

import "github.com/nebtex/hybrids/golang/hybrids"


type DocumentationReader struct {
	_table hybrids.TableReader
}

func (d *DocumentationReader) Short() (value string) {
	value, _ = d._table.String(0)
	return
}


func (d *DocumentationReader) Long() (value string) {
	value, _ = d._table.String(1)
	return
}
