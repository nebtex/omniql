package corev1

//ShardHolder ...
type ShardHolder interface {
    Hold(val bool)
}


//ShardDisposer ...
type ShardDisposer interface {
	Dispose(val bool)
}

//ShardHolderAndDisposer ...
type ShardHolderAndDisposer interface {
	ShardHolder
	ShardDisposer
}
	