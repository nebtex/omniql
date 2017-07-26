package Nextcorev1Native

//Table ...
type Table struct {

    RID []byte `json:"rid"`
    Metadata *Metadata `json:"metadata"`
    Fields []*Field `json:"fields"`
}
import(
    "github.com/nebtex/hybrids/golang/hybrids"
    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
)
//Table ...
type Table struct {

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
func (t *TableReader) Metadata() (mr corev1.MetadataReader, err error) {

	if t.metadata != nil {
		mr  =  t.metadata
	}

	return
}

//Fields ...
func (t *TableReader) Fields() corev1.VectorFieldReader {

	if t.fields != nil {
		return t.fields
	}

	return nil
}
	
//NewTableReader ...
func NewTableReader(t *Table) *TableReader{
	if t!=nil{
		return &TableReader{
		                                   _table:t,
metadata: NewMetadataReader(t.Metadata),
		                                   }
	}
	return nil
}

//VectorTableReader ...
type VectorTableReader struct {
    _vector  []*TableReader
}

//Len Returns the current size of this vector
func (vt *VectorTableReader) Len() (size int) {
    size = len(vt._vector)
    return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
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

//NewVectorTableReader ...
func NewVectorTableReader(vt []*Table) (vtr *VectorTableReader) {
    vtr = &VectorTableReader{}
	vtr._vector = make([]*TableReader, len(vt))

	for i := 0; i < len(vt); i++ {
		vtr._vector[i] = NewTableReader(vt[i])
	}
	return
}
