package Nextcorev1

//UnionResourceReader ...
type UnionResourceReader interface {
}

//VectorUnionResourceReader ...
type VectorUnionResourceReader interface {
	// Returns the current size of this vector
	Len() int

	//Get the item in the position i, if i < Len(),
	//if item does not exist should return the default value for the underlying data type
	//when i > Len() should return an VectorInvalidIndexError
	Get(i int) (item UnionResourceReader, err error)
}

//UnionResourceShard ...
type UnionResourceShard interface {
}

//UnionResourceWildcardShard ...
func UnionResourceWildcardShard(s UnionResourceShard) {

}

//UnionResourceForwardShard ...
func UnionResourceForwardShard(s UnionResourceShard) {

}
