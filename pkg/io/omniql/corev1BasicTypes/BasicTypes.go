package corev1BasicTypes

type enumType byte

const (
	None            enumType = iota
	Byte
	UnsignedByte
	Boolean
	Short
	UnsignedShort
	Integer
	UnsignedInteger
	Float
	Long
	UnsignedLong
	Double
	Vector
	String
)

var scalars []enumType

func init() {
	scalars = []enumType{Byte,
			     UnsignedByte,
			     Boolean,
			     Short,
			     UnsignedShort,
			     Integer,
			     UnsignedInteger,
			     Float,
			     Long,
			     UnsignedLong,
			     Double}

}

func (e enumType) String() (value string) {
	switch e {
	case None:
		value = "None"
	case Byte:
		value = "Byte"
	}
	return
}

func (e enumType) IsScalar() (result bool) {
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

func (e enumType) Scalars() (result []enumType) {
	return scalars
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
