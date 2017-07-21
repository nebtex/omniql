package corev1Hybrids



//EnumerationGroupReader allow to group enumerations
type EnumerationGroupReader interface {

    //Name ...
    Name() string

    //Documentation ...
    Documentation() DocumentationReader

    //Items ...
    Items() hybrids.VectorStringReader

}

type VectorEnumerationGroupReader interface {
     Len() int
     Get(i int) (item EnumerationGroupReader, err error)
}
