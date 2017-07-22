package Nextcorev1



//TableReader ...
type TableReader interface {

    //RID get resource id
    RID() ResourceIDReader

    //Meta ...
    Meta() (MetadataReader, error)

    //Fields ...
    Fields() VectorFieldReader

}

type VectorTableReader interface {
     Len() int
     Get(i int) (item TableReader, err error)
}
