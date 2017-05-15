package corev1

import "github.com/nebtex/omnibuff/pkg/io/omniql/corev1BasicTypes"

type DocumentationReader interface {
	Short() string
	Long() string
}

type OqlIDValueReader interface {
	Field() FieldReader
	Resource() ResourceReader
}

type VectorFieldReader interface {
	Len() int32
	Get(index int32) (FieldReader, error)
}

type MetadataReader interface {
	Application() string
	Name() string
	Kind() string
	Parent() OqlIDReader
	Documentation() DocumentationReader
}

type ResourceReader interface {
	OqlID() OqlIDReader
	Fields() VectorFieldReader
	Meta() MetadataReader
}

type OqlIDReader interface {
	Type() int
	ID() int64
	Name() string
	Local() bool
	Value() OqlIDValueReader
	Parent() OqlIDReader
}

type FieldReader interface {
	OqlID() OqlIDReader
	Name() string
	Type() string
	Items() string
	Documentation() DocumentationReader
	Required() bool
}

type GroupReader interface {
	OqlID() OqlIDReader
	Name() string
	Documentation() DocumentationReader
}


type EnumerationItemReader interface {
	Name() string
	Groups() VectorEnumerationGroupReader
}
type FieldErrorReader interface {
	Error() string
	ErrorCode() int64
	Application() string
	Field() FieldReader
	Parent()
	ID() int64
}

type FieldError struct {
	Error string
}

type FieldErrorOptions interface {
	UserMessage(str string)
	ErrorCode(code int)
	//Todo //user can extends this
}

type EnumerationGroupReader interface {
	OqlID() OqlIDReader
	Name() string
	Documentation() DocumentationReader
}
type FieldErrorReaderOptionsSetter func(FieldErrorOptions)

type FieldErrorReaderConstructor interface {
	FromNative() (FieldErrorReader, error)
	FromReader(FieldReader, ...FieldErrorReaderOptionsSetter) (FieldErrorReader)
}

type FieldErrorReaderConstructorFromNative func(fe *FieldError) FieldErrorReader

type ErrorWriter interface {
	Field() FieldErrorReaderConstructor
}

type EnumerationReader interface {
	OqlID() OqlIDReader
	Meta() MetadataReader
	Kind() corev1BasicTypes.Scalars
}

type Supplier interface {
}

type OmniQLResource interface {
	OqlID() OqlIDReader
}
type VectorEnumerationGroupReader interface {
	Len() int32
	Get(i int32) EnumerationGroupReader
}

type UnTypedNodeApex interface {
	// 1.*
	MatchAny() UnTypedTypeApex
	// 1.**
	MatchAnyLevel()
	// 1.1
	MatchID(int64) UnTypedTypeApex
	// 1.(4|3)+
	MatchAnyOf(...int64) UnTypedTypeApex
	// 1.(5|3)-
	MatchNoneOf(...int64) UnTypedTypeApex
	// 1.>
	Forward()
}

type EnumerationGroupNodeApex interface {
	// 1.*
	MatchAny() EnumerationGroupTypeApex
	// 1.**
	MatchAnyLevel() EnumerationGroupTypeApex
	// 1.1
	MatchID(int64) EnumerationGroupTypeApex
	// 1.(4|3)+
	MatchAnyOf(...int64) EnumerationGroupTypeApex
	// 1.(5|3)-
	MatchNoneOf(...int64) EnumerationGroupTypeApex
	// 1.>
	Forward()
}

type UnTypedTypeApex interface {
	// 1.*
	MatchAny() UnTypedNodeApex
	// 1.**
	MatchAnyLevel()
	// 1.1
	MatchID(int64) UnTypedNodeApex
	// 1.(4|3)+
	MatchAnyOf(...int64) UnTypedNodeApex
	// 1.(5|3)-
	MatchNoneOf(...int64) UnTypedNodeApex
	// 1.>
	Forward()
	ResourceNode()
	EnumerationGroupApex() EnumerationGroupNodeApex
}

type EnumerationGroupTypeApex interface {
	// 1.*
	MatchAny() UnTypedNodeApex
	// 1.**
	MatchAnyLevel() UnTypedNodeApex
	// 1.1
	MatchID(int64) UnTypedNodeApex
	// 1.(4|3)+
	MatchAnyOf(...int64) UnTypedNodeApex
	// 1.(5|3)-
	MatchNoneOf(...int64) UnTypedNodeApex
	// 1.>
	QuerySet() EnumerationGroupQuerySet
	EnumerationGroupApex() EnumerationGroupNodeApex
}

type EnumerationGroupQuerySet interface {
	Result() EnumerationGroupQueryResult
}

type EnumerationGroupQueryResult interface {
	Apply(func(egr EnumerationGroupReader, applyErr error) (stop bool, funcErr error)) (ran bool, err error)
}

//create apex
type ApexMaker interface {
	FromID(OqlIDReader) UnTypedTypeApex
}

type EnumerationItemStreamReader interface {
	Name() string
	Groups() VectorEnumerationGroupReader
}


type QueryChunkStreamer interface {
	FieldChunkStream(func(FragmentProgressReader, ChangeTreeReader, Unsubscribe func(str string) error, FieldReader, queryError error) error)
}

type QueryUnitStreamer interface {
	SetChunkStreamer(QueryChunkStreamer)
	FieldUnitStream(func(ChangeTreeReader, FieldReader, queryError error) error)
}

type FieldStreamReader interface {
	OqlID() OqlIDReader
	Name() string
	Type() string
	Items() string
	Documentation() (DocumentationStreamReader, err error)
	Required() bool
}

type Result interface {
	FieldNodes() (result []FieldStreamReader, next func() []FieldStreamReader, err error)
	EnumerationItemNodes()
}

type SyncQueryResult interface {
	FieldSubscripntion(func(ChangeTreeReader, FieldReader, queryError error) error)
}
