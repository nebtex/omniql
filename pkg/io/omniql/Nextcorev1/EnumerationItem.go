package Nextcorev1



//EnumerationItemReader ...
type EnumerationItemReader interface {

    //Name ...
    Name() string

    //Documentation ...
    Documentation() (DocumentationReader, error)

}

//VectorEnumerationItemReader ...
type VectorEnumerationItemReader interface {

    // Returns the current size of this vector
    Len() int

    //Get the item in the position i, if i < Len(),
    //if item does not exist should return the default value for the underlying data type
    //when i > Len() should return an VectorInvalidIndexError
    Get(i int) (item EnumerationItemReader, err error)
}
