package faker

import (
	"github.com/nebtex/omniql/commons/golang/oreflection"
	"github.com/nebtex/hybrids/golang/hybrids"
)

//go:generate mockery -name=FieldGenerator

type FieldGenerator interface {
	//generate random vector len
	VectorLen(path string, ot oreflection.OType, fn hybrids.FieldNumber) (int, error)

	//return if a field should be nil or not (don't use this on scalar or structs)
	ShouldBeNil(path string, ot oreflection.OType, fn hybrids.FieldNumber) (bool, error)

	//return a random Boolean tha follow the spec of th table
	Boolean(path string, table oreflection.OType, fn hybrids.FieldNumber) (bool, error)

	//return a random Int8 tha follow the spec of th table
	Int8(path string, table oreflection.OType, fn hybrids.FieldNumber) (int8, error)

	//return a random Uint8 tha follow the spec of th table
	Uint8(path string, table oreflection.OType, fn hybrids.FieldNumber) (uint8, error)

	//return a random Int16 tha follow the spec of th table
	Int16(path string, table oreflection.OType, fn hybrids.FieldNumber) (int16, error)

	//return a random Uint16 tha follow the spec of th table
	Uint16(path string, table oreflection.OType, fn hybrids.FieldNumber) (uint16, error)

	//return a random Int32 tha follow the spec of th table
	Int32(path string, table oreflection.OType, fn hybrids.FieldNumber) (int32, error)

	//return a random Uint32 tha follow the spec of th table
	Uint32(path string, table oreflection.OType, fn hybrids.FieldNumber) (uint32, error)

	//return a random Int64 tha follow the spec of th table
	Int64(path string, table oreflection.OType, fn hybrids.FieldNumber) (int64, error)

	//return a random Uint64 tha follow the spec of th table
	Uint64(path string, table oreflection.OType, fn hybrids.FieldNumber) (uint64, error)

	//return a random Float32 tha follow the spec of th table
	Float32(path string, table oreflection.OType, fn hybrids.FieldNumber) (float32, error)

	//return a random Float64 tha follow the spec of th table
	Float64(path string, table oreflection.OType, fn hybrids.FieldNumber) (float64, error)

	//return a random String tha follow the spec of th table
	String(path string, table oreflection.OType, fn hybrids.FieldNumber) (string, error)

	//return a random String tha follow the spec of th table
	Byte(path string, table oreflection.OType, fn hybrids.FieldNumber) ([]byte, error)

	//return a random String tha follow the spec of th table
	ResourceID(path string, table oreflection.OType, resourceIdType hybrids.ResourceIDType, fn hybrids.FieldNumber) ([]byte, error)
}
