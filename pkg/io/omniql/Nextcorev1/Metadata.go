package corev1Hybrids



//MetadataReader meta data
type MetadataReader interface {

    //Application ...
    Application() string

    //Name ...
    Name() string

    //Kind ...
    Kind() string

    //Parent ...
    Parent() string

    //Documentation ...
    Documentation() DocumentationReader

}

type VectorMetadataReader interface {
     Len() int
     Get(i int) (item MetadataReader, err error)
}
