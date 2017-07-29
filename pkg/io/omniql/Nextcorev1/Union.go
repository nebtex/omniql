package Nextcorev1

//UnionReader ...
type UnionReader interface {
	//RID get resource id
	RID() ResourceIDReader

	//Metadata ...
	Metadata() (MetadataReader, error)
}

//vectorunionreader ...
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
}

//UnionWildcardShard ...
func UnionWildcardShard(s UnionShard) {

}

//UnionForwardShard ...
func UnionForwardShard(s UnionShard) {

	s.Metadata(MetadataForwardShard)
}
