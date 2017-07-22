package Nextcorev1



//EnumerationReader ...
type EnumerationReader interface {

    //RID get resource id
    RID() ResourceIDReader

    //Meta ...
    Meta() (MetadataReader, error)

    //Items ...
    Items() VectorEnumerationItemReader

    //Groups ...
    Groups() VectorEnumerationGroupReader

}

type VectorEnumerationReader interface {
     Len() int
     Get(i int) (item EnumerationReader, err error)
}
