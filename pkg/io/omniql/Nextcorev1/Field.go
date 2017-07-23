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

type VectorFieldReader interface {
     Len() int
     Get(i int) (item FieldReader, err error)
}
