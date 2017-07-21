package corev1Hybrids



//UnionReader ...
type UnionReader interface {

    //RID get resource id
    RID() ResourceIDReader

    //Meta ...
    Meta() MetadataReader

}

type VectorUnionReader interface {
     Len() int
     Get(i int) (item UnionReader, err error)
}
