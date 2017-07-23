package Nextcorev1

type ResourceIDReader interface {
	Application() string
	Kind() string
	ID() string
	IsLocal() bool
	Parent() ResourceIDReader
}
