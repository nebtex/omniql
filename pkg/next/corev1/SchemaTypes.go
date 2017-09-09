package corev1
import(
    "github.com/nebtex/hybrids/golang/hybrids"
    "strings"
)
//SchemaTypes ...
type SchemaTypes uint16

const (
    //SchemaTypesNone ...
    SchemaTypesNone SchemaTypes = 0
    //Enumeration ...
    SchemaTypesEnumeration SchemaTypes = 1
    //Resource ...
    SchemaTypesResource SchemaTypes = 2
    //Table ...
    SchemaTypesTable SchemaTypes = 3
    //Union ...
    SchemaTypesUnion SchemaTypes = 4
    //EnumerationGroup ...
    SchemaTypesEnumerationGroup SchemaTypes = 5
    //EnumerationItem ...
    SchemaTypesEnumerationItem SchemaTypes = 6
    //UnionResource ...
    SchemaTypesUnionResource SchemaTypes = 7
    //UnionTable ...
    SchemaTypesUnionTable SchemaTypes = 8
    //Documentation ...
    SchemaTypesDocumentation SchemaTypes = 9
    //Field ...
    SchemaTypesField SchemaTypes = 10
    //Metadata ...
    SchemaTypesMetadata SchemaTypes = 11
)

var schema_types_map map[SchemaTypes]string
var schema_types_reverse_map map[string]SchemaTypes

var schema_types_resource_map map[SchemaTypes]bool

func init(){
	//init maps

    schema_types_map = map[SchemaTypes]string{}
    schema_types_reverse_map = map[string]SchemaTypes{}
    schema_types_map[SchemaTypesEnumeration] = "Enumeration"
    schema_types_reverse_map["Enumeration"] = SchemaTypesEnumeration
    schema_types_map[SchemaTypesResource] = "Resource"
    schema_types_reverse_map["Resource"] = SchemaTypesResource
    schema_types_map[SchemaTypesTable] = "Table"
    schema_types_reverse_map["Table"] = SchemaTypesTable
    schema_types_map[SchemaTypesUnion] = "Union"
    schema_types_reverse_map["Union"] = SchemaTypesUnion
    schema_types_map[SchemaTypesEnumerationGroup] = "EnumerationGroup"
    schema_types_reverse_map["EnumerationGroup"] = SchemaTypesEnumerationGroup
    schema_types_map[SchemaTypesEnumerationItem] = "EnumerationItem"
    schema_types_reverse_map["EnumerationItem"] = SchemaTypesEnumerationItem
    schema_types_map[SchemaTypesUnionResource] = "UnionResource"
    schema_types_reverse_map["UnionResource"] = SchemaTypesUnionResource
    schema_types_map[SchemaTypesUnionTable] = "UnionTable"
    schema_types_reverse_map["UnionTable"] = SchemaTypesUnionTable
    schema_types_map[SchemaTypesDocumentation] = "Documentation"
    schema_types_reverse_map["Documentation"] = SchemaTypesDocumentation
    schema_types_map[SchemaTypesField] = "Field"
    schema_types_reverse_map["Field"] = SchemaTypesField
    schema_types_map[SchemaTypesMetadata] = "Metadata"
    schema_types_reverse_map["Metadata"] = SchemaTypesMetadata


    schema_types_resource_map =  map[SchemaTypes]bool{}
    schema_types_resource_map[SchemaTypesTable] = true
    schema_types_resource_map[SchemaTypesUnion] = true
    schema_types_resource_map[SchemaTypesResource] = true
    schema_types_resource_map[SchemaTypesEnumeration] = true
}

//String stringer
func (st SchemaTypes) String() (value string) {
	value = schema_types_map[st]
	return
}

//IsValid check if the variable has a valid enumeration value
func (st SchemaTypes) IsValid() (result bool) {
	_, result = schema_types_map[st]
	return
}


//IsResource ...
func (st SchemaTypes) IsResource() (result bool) {
    _,result = schema_types_resource_map[st]
    return

}

//SchemaTypesFromString convert a string to its SchemaTypes representation
func FromStringToSchemaTypes(str string) (value SchemaTypes) {
    var ok bool
    value, ok = schema_types_reverse_map[strings.Title(strings.TrimSpace(str))]
    if !ok{
        value = SchemaTypesNone
	}
	return
}

//VectorSchemaTypes ...
type VectorSchemaTypes interface {

	// Returns the current size of this vector
	Len() int

	// Get the item in the position i, if i < Len(),
	// if item does not exist should return the default value for the underlying data type
	// when i > Len() should return an VectorInvalidIndexError
	Get(i int) (SchemaTypes, error)
}

type vector_schema_types struct {
	_vector []SchemaTypes
}

//Len Returns the current size of this vector
func (vst *vector_schema_types) Len() (size int) {
	size = len(vst._vector)
	return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
func (vst *vector_schema_types) Get(i int) (item SchemaTypes, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vst._vector)}
		return
	}

	if i > len(vst._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vst._vector)}
		return
	}

	item = vst._vector[i]
	return

}

//NewVectorSchemaTypes ...
func NewVectorSchemaTypes(v []SchemaTypes) VectorSchemaTypes {
	return &vector_schema_types{_vector: v}
}
