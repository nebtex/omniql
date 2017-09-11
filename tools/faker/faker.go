package faker

import "github.com/nebtex/omniql/tools/faker/fieldgen"

//go:generate go run scalar-generator.go
type Json struct {
	fieldGen fieldgen.FieldGenerator
}

func iToBool(i int) (r bool) {
	if i == 0 {
		return
	}
	r = true
	return
}
