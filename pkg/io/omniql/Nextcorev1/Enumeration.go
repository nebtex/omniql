package Nextcorev1

//EnumerationReader ...
type EnumerationReader interface {
	//RID get resource id
	RID() ResourceIDReader

	//Metadata ...
	Metadata() (MetadataReader, error)

	//Kind ...
	Kind() BasicType

	//Items ...
	Items() VectorEnumerationItemReader

	//Groups ...
	Groups() VectorEnumerationGroupReader
}

//VectorEnumerationReader ...
type VectorEnumerationReader interface {
	// Returns the current size of this VECTOR
	Len() int

	//Get the item in the position i, if i < Len(),
	//if item does not exist should return the default value for the underlying data type
	//when i > Len() should return an VectorInvalidIndexError
	Get(i int) (item EnumerationReader, err error)
}

//EnumerationShard ...
type EnumerationShard interface {
	Metadata(func(MetadataShard)) ShardHolderAndDisposer
	Kind() ShardHolderAndDisposer
	Items(func(EnumerationItemShard)) ShardHolderAndDisposer
	Groups(func(EnumerationGroupShard)) ShardHolderAndDisposer
}

//EnumerationWildcardShard ...
func EnumerationWildcardShard(s EnumerationShard) {

	s.Kind()
}

//EnumerationForwardShard ...
func EnumerationForwardShard(s EnumerationShard) {

	s.Metadata(MetadataForwardShard)
	s.Kind()
	s.Items(EnumerationItemForwardShard)
	s.Groups(EnumerationGroupForwardShard)
}
