package Nextcorev1



//UnionResourceReader ...
type UnionResourceReader interface {

}

type VectorUnionResourceReader interface {
     Len() int
     Get(i int) (item UnionResourceReader, err error)
}
