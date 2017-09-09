package faker

//go:generate go run scalar-generator.go
type Json struct {
}

func iToBool(i int) (r bool) {
	if i == 0 {
		return
	}
	r = true
	return
}
