package Nextcorev1



//ResourceReader Resource type
type ResourceReader interface {

    //RID get resource id
    RID() ResourceIDReader

    //Meta ...
    Meta() (MetadataReader, error)

    //Fields ...
    Fields() VectorFieldReader

}

type VectorResourceReader interface {
     Len() int
     Get(i int) (item ResourceReader, err error)
}
