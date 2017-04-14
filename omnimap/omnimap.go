package omnimap

//Structure ...

type Structure struct {
	//Localization in the parent buffer
	Index           int32
	path            []int32
	buffer          []byte
	ParentResource  *Resource
	ParentStructure *Structure
}

//Resource ...
type Resource struct {
	Index  int32
	buf    []byte
	path   []int64
	Field  map[string]Structure
}

//Map ...
type Map struct {
	Application string
	buffer      []byte
	Resources   map[string]Resource
}
