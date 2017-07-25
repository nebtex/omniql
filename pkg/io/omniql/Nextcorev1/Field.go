package Nextcorev1

//FieldReader ...
type FieldReader interface {

    //Name ...
    Name() string

    //Type ...
    Type() string

    //Documentation ...
    Documentation() (DocumentationReader, error)

    //Default String representation of the default value
    Default() string

}

//VectorFieldReader ...
type VectorFieldReader interface {

    // Returns the current size of this vector
    Len() int

    //Get the item in the position i, if i < Len(),
    //if item does not exist should return the default value for the underlying data type
    //when i > Len() should return an VectorInvalidIndexError
    Get(i int) (item FieldReader, err error)
}
