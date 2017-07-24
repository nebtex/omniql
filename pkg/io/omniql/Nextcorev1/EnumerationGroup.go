package Nextcorev1


import(
    "github.com/nebtex/hybrids/golang/hybrids"
)
//EnumerationGroupReader ...
type EnumerationGroupReader interface {

    //Name ...
    Name() string

    //Documentation ...
    Documentation() (DocumentationReader, error)

    //Items ...
    Items() hybrids.VectorStringReader

}

//VectorEnumerationGroupReader ...
type VectorEnumerationGroupReader interface {

    // Returns the current size of this vector
    Len() int

    //Get the item in the position i, if i < Len(),
    //if item does not exist should return the default value for the underlying data type
    //when i > Len() should return an VectorInvalidIndexError
    Get(i int) (item EnumerationGroupReader, err error)
}
