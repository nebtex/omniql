package Nextcorev1

//ResourceIDReader ...
type ResourceIDReader interface {
	//Application ...
	Application() string

	//Kind ...
	Kind() ApplicationType

	//ID ...
	ID() string

	//Parent ...
	Parent() (ResourceIDReader, error)
}

//VectorResourceIDReader ...
type VectorResourceIDReader interface {
	// Returns the current size of this vector
	Len() int

	//Get the item in the position i, if i < Len(),
	//if item does not exist should return the default value for the underlying data type
	//when i > Len() should return an VectorInvalidIndexError
	Get(i int) (item ResourceIDReader, err error)
}

//ResourceIDShard ...
type ResourceIDShard interface {
	Application() ShardHolderAndDisposer
	Kind() ShardHolderAndDisposer
	ID() ShardHolderAndDisposer
	Parent(func(ResourceIDShard)) ShardHolderAndDisposer
}

//ResourceIDWildcardShard ...
func ResourceIDWildcardShard(s ResourceIDShard) {

	s.Application()
	s.Kind()
	s.ID()
}

//ResourceIDForwardShard ...
func ResourceIDForwardShard(s ResourceIDShard) {

	s.Application()
	s.Kind()
	s.ID()
	s.Parent(ResourceIDForwardShard)
}
