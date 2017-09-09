package corev1

//MetadataReader ...
type MetadataReader interface {

    //Application ...
    Application() string

    //Name ...
    Name() string

    //Kind ...
    Kind() ApplicationType

    //Parent ...
    Parent() string

    //Documentation ...
    Documentation() (DocumentationReader, error)

}

//VectorMetadataReader ...
type VectorMetadataReader interface {

    // Returns the current size of this vector
    Len() int

    //Get the item in the position i, if i < Len(),
    //if item does not exist should return the default value for the underlying data type
    //when i > Len() should return an VectorInvalidIndexError
    Get(i int) (item MetadataReader, err error)
}

//MetadataShard ...
type MetadataShard interface {

    Application() ShardHolderAndDisposer
    Name() ShardHolderAndDisposer
    Kind() ShardHolderAndDisposer
    Parent() ShardHolderAndDisposer
    Documentation(func(DocumentationShard)) ShardHolderAndDisposer
}

//MetadataWildcardShard ...
func MetadataWildcardShard(s MetadataShard){

    s.Application()
    s.Name()
    s.Kind()
    s.Parent()
}

//MetadataForwardShard ...
func MetadataForwardShard(s MetadataShard){

    s.Application()
    s.Name()
    s.Kind()
    s.Parent()
    s.Documentation(DocumentationForwardShard)
}
