package Nextcorev1

type ResourceUnionReader interface {
	Kind() ApplicationType
	Resource() ResourceReader
	Enumeration() EnumerationReader
	Table() TableReader
}

type VectorResourceUnion interface {
	Get(i int) (ResourceUnionReader, error)
	Len() int
}

type CollectionQueryReader interface {
	Pattern()string
	Application() string
	Ref() string
	Pager()
}

type CollectionQueryWriter interface {

}

type CollectionResultReader interface {
	TotalCount() int64
	Objects() VectorResourceUnion
	PageInfo() string
}

type CollectionPager interface {
	Forward(int, int)
	Backward(int, int)
}
