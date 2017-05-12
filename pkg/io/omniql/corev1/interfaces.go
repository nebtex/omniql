package corev1

type DocumentationReader interface {
	Short() string
	Long() string
}
type FieldReader interface {
	Name() string
	Type() string
	Items() string
	Documentation() DocumentationReader
	Required() bool
}
