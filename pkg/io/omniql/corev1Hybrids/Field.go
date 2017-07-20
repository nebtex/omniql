package corev1Hybrids

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//FieldReader field type
type FieldReader struct {
    _table hybrids.TableReader
	_resource hybrids.ResourceReader
    documentation corev1.DocumentationReader
}

//Name ...
func (f *FieldReader) Name() (value string) {
	value, _ = f._table.String(0)
	return
}

//Type ...
func (f *FieldReader) Type() (value string) {
	value, _ = f._table.String(1)
	return
}

//Documentation ...
func (f *FieldReader) Documentation() corev1.DocumentationReader {

	if f.documentation != nil {
		return f.documentation
	}

	return NewDocumentationReader(f._table.Table(4))
}

//Default String representation of the default value
func (f *FieldReader) Default() (value string) {
	value, _ = f._table.String(5)
	return
}

func NewFieldReader(r hybrids.ResourceReader) corev1.FieldReader{
	if t==nil{
		return nil
	}
	return &FieldReader{_table:r.RootTable(), _resource:r}
}