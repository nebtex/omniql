package oreflection

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"github.com/nebtex/omniql/pkg/next/corev1"
)

//go:generate mockery -name=LookupFields
//LookupFields ...
type LookupFields interface {
	ByPosition(fn hybrids.FieldNumber) (Field, ok bool)
	BySnakeCase(fieldName string) (f Field, ok bool)
	ByCamelCase(fieldName string) (f Field, ok bool)
}

//go:generate mockery -name=LookupEnumeration
//LookupEnumeration ...
type LookupEnumeration interface {
	ByUint8ToCamelCase(input uint8) (value string, ok bool)
	ByUint16ToCamelCase(input uint16) (value string, ok bool)
	ByUint32ToCamelCase(input uint32) (value string, ok bool)
	ByUint8ToSnakeCase(input uint8) (value string, ok bool)
	ByUint16ToSnakeCase(input uint16) (value string, ok bool)
	ByUint32ToSnakeCase(input uint32) (value string, ok bool)
	ByStringToUint8(input string) (value uint8, ok bool)
	ByStringToUint16(input string) (value uint16, ok bool)
	ByStringToUint32(input string) (value uint32, ok bool)
}

//go:generate mockery -name=LookupTableOnUnion
//LookupTableOnUnion ...
type LookupTableOnUnion interface {
	ByString(input string) (position hybrids.UnionKind, tableType OType, ok bool)
	ByPositionToCamelCase(input hybrids.UnionKind) (value string, tableType OType, ok bool)
	ByPositionToSnakeCase(input hybrids.UnionKind) (value string, tableType OType, ok bool)
}

//go:generate mockery -name=Table
//Table ...
type Table interface {
	Schema() corev1.TableReader
	FieldCount() hybrids.FieldNumber
	LookupFields() LookupFields
}

//go:generate mockery -name=Resource
//Resource ...
type Resource interface {
	Table() Table
	ResourceIDType() hybrids.ResourceIDType
}

//go:generate mockery -name=Struct
//Struct ...
type Struct interface {
	Id() string
	Package() string
	FieldCount() hybrids.FieldNumber
	LookupFields() LookupFields
}

//go:generate mockery -name=Items
//Items ...
type Items interface {
	ValueType() OType
	HybridType() hybrids.Types
}

//go:generate mockery -name=Field
//Field ...
type Field interface {
	Id() string //full id of this field, with the version, to allow query the backend

	GlobalPackage() string
	InnerPackage() string

	//table or struct
	Parent() OType

	//Otype of the field value
	//this return nil when is a scalar or vector scalar  unless tha the field is an enumeration enumeration
	ValueType() OType

	//position of this field in the table
	Position() hybrids.FieldNumber

	//field name
	Name() string

	//the underlying data type
	HybridType() hybrids.Types

	//if is a vector the item type
	Items() Items
}

//go:generate mockery -name=Enumeration
//Enumeration ...
type Enumeration interface {
	Lookup() LookupEnumeration
	//the underlying data type
	HybridType() hybrids.Types
}

//go:generate mockery -name=Union
//Union ...
type Union interface {
	LookupTable() LookupTableOnUnion
}

//go:generate mockery -name=Application
//Application ...
type Application interface {
	Name() string
	ResourceList() (ResourceKind uint16, oType OType)
}

//go:generate mockery -name=OType
//OType represent any omniql type
type OType interface {
	//Table, Field, Enumeration, etc..
	Kind() corev1.SchemaTypes

	//if this is a vector this will return the item type
	Enumeration() Enumeration
	Union() Union
	Table() Table
	Resource() Resource
	Field() Field
	Struct() Struct
	Application() Application
}
