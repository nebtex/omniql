package Nextcorev1
import(
    "github.com/nebtex/hybrids/golang/hybrids"
)
//ApplicationType ...
type ApplicationType uint16

const (
    //ApplicationTypeNone ...
    ApplicationTypeNone ApplicationType = 0
    //Enumeration ...
    ApplicationTypeEnumeration ApplicationType = 1
    //Resource ...
    ApplicationTypeResource ApplicationType = 2
    //Table ...
    ApplicationTypeTable ApplicationType = 3
    //Union ...
    ApplicationTypeUnion ApplicationType = 4
    //EnumerationGroup ...
    ApplicationTypeEnumerationGroup ApplicationType = 5
    //EnumerationItem ...
    ApplicationTypeEnumerationItem ApplicationType = 6
    //UnionResource ...
    ApplicationTypeUnionResource ApplicationType = 7
    //UnionTable ...
    ApplicationTypeUnionTable ApplicationType = 8
    //Documentation ...
    ApplicationTypeDocumentation ApplicationType = 9
    //Field ...
    ApplicationTypeField ApplicationType = 10
    //Metadata ...
    ApplicationTypeMetadata ApplicationType = 11
)

var application_type_map map[ApplicationType]string
var application_type_reverse_map map[string]ApplicationType

var application_type_resource_map map[ApplicationType]bool

func init(){
	//init maps

    application_type_map = map[ApplicationType]string{}
    application_type_reverse_map = map[string]ApplicationType{}
    application_type_map[ApplicationTypeEnumeration] = "Enumeration"
    application_type_reverse_map["Enumeration"] = ApplicationTypeEnumeration
    application_type_map[ApplicationTypeResource] = "Resource"
    application_type_reverse_map["Resource"] = ApplicationTypeResource
    application_type_map[ApplicationTypeTable] = "Table"
    application_type_reverse_map["Table"] = ApplicationTypeTable
    application_type_map[ApplicationTypeUnion] = "Union"
    application_type_reverse_map["Union"] = ApplicationTypeUnion
    application_type_map[ApplicationTypeEnumerationGroup] = "EnumerationGroup"
    application_type_reverse_map["EnumerationGroup"] = ApplicationTypeEnumerationGroup
    application_type_map[ApplicationTypeEnumerationItem] = "EnumerationItem"
    application_type_reverse_map["EnumerationItem"] = ApplicationTypeEnumerationItem
    application_type_map[ApplicationTypeUnionResource] = "UnionResource"
    application_type_reverse_map["UnionResource"] = ApplicationTypeUnionResource
    application_type_map[ApplicationTypeUnionTable] = "UnionTable"
    application_type_reverse_map["UnionTable"] = ApplicationTypeUnionTable
    application_type_map[ApplicationTypeDocumentation] = "Documentation"
    application_type_reverse_map["Documentation"] = ApplicationTypeDocumentation
    application_type_map[ApplicationTypeField] = "Field"
    application_type_reverse_map["Field"] = ApplicationTypeField
    application_type_map[ApplicationTypeMetadata] = "Metadata"
    application_type_reverse_map["Metadata"] = ApplicationTypeMetadata


    application_type_resource_map =  map[ApplicationType]bool{}
    application_type_resource_map[ApplicationTypeTable] = true
    application_type_resource_map[ApplicationTypeUnion] = true
    application_type_resource_map[ApplicationTypeResource] = true
    application_type_resource_map[ApplicationTypeEnumeration] = true
}

//String stringer
func (at ApplicationType) String() (value string) {
	value = application_type_map[at]
	return
}

//IsValid check if the variable has a valid enumeration value
func (at ApplicationType) IsValid() (result bool) {
	_, result = application_type_map[at]
	return
}


//IsResource ...
func (at ApplicationType) IsResource() (result bool) {
    _,result = application_type_resource_map[at]
    return

}

//ApplicationTypeFromString convert a string to its ApplicationType representation
func FromStringToApplicationType(str string) (value ApplicationType) {
    var ok bool
    value, ok = application_type_reverse_map[str]
    if !ok{
        value = ApplicationTypeNone
	}
	return
}

//VectorApplicationType ...
type VectorApplicationType interface {

	// Returns the current size of this vector
	Len() int

	// Get the item in the position i, if i < Len(),
	// if item does not exist should return the default value for the underlying data type
	// when i > Len() should return an VectorInvalidIndexError
	Get(i int) (ApplicationType, error)
}

type vector_application_type struct {
	_vector []ApplicationType
}

//Len Returns the current size of this vector
func (vat *vector_application_type) Len() (size int) {
	size = len(vat._vector)
	return
}

//Get the item in the position i, if i < Len(),
//if item does not exist should return the default value for the underlying data type
//when i > Len() should return an VectorInvalidIndexError
func (vat *vector_application_type) Get(i int) (item ApplicationType, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vat._vector)}
		return
	}

	if i > len(vat._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vat._vector)}
		return
	}

	item = vat._vector[i]
	return

}

//NewVectorApplicationType ...
func NewVectorApplicationType(v []ApplicationType) VectorApplicationType {
	return &vector_application_type{_vector: v}
}
