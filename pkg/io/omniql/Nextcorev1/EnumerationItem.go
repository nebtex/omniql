package corev1Hybrids



//EnumerationItemReader ...
type EnumerationItemReader interface {

    //Name ...
    Name() string

    //Documentation ...
    Documentation() DocumentationReader

}

type VectorEnumerationItemReader interface {
     Len() int
     Get(i int) (item EnumerationItemReader, err error)
}
