package corev1
import(
    "github.com/nebtex/hybrids/golang/hybrids"
    "strings"
)
//EnumerationTypes ...
type EnumerationTypes uint16

const (
    //EnumerationTypesNone ...
    EnumerationTypesNone EnumerationTypes = 0
    //Uint8 ...
    EnumerationTypesUint8 EnumerationTypes = 1
    //Uint16 ...
    EnumerationTypesUint16 EnumerationTypes = 2
    //Uint32 ...
    EnumerationTypesUint32 EnumerationTypes = 3
)

var enumeration_types_map map[EnumerationTypes]string
var enumeration_types_reverse_map map[string]EnumerationTypes


func init(){
	//init maps

    enumeration_types_map = map[EnumerationTypes]string{}
    enumeration_types_reverse_map = map[string]EnumerationTypes{}
    enumeration_types_map[EnumerationTypesUint8] = "Uint8"
    enumeration_types_reverse_map["Uint8"] = EnumerationTypesUint8
    enumeration_types_map[EnumerationTypesUint16] = "Uint16"
    enumeration_types_reverse_map["Uint16"] = EnumerationTypesUint16
    enumeration_types_map[EnumerationTypesUint32] = "Uint32"
    enumeration_types_reverse_map["Uint32"] = EnumerationTypesUint32


}

//String stringer
func (et EnumerationTypes) String() (value string) {
	value = enumeration_types_map[et]
	return
}

//IsValid check if the variable has a valid enumeration value
func (et EnumerationTypes) IsValid() (result bool) {
	_, result = enumeration_types_map[et]
	return
}



//EnumerationTypesFromString convert a string to its EnumerationTypes representation
func FromStringToEnumerationTypes(str string) (value EnumerationTypes) {
    var ok bool
    value, ok = enumeration_types_reverse_map[strings.Title(strings.TrimSpace(str))]
    if !ok{
        value = EnumerationTypesNone
	}
	return
}

//VectorEnumerationTypes ...
type VectorEnumerationTypes interface {

	// Returns the current size of this vector
	Len() int

	// Get the item in the position i, if i < Len(),
	// if item does not exist should return the default value for the underlying data type
	// when i > Len() should return an VectorInvalidIndexError
	Get(i int) (EnumerationTypes, error)
}

type vector_enumeration_types struct {
	_vector []EnumerationTypes
}

//Len Returns the current size of this vector
func (vet *vector_enumeration_types) Len() (size int) {
	size = len(vet._vector)
	return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
func (vet *vector_enumeration_types) Get(i int) (item EnumerationTypes, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vet._vector)}
		return
	}

	if i > len(vet._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vet._vector)}
		return
	}

	item = vet._vector[i]
	return

}

//NewVectorEnumerationTypes ...
func NewVectorEnumerationTypes(v []EnumerationTypes) VectorEnumerationTypes {
	return &vector_enumeration_types{_vector: v}
}
