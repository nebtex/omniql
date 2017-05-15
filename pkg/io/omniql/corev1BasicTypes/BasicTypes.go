package corev1BasicTypes

type BasicTypes byte
type Scalars byte

const (
	//None ghjgkjgggg
	//hjhk jnknlnlknl kjjjjlkkn jkjkj
	//kklkkokoko jllklkjljljjljljjlkjljlkjllll
	None            BasicTypes = 1
	Byte            BasicTypes = 2
	UnsignedByte    BasicTypes = 3
	Boolean         BasicTypes = 4
	//Short
	//UnsignedShort
	//Integer
	//UnsignedInteger
	//Float
	//Long
	//UnsignedLong
	//Double
	//Vector
	//String
)


const (
	ScalarNone            Scalars = iota
	ScalarByte
	ScalarUnsignedByte
	ScalarBoolean
	ScalarShort
	ScalarUnsignedShort
	ScalarInteger
	ScalarUnsignedInteger
	ScalarFloat
	ScalarLong
	ScalarUnsignedLong
	ScalarDouble
)

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

func (e enumType) FromString(str string) (value enumType) {
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

func (e enumType) FromStrings(s []string) ([]enumType) {
	value := make([]enumType, len(s))
	for i := 0; i < len(s); i++ {
		value[i] = None.FromString(s[i])
	}
	return value
}
