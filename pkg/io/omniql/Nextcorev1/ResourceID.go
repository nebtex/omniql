package corev1Hybrids

type ResourceIDReader interface {
	Application() string
	Kind() string
	ID() string
	IsLocal() bool
	Parent() ResourceIDReader
}