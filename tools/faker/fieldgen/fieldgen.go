package fieldgen

import (
	"github.com/nebtex/omniql/commons/golang/oreflection"
	"github.com/nebtex/hybrids/golang/hybrids"
)

//go:generate mockery -name=EnumerationGenerator
type EnumerationGenerator interface {
	StringEnumeration(path string, fieldType oreflection.OType) (string, error)
	Uint8Enumeration(path string, fieldType oreflection.OType) (uint8, error)
	Uint16Enumeration(path string, fieldType oreflection.OType) (uint16, error)
	Uint32Enumeration(path string, fieldType oreflection.OType) (uint32, error)
	ShouldGenerateString(path string, fieldType oreflection.OType) (bool, error)
}

//go:generate mockery -name=FieldGenerator
//FieldGenerator generate random data using the reflection interface
type FieldGenerator interface {
	//generate random vector len
	VectorLen(path string, ot oreflection.OType) (int, error)

	//return if a field should be nil or not (don't use this on scalar or structs)
	ShouldBeNil(path string, ot oreflection.OType) (bool, error)

	//return a random Boolean tha follow the spec of th field
	Boolean(path string, field oreflection.OType) (bool, error)

	//return a random Int8 tha follow the spec of th field
	Int8(path string, field oreflection.OType) (int8, error)

	//return a random Uint8 tha follow the spec of th field
	Uint8(path string, field oreflection.OType) (uint8, error)

	//return a random Int16 tha follow the spec of th field
	Int16(path string, field oreflection.OType) (int16, error)

	//return a random Uint16 tha follow the spec of th field
	Uint16(path string, field oreflection.OType) (uint16, error)

	//return a random Int32 tha follow the spec of th field
	Int32(path string, field oreflection.OType) (int32, error)

	//return a random Uint32 tha follow the spec of th field
	Uint32(path string, field oreflection.OType) (uint32, error)

	//return a random Int64 tha follow the spec of th field
	Int64(path string, field oreflection.OType) (int64, error)

	//return a random Uint64 tha follow the spec of th field
	Uint64(path string, field oreflection.OType) (uint64, error)

	//return a random Float32 tha follow the spec of th field
	Float32(path string, field oreflection.OType) (float32, error)

	//return a random Float64 tha follow the spec of th field
	Float64(path string, field oreflection.OType) (float64, error)

	//return a random String tha follow the spec of th field
	String(path string, field oreflection.OType) (string, error)

	//return a random String tha follow the spec of th field
	Byte(path string, field oreflection.OType) ([]byte, error)

	//return a random String tha follow the spec of th field
	ResourceID(path string, field oreflection.OType, resourceIdType hybrids.ResourceIDType) ([]byte, error)

	//return a random Enumeration,
	Enumeration() EnumerationGenerator

	ShouldGenerateField(path string, table oreflection.OType, fn hybrids.FieldNumber) (bool, error)
}
