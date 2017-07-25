package Nextcorev1
import(
    "github.com/nebtex/hybrids/golang/hybrids"
    "strings"
)
//BasicType ...
type BasicType int8

const (
    //BasicTypeNone ...
    BasicTypeNone BasicType = 0
    //Integer8 ...
    BasicTypeInteger8 BasicType = 1
    //UnsignedInteger8 ...
    BasicTypeUnsignedInteger8 BasicType = 2
    //Boolean ...
    BasicTypeBoolean BasicType = 3
    //Short ...
    BasicTypeShort BasicType = 4
    //UnsignedShort ...
    BasicTypeUnsignedShort BasicType = 5
    //Integer ...
    BasicTypeInteger BasicType = 6
    //UnsignedInteger ...
    BasicTypeUnsignedInteger BasicType = 7
    //Float ...
    BasicTypeFloat BasicType = 8
    //Long ...
    BasicTypeLong BasicType = 9
    //UnsignedLong ...
    BasicTypeUnsignedLong BasicType = 10
    //Double ...
    BasicTypeDouble BasicType = 11
    //Vector ...
    BasicTypeVector BasicType = 12
    //String ...
    BasicTypeString BasicType = 13
    //Bytes ...
    BasicTypeBytes BasicType = 14
)

var basic_type_map map[BasicType]string
var basic_type_reverse_map map[string]BasicType

var basic_type_scalar_map map[BasicType]bool

func init(){
	//init maps

    basic_type_map = map[BasicType]string{}
    basic_type_reverse_map = map[string]BasicType{}
    basic_type_map[BasicTypeInteger8] = "Integer8"
    basic_type_reverse_map["Integer8"] = BasicTypeInteger8
    basic_type_map[BasicTypeUnsignedInteger8] = "UnsignedInteger8"
    basic_type_reverse_map["UnsignedInteger8"] = BasicTypeUnsignedInteger8
    basic_type_map[BasicTypeBoolean] = "Boolean"
    basic_type_reverse_map["Boolean"] = BasicTypeBoolean
    basic_type_map[BasicTypeShort] = "Short"
    basic_type_reverse_map["Short"] = BasicTypeShort
    basic_type_map[BasicTypeUnsignedShort] = "UnsignedShort"
    basic_type_reverse_map["UnsignedShort"] = BasicTypeUnsignedShort
    basic_type_map[BasicTypeInteger] = "Integer"
    basic_type_reverse_map["Integer"] = BasicTypeInteger
    basic_type_map[BasicTypeUnsignedInteger] = "UnsignedInteger"
    basic_type_reverse_map["UnsignedInteger"] = BasicTypeUnsignedInteger
    basic_type_map[BasicTypeFloat] = "Float"
    basic_type_reverse_map["Float"] = BasicTypeFloat
    basic_type_map[BasicTypeLong] = "Long"
    basic_type_reverse_map["Long"] = BasicTypeLong
    basic_type_map[BasicTypeUnsignedLong] = "UnsignedLong"
    basic_type_reverse_map["UnsignedLong"] = BasicTypeUnsignedLong
    basic_type_map[BasicTypeDouble] = "Double"
    basic_type_reverse_map["Double"] = BasicTypeDouble
    basic_type_map[BasicTypeVector] = "Vector"
    basic_type_reverse_map["Vector"] = BasicTypeVector
    basic_type_map[BasicTypeString] = "String"
    basic_type_reverse_map["String"] = BasicTypeString
    basic_type_map[BasicTypeBytes] = "Bytes"
    basic_type_reverse_map["Bytes"] = BasicTypeBytes


    basic_type_scalar_map =  map[BasicType]bool{}
    basic_type_scalar_map[BasicTypeInteger8] = true
    basic_type_scalar_map[BasicTypeUnsignedInteger8] = true
    basic_type_scalar_map[BasicTypeBoolean] = true
    basic_type_scalar_map[BasicTypeShort] = true
    basic_type_scalar_map[BasicTypeUnsignedShort] = true
    basic_type_scalar_map[BasicTypeInteger] = true
    basic_type_scalar_map[BasicTypeUnsignedInteger] = true
    basic_type_scalar_map[BasicTypeFloat] = true
    basic_type_scalar_map[BasicTypeLong] = true
    basic_type_scalar_map[BasicTypeUnsignedLong] = true
    basic_type_scalar_map[BasicTypeDouble] = true
}

//String stringer
func (bt BasicType) String() (value string) {
	value = basic_type_map[bt]
	return
}

//IsValid check if the variable has a valid enumeration value
func (bt BasicType) IsValid() (result bool) {
	_, result = basic_type_map[bt]
	return
}


//IsScalar ...
func (bt BasicType) IsScalar() (result bool) {
    _,result = basic_type_scalar_map[bt]
    return

}

//BasicTypeFromString convert a string to its BasicType representation
func FromStringToBasicType(str string) (value BasicType) {
    var ok bool
    value, ok = basic_type_reverse_map[strings.Title(strings.TrimSpace(str))]
    if !ok{
        value = BasicTypeNone
	}
	return
}

//VectorBasicType ...
type VectorBasicType interface {

	// Returns the current size of this vector
	Len() int

	// Get the item in the position i, if i < Len(),
	// if item does not exist should return the default value for the underlying data type
	// when i > Len() should return an VectorInvalidIndexError
	Get(i int) (BasicType, error)
}

type vector_basic_type struct {
	_vector []BasicType
}

//Len Returns the current size of this vector
func (vbt *vector_basic_type) Len() (size int) {
	size = len(vbt._vector)
	return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
func (vbt *vector_basic_type) Get(i int) (item BasicType, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vbt._vector)}
		return
	}

	if i > len(vbt._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vbt._vector)}
		return
	}

	item = vbt._vector[i]
	return

}

//NewVectorBasicType ...
func NewVectorBasicType(v []BasicType) VectorBasicType {
	return &vector_basic_type{_vector: v}
}
