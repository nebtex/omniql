package Nextcorev1Native

//Union ...
type Union struct {

    RID []byte `json:"rid"`
    Metadata *Metadata `json:"metadata"`
}
import(
    "github.com/nebtex/hybrids/golang/hybrids"
    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
)
//Union ...
type Union struct {

    Metadata *Metadata `json:"metadata"`
}

//UnionReader ...
type UnionReader struct {
    _union *Union
    metadata *MetadataReader
}

//RID get resource id
func (u *UnionReader) RID() corev1.ResourceIDReader {
	return u._rid
}

//Metadata ...
func (u *UnionReader) Metadata() (mr corev1.MetadataReader, err error) {

	if u.metadata != nil {
		mr  =  u.metadata
	}

	return
}

//NewUnionReader ...
func NewUnionReader(u *Union) *UnionReader{
	if u!=nil{
		return &UnionReader{
		                                   _union:u,
metadata: NewMetadataReader(u.Metadata),
		                                   }
	}
	return nil
}

//VectorUnionReader ...
type VectorUnionReader struct {
    _vector  []*UnionReader
}

//Len Returns the current size of this vector
func (vu *VectorUnionReader) Len() (size int) {
    size = len(vu._vector)
    return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
func (vu *VectorUnionReader) Get(i int) (item corev1.UnionReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vu._vector)}
		return
	}

	if i > len(vu._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vu._vector)}
		return
	}

	item = vu._vector[i]
	return


}

//NewVectorUnionReader ...
func NewVectorUnionReader(vu []*Union) (vur *VectorUnionReader) {
    vur = &VectorUnionReader{}
	vur._vector = make([]*UnionReader, len(vu))

	for i := 0; i < len(vu); i++ {
		vur._vector[i] = NewUnionReader(vu[i])
	}
	return
}
