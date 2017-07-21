package corev1Hybrids



//TableReader ...
type TableReader interface {

    //RID get resource id
    RID() ResourceIDReader

    //Meta ...
    Meta() MetadataReader

    //Fields ...
    Fields() VectorFieldReader

}

type VectorTableReader interface {
     Len() int
     Get(i int) (item TableReader, err error)
}
