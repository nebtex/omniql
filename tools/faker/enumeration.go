package faker
/*
import (
	"github.com/nebtex/omniql/commons/golang/oreflection"
	"github.com/nebtex/hybrids/golang/hybrids"
	"github.com/jmcvetta/randutil"
)

func (j *Json) fakeEnum(out map[string]interface{}, fieldName string, ot oreflection.OType) (err error) {
	var enum string
	var ok bool
	maxEnum := ot.Enumeration().Items().Len()

	value, err := randutil.IntRange(0, maxEnum)

	if err != nil {
		return err
	}

	switch ot.HybridType() {

	case hybrids.Uint8:
		enum, ok = ot.LookupEnumeration().ByUint8ToCamelCase(uint8(value))

		if ok {
			//write the enum as string
			out[fieldName] = enum
		} else {
			//write the enum as json number
			out[fieldName] = float64(value)
		}
		return

	case hybrids.Uint16:
		enum, ok = ot.LookupEnumeration().ByUint16ToCamelCase(uint16(value))

		if ok {
			//write the enum as string
			out[fieldName] = enum
		} else {
			//write the enum as json number
			out[fieldName] = float64(value)
		}
		return

	case hybrids.Uint32:
		enum, ok = ot.LookupEnumeration().ByUint32ToCamelCase(uint32(value))

		if ok {
			//write the enum as string
			out[fieldName] = enum
		} else {
			//write the enum as json number
			out[fieldName] = float64(value)
		}
		return

	}
}

func (j *Json) fakeVectorEnum(out map[string]interface{}, fieldName string, ot oreflection.OType) (err error) {
	var choice randutil.Choice
	var vLen int

	// generate nil or vector
	nullVector := randutil.Choice{10, int(0)}
	newVector := randutil.Choice{100, int(1)}

	choice, err = randutil.WeightedChoice([]randutil.Choice{nullVector, newVector})

	if err != nil {
		return err
	}

	item, _ := choice.Item.(int)

	if item == 0 {
		out[fieldName] = nil
		return
	}

	//generate vector len
	vLen, err = randutil.IntRange(1, 127)

	if err != nil {
		return err
	}

	r := make([]interface{}, 0, vLen)

	maxEnum := ot.Enumeration().Items().Len()

	switch ot.Items().HybridType() {

	case hybrids.Uint8:

		for i := 0; i < vLen; i++ {
			value, err := randutil.IntRange(0, maxEnum)

			if err != nil {
				return err
			}

			enumName, ok := ot.LookupEnumeration().ByUint8ToCamelCase(uint8(value))

			if ok {
				//write the enum as string
				r = append(r, enumName)
			} else {
				//write the enum as json number
				r = append(r, float64(value))
			}
		}

	case hybrids.Uint16:
		for i := 0; i < vLen; i++ {
			value, err := randutil.IntRange(0, maxEnum)

			if err != nil {
				return err
			}

			enumName, ok := ot.LookupEnumeration().ByUint16ToCamelCase(uint16(value))

			if ok {
				//write the enum as string
				r = append(r, enumName)
			} else {
				//write the enum as json number
				r = append(r, float64(value))
			}
		}

	case hybrids.Uint32:
		for i := 0; i < vLen; i++ {
			value, err := randutil.IntRange(0, maxEnum)

			if err != nil {
				return err
			}

			enumName, ok := ot.LookupEnumeration().ByUint32ToCamelCase(uint32(value))

			if ok {
				//write the enum as string
				r = append(r, enumName)
			} else {
				//write the enum as json number
				r = append(r, float64(value))
			}
		}

	}

	out[fieldName] = r
	return
}
*/