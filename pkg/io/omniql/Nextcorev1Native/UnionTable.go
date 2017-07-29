package Nextcorev1Native
import(
    "github.com/nebtex/hybrids/golang/hybrids"
    "github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1"
)
//UnionTable ...
type UnionTable struct {

}

//UnionTableReader ...
type UnionTableReader struct {
    _uniontable *UnionTable
}

//NewUnionTableReader ...
func NewUnionTableReader(ut *UnionTable) *UnionTableReader{
	if ut!=nil{
		return &UnionTableReader{
		                                   _uniontable:ut,
		                                   }
	}
	return nil
}



//VectorUnionTableReader ...
type VectorUnionTableReader struct {
    _vector  []*UnionTableReader
}

//Len Returns the current size of this vector
func (vut *VectorUnionTableReader) Len() (size int) {
    size = len(vut._vector)
    return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
func (vut *VectorUnionTableReader) Get(i int) (item Nextcorev1.UnionTableReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vut._vector)}
		return
	}

	if i > len(vut._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vut._vector)}
		return
	}

	item = vut._vector[i]
	return


}

//NewVectorUnionTableReader ...
func NewVectorUnionTableReader(vut []*UnionTable) (vutr *VectorUnionTableReader) {
    vutr = &VectorUnionTableReader{}
	vutr._vector = make([]*UnionTableReader, len(vut))

	for i := 0; i < len(vut); i++ {
		vutr._vector[i] = NewUnionTableReader(vut[i])
	}
	return
}
