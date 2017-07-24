package Nextcorev1

type ResourceIDReader interface {
	Application() string
	Kind() ApplicationType
	ID() string
	IsLocal() bool
	Parent() ResourceIDReader
}
