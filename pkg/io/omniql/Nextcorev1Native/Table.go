package Nextcorev1Native

import "github.com/nebtex/omnibuff/pkg/io/omniql/corev1"

//Table ...
type Table struct {

    RID []byte `json:"rid"`
    Meta *Metadata `json:"meta"`
    Fields []*Field `json:"fields"`
}

//TableReader ...
type TableReader struct {
    _table *Table
//TableReader ...
type TableReader struct {
    _table *Table
    meta *MetadataReader
    fields *VectorFieldReader
}

}

//RID get resource id
func (t *TableReader) RID() corev1.ResourceIDReader {
	return t._rid
}

//Meta ...
func (t *TableReader) Meta() corev1.MetadataReader {

	if t.meta != nil {
		return t.meta
	}

	return nil
}

//Fields ...
func (t *TableReader) Fields() corev1.VectorFieldReader {

	if t.fields != nil {
		return t.fields
	}

	return nil
}
	
func NewTableReader(t hybrids.TableReader) corev1.TableReader{
	if t==nil{
		return nil
	}
	return &TableReader{_table:t}
}

type VectorTableReader struct {
    _vector  []*TableReader
}

func (vt *VectorTableReader) Len() (size int) {
    size = len(vt._vector)
    return
}

func (vt *VectorTableReader) Get(i int) (item corev1.TableReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vt._vector)}
		return
	}

	if i > len(vt._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vt._vector)}
		return
	}

	item = vt._vector[i]
	return


}

func NewVectorTableReader(v hybrids.VectorTableReader) corev1.VectorTableReader {
    if v == nil {
        return nil
    }
    return &VectorTableReader{_vectorHybrid: v}
}
