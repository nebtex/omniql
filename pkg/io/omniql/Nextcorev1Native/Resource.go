package Nextcorev1Native

//Resource ...
type Resource struct {

    RID []byte `json:"rid"`
    Metadata *Metadata `json:"metadata"`
    Fields []*Field `json:"fields"`
}

//Resource ...
type Resource struct {

    Metadata *Metadata `json:"metadata"`
    Fields []*Field `json:"fields"`
}

//ResourceReader ...
type ResourceReader struct {
    _resource *Resource
    metadata *MetadataReader
    fields *VectorFieldReader
}

//RID get resource id
func (r *ResourceReader) RID() corev1.ResourceIDReader {
	return r._rid
}

//Metadata ...
func (r *ResourceReader) Metadata() corev1.MetadataReader {

	if r.metadata != nil {
		return r.metadata
	}

	return nil
}

//Fields ...
func (r *ResourceReader) Fields() corev1.VectorFieldReader {

	if r.fields != nil {
		return r.fields
	}

	return nil
}
	
//NewResourceReader ...
func NewResourceReader(r *ResourceReader) corev1.ResourceReader{
	if r!=nil{
		return &ResourceReader{_resource:r}
	}
	return nil
}

type VectorResourceReader struct {
    _vector  []*ResourceReader
}

func (vr *VectorResourceReader) Len() (size int) {
    size = len(vr._vector)
    return
}

func (vr *VectorResourceReader) Get(i int) (item corev1.ResourceReader, err error) {

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


func NewVectorResourceReader(v hybrids.VectorTableReader) corev1.VectorResourceReader {
    if v == nil {
        return nil
    }
    return &VectorResourceReader{_vectorHybrid: v}
}
