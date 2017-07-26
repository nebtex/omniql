package Nextcorev1Native
import(
    "github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1"
    "github.com/nebtex/hybrids/golang/hybrids"
)
//Metadata ...
type Metadata struct {

    Application string `json:"application"`
    Name string `json:"name"`
    Kind string `json:"kind"`
    Parent string `json:"parent"`
    Documentation *Documentation `json:"documentation"`
}

//MetadataReader ...
type MetadataReader struct {
    _metadata *Metadata
    documentation *DocumentationReader
}

//Application ...
func (m *MetadataReader) Application() (value string) {
	value = m._metadata.Application
	return
}

//Name ...
func (m *MetadataReader) Name() (value string) {
	value = m._metadata.Name
	return
}

//Kind ...
func (m *MetadataReader) Kind() (value Nextcorev1.ApplicationType) {
	value = Nextcorev1.FromStringToApplicationType(m._metadata.Kind)

	if !value.IsResource(){
		value = Nextcorev1.ApplicationTypeNone
	}

	return
}

//Parent ...
func (m *MetadataReader) Parent() (value string) {
	value = m._metadata.Parent
	return
}

//Documentation ...
func (m *MetadataReader) Documentation() (dr Nextcorev1.DocumentationReader, err error) {

	if m.documentation != nil {
		dr  =  m.documentation
	}

	return
}

//NewMetadataReader ...
func NewMetadataReader(m *Metadata) *MetadataReader{
	if m!=nil{
		return &MetadataReader{
		                                   _metadata:m,
documentation: NewDocumentationReader(m.Documentation),
		                                   }
	}
	return nil
}

//VectorMetadataReader ...
type VectorMetadataReader struct {
    _vector  []*MetadataReader
}

//Len Returns the current size of this vector
func (vm *VectorMetadataReader) Len() (size int) {
    size = len(vm._vector)
    return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
func (vm *VectorMetadataReader) Get(i int) (item Nextcorev1.MetadataReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vm._vector)}
		return
	}

	if i > len(vm._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vm._vector)}
		return
	}

	item = vm._vector[i]
	return


}

//NewVectorMetadataReader ...
func NewVectorMetadataReader(vm []*Metadata) (vmr *VectorMetadataReader) {
    vmr = &VectorMetadataReader{}
	vmr._vector = make([]*MetadataReader, len(vm))

	for i := 0; i < len(vm); i++ {
		vmr._vector[i] = NewMetadataReader(vm[i])
	}
	return
}
