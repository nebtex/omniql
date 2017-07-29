package Nextcorev1

//TableReader ...
type TableReader interface {
	//RID get resource id
	RID() ResourceIDReader

	//Metadata ...
	Metadata() (MetadataReader, error)

	//Fields ...
	Fields() VectorFieldReader
}

//VectorTableReader ...
type VectorTableReader interface {
	// Returns the current size of this vector
	Len() int

	//Get the item in the position i, if i < Len(),
	//if item does not exist should return the default value for the underlying data type
	//when i > Len() should return an VectorInvalidIndexError
	Get(i int) (item TableReader, err error)
}

//TableShard ...
type TableShard interface {
	Metadata(func(MetadataShard)) ShardHolderAndDisposer
	Fields(func(FieldShard)) ShardHolderAndDisposer
}

//TableWildcardShard ...
func TableWildcardShard(s TableShard) {

}

//TableForwardShard ...
func TableForwardShard(s TableShard) {

	s.Metadata(MetadataForwardShard)
	s.Fields(FieldForwardShard)
}
