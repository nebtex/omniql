package Nextcorev1Native

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1"
)

//UnionResource ...
type UnionResource struct {
}

//UnionResourceReader ...
type UnionResourceReader struct {
	_unionresource *UnionResource
}

//NewUnionResourceReader ...
func NewUnionResourceReader(ur *UnionResource) *UnionResourceReader {
	if ur != nil {
		return &UnionResourceReader{
			_unionresource: ur,
		}
	}
	return nil
}

//VectorUnionResourceReader ...
type VectorUnionResourceReader struct {
	_vector []*UnionResourceReader
}

//Len Returns the current size of this vector
func (vur *VectorUnionResourceReader) Len() (size int) {
	size = len(vur._vector)
	return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
func (vur *VectorUnionResourceReader) Get(i int) (item Nextcorev1.UnionResourceReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vur._vector)}
		return
	}

	if i > len(vur._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vur._vector)}
		return
	}

	item = vur._vector[i]
	return

}

//NewVectorUnionResourceReader ...
func NewVectorUnionResourceReader(vur []*UnionResource) (vurr *VectorUnionResourceReader) {
	vurr = &VectorUnionResourceReader{}
	vurr._vector = make([]*UnionResourceReader, len(vur))

	for i := 0; i < len(vur); i++ {
		vurr._vector[i] = NewUnionResourceReader(vur[i])
	}
	return
}
