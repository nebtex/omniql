package Nextcorev1



//UnionReader ...
type UnionReader interface {

    //RID get resource id
    RID() ResourceIDReader

    //Meta ...
    Meta() (MetadataReader, error)

}

type VectorUnionReader interface {
     Len() int
     Get(i int) (item UnionReader, err error)
}
