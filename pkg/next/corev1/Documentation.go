package corev1

//DocumentationReader ...
type DocumentationReader interface {

    //Short ...
    Short() string

    //Long ...
    Long() string

}

//VectorDocumentationReader ...
type VectorDocumentationReader interface {

    // Returns the current size of this vector
    Len() int

    //Get the item in the position i, if i < Len(),
    //if item does not exist should return the default value for the underlying data type
    //when i > Len() should return an VectorInvalidIndexError
    Get(i int) (item DocumentationReader, err error)
}

//DocumentationShard ...
type DocumentationShard interface {

    Short() ShardHolderAndDisposer
    Long() ShardHolderAndDisposer
}

//DocumentationWildcardShard ...
func DocumentationWildcardShard(s DocumentationShard){

    s.Short()
    s.Long()
}

//DocumentationForwardShard ...
func DocumentationForwardShard(s DocumentationShard){

    s.Short()
    s.Long()
}
