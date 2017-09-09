package corev1



//ResourceReader ...
type ResourceReader interface {

    //RID get resource id
    RID() ResourceIDReader

    //Metadata ...
    Metadata() (MetadataReader, error)

    //Fields ...
    Fields() VectorFieldReader

}

//VectorResourceReader ...
type VectorResourceReader interface {

    // Returns the current size of this vector
    Len() int

    //Get the item in the position i, if i < Len(),
    //if item does not exist should return the default value for the underlying data type
    //when i > Len() should return an VectorInvalidIndexError
    Get(i int) (item ResourceReader, err error)
}

//ResourceShard ...
type ResourceShard interface {

    Metadata(func(MetadataShard)) ShardHolderAndDisposer
    Fields(func(FieldShard)) ShardHolderAndDisposer
}

//ResourceWildcardShard ...
func ResourceWildcardShard(s ResourceShard){

}

//ResourceForwardShard ...
func ResourceForwardShard(s ResourceShard){

    s.Metadata(MetadataForwardShard)
    s.Fields(FieldForwardShard)
}
