package corev1
import(
    "github.com/nebtex/hybrids/golang/hybrids"
    "strings"
)
//UnionTypes ...
type UnionTypes uint16

const (
    //UnionTypesNone ...
    UnionTypesNone UnionTypes = 0
    //Table ...
    UnionTypesTable UnionTypes = 1
    //Resources ...
    UnionTypesResources UnionTypes = 2
    //ExternalResources ...
    UnionTypesExternalResources UnionTypes = 3
)

var union_types_map map[UnionTypes]string
var union_types_reverse_map map[string]UnionTypes


func init(){
	//init maps

    union_types_map = map[UnionTypes]string{}
    union_types_reverse_map = map[string]UnionTypes{}
    union_types_map[UnionTypesTable] = "Table"
    union_types_reverse_map["Table"] = UnionTypesTable
    union_types_map[UnionTypesResources] = "Resources"
    union_types_reverse_map["Resources"] = UnionTypesResources
    union_types_map[UnionTypesExternalResources] = "ExternalResources"
    union_types_reverse_map["ExternalResources"] = UnionTypesExternalResources


}

//String stringer
func (ut UnionTypes) String() (value string) {
	value = union_types_map[ut]
	return
}

//IsValid check if the variable has a valid enumeration value
func (ut UnionTypes) IsValid() (result bool) {
	_, result = union_types_map[ut]
	return
}



//UnionTypesFromString convert a string to its UnionTypes representation
func FromStringToUnionTypes(str string) (value UnionTypes) {
    var ok bool
    value, ok = union_types_reverse_map[strings.Title(strings.TrimSpace(str))]
    if !ok{
        value = UnionTypesNone
	}
	return
}

//VectorUnionTypes ...
type VectorUnionTypes interface {

	// Returns the current size of this vector
	Len() int

	// Get the item in the position i, if i < Len(),
	// if item does not exist should return the default value for the underlying data type
	// when i > Len() should return an VectorInvalidIndexError
	Get(i int) (UnionTypes, error)
}

type vector_union_types struct {
	_vector []UnionTypes
}

//Len Returns the current size of this vector
func (vut *vector_union_types) Len() (size int) {
	size = len(vut._vector)
	return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
func (vut *vector_union_types) Get(i int) (item UnionTypes, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vut._vector)}
		return
	}

	if i > len(vut._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vut._vector)}
		return
	}

	item = vut._vector[i]
	return

}

//NewVectorUnionTypes ...
func NewVectorUnionTypes(v []UnionTypes) VectorUnionTypes {
	return &vector_union_types{_vector: v}
}
