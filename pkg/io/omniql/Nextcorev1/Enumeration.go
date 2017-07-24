package Nextcorev1



//EnumerationReader ...
type EnumerationReader interface {

    //RID get resource id
    RID() ResourceIDReader

    //Meta ...
    Meta() (MetadataReader, error)

    //Kind ...
    Kind() BasicType

    //Items ...
    Items() VectorEnumerationItemReader

    //Groups ...
    Groups() VectorEnumerationGroupReader

}

//VectorEnumerationReader ...
type VectorEnumerationReader interface {

    // Returns the current size of this vector
    Len() int

    //Get the item in the position i, if i < Len(),
    //if item does not exist should return the default value for the underlying data type
    //when i > Len() should return an VectorInvalidIndexError
    Get(i int) (item EnumerationReader, err error)
}
