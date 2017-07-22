package Nextcorev1



//UnionTableReader ...
type UnionTableReader interface {

}

type VectorUnionTableReader interface {
     Len() int
     Get(i int) (item UnionTableReader, err error)
}
