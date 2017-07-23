package corev1BasicTypes
//Doc
type BasicTypes byte

const (
	//Dodc
	BasicTypesNone            BasicTypes = iota
	BasicTypesByte
	BasicTypesUnsignedByte
	BasicTypesBoolean
	BasicTypesShort
	BasicTypesUnsignedShort
	BasicTypesInteger
	BasicTypesUnsignedInteger
	BasicTypesFloat
	BasicTypesLong
	BasicTypesUnsignedLong
	BasicTypesDouble
	BasicTypesVector
	BasicTypesString
)


func init(){
	init maps
}

//stringer
func (e BasicTypes) String() (value string) {
	switch e {
	case None:
		value = "None"
	case Byte:
		value = "Byte"
	}
	return
}

//

func (e BasicTypes) IsValid() (result bool) {
	switch e {
	case Byte,
		UnsignedByte,
		Boolean,
		Short,
		UnsignedShort,
		Integer,
		UnsignedInteger,
		Float,
		Long,
		UnsignedLong,
		Double:
		result = true
	}
	return

}
func (e BasicTypes) IsScalar() (result bool) {
	switch e {
	case Byte,
		UnsignedByte,
		Boolean,
		Short,
		UnsignedShort,
		Integer,
		UnsignedInteger,
		Float,
		Long,
		UnsignedLong,
		Double:
		result = true
	}
	return
}

func BasicTypesFromString(str string) (value BasicTypes) {
	value = None

	switch str {
	case "Byte":
		value = Byte
	case "UnsignedByte":
		value = UnsignedByte
	case "Boolean":
		value = Boolean
	case "Short":
		value = Short
	case "UnsignedShort":
		value = UnsignedShort
	case "Integer":
		value = Integer
	case "UnsignedInteger":
		value = UnsignedInteger
	case "Float":
		value = Float
	case "Long":
		value = Long
	case "UnsignedLong":
		value = UnsignedLong
	case "Double":
		value = Double

	}
	return
}
