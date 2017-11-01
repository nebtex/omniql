package _go

type DNS interface {
}

type Resource interface {
}

type Computation interface {
	Call()
}

type Partition interface {
	Scan("ssd/ddd/**")

}

type tree struct{
	nodes []tree
}