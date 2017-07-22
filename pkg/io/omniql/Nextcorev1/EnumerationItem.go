package Nextcorev1



//EnumerationItemReader ...
type EnumerationItemReader interface {

    //Name ...
    Name() string

    //Documentation ...
    Documentation() (DocumentationReader, error)

}

type VectorEnumerationItemReader interface {
     Len() int
     Get(i int) (item EnumerationItemReader, err error)
}
