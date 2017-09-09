package corev1


import(
    "github.com/nebtex/hybrids/golang/hybrids"
)
//UnionReader ...
type UnionReader interface {

    //RID get resource id
    RID() ResourceIDReader

    //Metadata ...
    Metadata() (MetadataReader, error)

    //Items ...
    Items() hybrids.VectorStringReader

}

//VectorUnionReader ...
type VectorUnionReader interface {

    // Returns the current size of this vector
    Len() int

    //Get the item in the position i, if i < Len(),
    //if item does not exist should return the default value for the underlying data type
    //when i > Len() should return an VectorInvalidIndexError
    Get(i int) (item UnionReader, err error)
}

//UnionShard ...
type UnionShard interface {

    Metadata(func(MetadataShard)) ShardHolderAndDisposer
    Items() ShardHolderAndDisposer
}

//UnionWildcardShard ...
func UnionWildcardShard(s UnionShard){

    s.Items()
}

//UnionForwardShard ...
func UnionForwardShard(s UnionShard){

    s.Metadata(MetadataForwardShard)
    s.Items()
}
