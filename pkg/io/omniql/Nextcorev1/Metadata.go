package Nextcorev1



//MetadataReader ...
type MetadataReader interface {

    //Application ...
    Application() string

    //Name ...
    Name() string

    //Kind ...
    Kind() ApplicationType

    //Parent ...
    Parent() string

    //Documentation ...
    Documentation() (DocumentationReader, error)

}

type VectorMetadataReader interface {
     Len() int
     Get(i int) (item MetadataReader, err error)
}
