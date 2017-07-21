package corev1Hybrids



//UnionTableReader ...
type UnionTableReader interface {

}

type VectorUnionTableReader interface {
     Len() int
     Get(i int) (item UnionTableReader, err error)
}
