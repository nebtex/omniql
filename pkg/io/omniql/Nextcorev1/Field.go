package corev1Hybrids



//FieldReader field type
type FieldReader interface {

    //Name ...
    Name() string

    //Type ...
    Type() string

    //Documentation ...
    Documentation() DocumentationReader

    //Default String representation of the default value
    Default() string

}

type VectorFieldReader interface {
     Len() int
     Get(i int) (item FieldReader, err error)
}
