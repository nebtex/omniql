package Nextcorev1Native

//Table ...
type Table struct {

    RID []byte `json:"rid"`
    Metadata *Metadata `json:"metadata"`
    Fields []*Field `json:"fields"`
}

//TableReader ...
type TableReader struct {
    _table *Table
    metadata *MetadataReader
    fields *VectorFieldReader
}

//RID get resource id
func (t *TableReader) RID() corev1.ResourceIDReader {
	return t._rid
}

//Metadata ...
func (t *TableReader) Metadata() corev1.MetadataReader {

	if t.metadata != nil {
		return t.metadata
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
	
//NewTableReader ...
func NewTableReader(t *TableReader) corev1.TableReader{
	if t!=nil{
		return &TableReader{_table:t}
	}
	return nil
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
