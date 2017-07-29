package Nextcorev1Native

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1"
)

//Resource ...
type Resource struct {
	RID      []byte `json:"rid"`
	Metadata *Metadata `json:"metadata"`
	Fields   []*Field `json:"fields"`
}

//ResourceReader ...
type ResourceReader struct {
	_resource *Resource
	metadata  *MetadataReader
	fields    *VectorFieldReader
}

//RID get resource id
func (r *ResourceReader) RID() Nextcorev1.ResourceIDReader {
	return r._rid
}

//Metadata ...
func (r *ResourceReader) Metadata() (mr Nextcorev1.MetadataReader, err error) {

	if r.metadata != nil {
		mr = r.metadata
	}

	return
}

//Fields ...
func (r *ResourceReader) Fields() Nextcorev1.VectorFieldReader {

	if r.fields != nil {
		return r.fields
	}

	return nil
}

//NewResourceReader ...
func NewResourceReader(r *Resource) *ResourceReader {
	if r != nil {
		return &ResourceReader{
			_resource: r,
			metadata:  NewMetadataReader(r.Metadata),
		}
	}
	return nil
}

//VectorResourceReader ...
type VectorResourceReader struct {
	_vector []*ResourceReader
}

//Len Returns the current size of this vector
func (vr *VectorResourceReader) Len() (size int) {
	size = len(vr._vector)
	return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
func (vr *VectorResourceReader) Get(i int) (item Nextcorev1.ResourceReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vr._vector)}
		return
	}

	if i > len(vr._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vr._vector)}
		return
	}

	item = vr._vector[i]
	return

}

//NewVectorResourceReader ...
func NewVectorResourceReader(vr []*Resource) (vrr *VectorResourceReader) {
	vrr = &VectorResourceReader{}
	vrr._vector = make([]*ResourceReader, len(vr))

	for i := 0; i < len(vr); i++ {
		vrr._vector[i] = NewResourceReader(vr[i])
	}
	return
}
