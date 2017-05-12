package corev1Table

type DocumentationReader interface {
	Short() string
	Long() string
}
type FieldReader interface {
	Name()
	Type() string
	Items() string
	Documentation() DocumentationReader
	Required() bool
}
