package corev1Hybrids



//ResourceReader Resource type
type ResourceReader interface {

    //RID get resource id
    RID() ResourceIDReader

    //Meta ...
    Meta() MetadataReader

    //Fields ...
    Fields() VectorFieldReader

}

type VectorResourceReader interface {
     Len() int
     Get(i int) (item ResourceReader, err error)
}
