package Nextcorev1



//EnumerationGroupReader allow to group enumerations
type EnumerationGroupReader interface {

    //Name ...
    Name() string

    //Documentation ...
    Documentation() (DocumentationReader, error)

    //Items ...
    Items() hybrids.VectorStringReader

}

type VectorEnumerationGroupReader interface {
     Len() int
     Get(i int) (item EnumerationGroupReader, err error)
}
