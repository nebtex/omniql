package faker

//go:generate go run scalar-generator.go
type Json struct {
	fieldGen FieldGenerator
}

func iToBool(i int) (r bool) {
	if i == 0 {
		return
	}
	r = true
	return
}
