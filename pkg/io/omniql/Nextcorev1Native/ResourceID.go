package Nextcorev1Native
import(
    "github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1"
    "github.com/nebtex/hybrids/golang/hybrids"
)
//ResourceID ...
type ResourceID struct {

    Application string `json:"application"`
    Kind string `json:"kind"`
    ID string `json:"id"`
    Parent *ResourceID `json:"parent"`
}

//ResourceIDReader ...
type ResourceIDReader struct {
    _resourceid *ResourceID
    parent *ResourceIDReader
}

//Application ...
func (ri *ResourceIDReader) Application() (value string) {
	value = ri._resourceid.Application
	return
}

//Kind ...
func (ri *ResourceIDReader) Kind() (value Nextcorev1.ApplicationType) {
	value = Nextcorev1.FromStringToApplicationType(ri._resource_id.Kind)

	if !value.IsResource(){
		value = Nextcorev1.ApplicationTypeNone
	}

	return
}

//ID ...
func (ri *ResourceIDReader) ID() (value string) {
	value = ri._resourceid.ID
	return
}



//Parent ...
func (ri *ResourceIDReader) Parent() (rir Nextcorev1.ResourceIDReader, err error) {

	if ri.parent != nil {
		rir  =  ri.parent
	}

	return
}

//NewResourceIDReader ...
func NewResourceIDReader(ri *ResourceID) *ResourceIDReader{
	if ri!=nil{
		return &ResourceIDReader{
		                                   _resourceid:ri,
parent: NewResourceIDReader(ri.Parent),
		                                   }
	}
	return nil
}



//VectorResourceIDReader ...
type VectorResourceIDReader struct {
    _vector  []*ResourceIDReader
}

//Len Returns the current size of this vector
func (vri *VectorResourceIDReader) Len() (size int) {
    size = len(vri._vector)
    return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
func (vri *VectorResourceIDReader) Get(i int) (item Nextcorev1.ResourceIDReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vri._vector)}
		return
	}

	if i > len(vri._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vri._vector)}
		return
	}

	item = vri._vector[i]
	return


}

//NewVectorResourceIDReader ...
func NewVectorResourceIDReader(vri []*ResourceID) (vrir *VectorResourceIDReader) {
    vrir = &VectorResourceIDReader{}
	vrir._vector = make([]*ResourceIDReader, len(vri))

	for i := 0; i < len(vri); i++ {
		vrir._vector[i] = NewResourceIDReader(vri[i])
	}
	return
}
