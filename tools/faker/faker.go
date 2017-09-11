package faker

import "github.com/nebtex/omniql/tools/faker/fieldgen"

//go:generate go run scalar-generator.go
type Json struct {
	fieldGen fieldgen.FieldGenerator
}