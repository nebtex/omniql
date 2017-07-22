package Nextcorev1Native

import ("github.com/nebtex/hybrids/golang/hybrids"
	    "github.com/nebtex/omnibuff/pkg/io/omniql/corev1")



//Union ...
type Union struct {

    RID []byte `json:"rid"`
    Meta *Metadata `json:"meta"`
}

//UnionReader ...
type UnionReader struct {
    _union *Union
    meta *MetadataReader
}

//RID get resource id
func (u *UnionReader) RID() corev1.ResourceIDReader {
	return u._rid
}

//Meta ...
func (u *UnionReader) Meta() corev1.MetadataReader {

	if u.meta != nil {
		return u.meta
	}

	return nil
}

func NewUnionReader(t hybrids.TableReader) corev1.UnionReader{
	if t==nil{
		return nil
	}
	return &UnionReader{_table:t}
}

type VectorUnionReader struct {
    _vector  []*UnionReader
}

func (vu *VectorUnionReader) Len() (size int) {
    size = len(vu._vector)
    return
}

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

func NewVectorUnionReader(v hybrids.VectorTableReader) corev1.VectorUnionReader {
    if v == nil {
        return nil
    }
    return &VectorUnionReader{_vectorHybrid: v}
}
