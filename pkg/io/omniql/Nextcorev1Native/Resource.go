package Nextcorev1Native

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//Resource Resource type
type Resource struct {

    RID []byte `json:"rid"`
    Meta *Metadata `json:"meta"`
    Fields []*Field `json:"fields"`
}

//ResourceReader Resource type
type ResourceReader struct {
    _resource *Resource
    meta *MetadataReader
    fields *VectorFieldReader
}

//RID get resource id
func (r *ResourceReader) RID() corev1.ResourceIDReader {
	return r._rid
}

//Meta ...
func (r *ResourceReader) Meta() corev1.MetadataReader {

	if r.meta != nil {
		return r.meta
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
	
func NewResourceReader(t hybrids.TableReader) corev1.ResourceReader{
	if t==nil{
		return nil
	}
	return &ResourceReader{_table:t}
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
