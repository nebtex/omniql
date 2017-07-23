package Nextcorev1



//DocumentationReader ...
type DocumentationReader interface {

    //Short ...
    Short() string

    //Long ...
    Long() string

}

type VectorDocumentationReader interface {
     Len() int
     Get(i int) (item DocumentationReader, err error)
}
