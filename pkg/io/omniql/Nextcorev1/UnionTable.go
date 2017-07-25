package Nextcorev1

//UnionTableReader ...
type UnionTableReader interface {

}

//VectorUnionTableReader ...
type VectorUnionTableReader interface {

    // Returns the current size of this vector
    Len() int

    //Get the item in the position i, if i < Len(),
    //if item does not exist should return the default value for the underlying data type
    //when i > Len() should return an VectorInvalidIndexError
    Get(i int) (item UnionTableReader, err error)
}
