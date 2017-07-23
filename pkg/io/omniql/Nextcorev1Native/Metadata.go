package Nextcorev1Native

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



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
func (m *MetadataReader) Kind() (value string) {
	value = m._metadata.Kind
	return
}

//Parent ...
func (m *MetadataReader) Parent() (value string) {
	value = m._metadata.Parent
	return
}

//Documentation ...
func (m *MetadataReader) Documentation() Nextcorev1.DocumentationReader {

	if m.documentation != nil {
		return m.documentation
	}

	return nil
}

func NewMetadataReader(t hybrids.TableReader) Nextcorev1.MetadataReader{
	if t==nil{
		return nil
	}
	return &MetadataReader{_table:t}
}

type VectorMetadataReader struct {
    _vector  []*MetadataReader
}

func (vm *VectorMetadataReader) Len() (size int) {
    size = len(vm._vector)
    return
}

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

func NewVectorMetadataReader(v hybrids.VectorTableReader) Nextcorev1.VectorMetadataReader {
    if v == nil {
        return nil
    }
    return &VectorMetadataReader{_vectorHybrid: v}
}
