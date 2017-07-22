package Nextcorev1



//DocumentationReader documentation type
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
