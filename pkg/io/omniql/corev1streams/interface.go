package corev1streams

import "github.com/nebtex/omnibuff/pkg/io/omniql/corev1"

type VectorInterfaceStreamReader interface {
	Len() int32
	Get(int32) (InterfaceStreamReader, error)
}

type VectorResourceStreamReader interface {
	Len() int32
	Get(int32) (ResourceStreamReader, error)
}

type VectorTableStreamReader interface {
	Len() int32
	Get(int32) (ResourceStreamReader, error)
}

type VectorStringStreamReader interface {
	Len() int32
	Get(int32) (string, error)
}
type VectorEnumerationItemStreamReader interface {
	Len() int32
	Get(int32) (EnumerationItemStreamReader, error)
}

type VectorFieldStreamReader interface {
	Len() int32
	Get(int32) (FieldStreamReader, error)
}

type VectorEnumerationGroupStreamReader interface {
	Len() int32
	Get(int32) (EnumerationGroupStreamReader, error)
}

type UnionInterfaceStreamReader interface {
	Items() (VectorInterfaceStreamReader, error)
}

type UnionResourceStreamReader interface {
	Items() (VectorResourceStreamReader, error)
}

type UnionTableStreamReader interface {
	Items() (VectorTableStreamReader, error)
}

type DocumentationStreamReader interface {
	Short() string
	Long() string
}

type MetadataStreamReader interface {
	Application() string
	Name() string
	Documentation() (DocumentationStreamReader, error)
	Groups() (VectorEnumerationGroupStreamReader, error)
}

type EnumerationStreamReader interface {
	Meta() (MetadataStreamReader, error)
	Kind()
	Items() (VectorEnumerationItemStreamReader, error)
}

type EnumerationGroupStreamReader interface {
	Documentation() (DocumentationStreamReader, error)
}

type EnumerationItemStreamReader interface {
	Name() string
	Groups() (ver VectorEnumerationGroupStreamReader, err error)
}

type FieldStreamReader interface {
	Name() string
	Type() string
	Required() bool
	Deprecated() bool
	Documentation() (DocumentationStreamReader, error)
	Default() string
}

type InterfaceStreamReader interface {
	Meta() (MetadataStreamReader, error)
	Fields() (VectorFieldStreamReader, error)
}

type ResourceStreamReader interface {
	Meta() (MetadataStreamReader, error)
	Fields() (VectorFieldStreamReader, error)
}

type TableStreamReader interface {
	Meta() (MetadataStreamReader, error)
	Fields() (VectorFieldStreamReader, error)
}

type UnionStreamReader interface {
	Meta() (MetadataStreamReader, error)
	Type() (MetadataStreamReader, error)
}

type FieldChunkStreamObserver func(FragmentProgressReader, ChangeTreeReader, Unsubscribe func() error, fr corev1.FieldReader, queryError error) error

//when all the subscription are closed the ObserveFieldChunkStream will advice the api about
type NodeFieldChunkStreamReader interface {
	ObserveFieldChunkStream(FieldChunkStreamObserver)
}

type NodeFieldResultStreamReader interface {
	FieldStreamResultSet() (fs []FieldStreamReader, requested bool, err error)
}

type EnumerationItemChunkStreamObserver func(FragmentProgressReader, ChangeTreeReader, Unsubscribe func() error, fr corev1.EnumerationItemReader, queryError error) error

//when all the subscription are closed the ObserveFieldChunkStream will advice the api about
type NodeEnumerationItemChunkStreamReader interface {
	ObserveEnumerationItemChunkStream(EnumerationItemChunkStreamObserver)
}

type NodeEnumerationItemStreamReader interface {
	EnumerationItemStreamResultSet() (esr []EnumerationStreamReader, requested bool, err error)
}

type EnumerationGroupChunkStreamObserver func(FragmentProgressReader, ChangeTreeReader, Unsubscribe func() error, fr corev1.EnumerationGroupReader, queryError error) error

//when all the subscription are closed the ObserveFieldChunkStream will notify the api about
type NodeEnumerationGroupChunkStreamReader interface {
	ObserveEnumerationGroupChunkStream(EnumerationGroupChunkStreamObserver)
}

type NodeEnumerationGroupStreamReader interface {
	EnumerationGroupStreamResultSet() (esr []NodeEnumerationGroupStreamReader, requested bool, err error)
}

type NodeResultChunkStreamReader interface {
	NodeFieldResultStreamReader
	NodeEnumerationItemChunkStreamReader
	NodeEnumerationGroupStreamReader
}

type ResultChunkStreamReader interface {
	NodeResultChunkStreamReader
}


type ResultSetChunkSubscriber interface {
	ObserveChunks() ResultChunkStreamReader
}

type QuerySetRequestCUnitSubscriber interface {
	ObserveUnits() ResultChunkStreamReader
}

type QuerySetRequestConsumer interface {
	Consume() ResultChunkStreamReader
}

type QuerySetRequest interface {
	ResultSetChunkSubscriber
	QuerySetRequestCUnitSubscriber
	QuerySetRequestConsumer
}

type QueryWriter interface {
	StartNodeEnumerationItem(NodeEnumearionItemApex) NodeEnumerationQueryWriter
	Execute() (QuerySetRequest, error)
}

type NodeEnumerationQueryWriter interface {
	End() QueryWriter
}
