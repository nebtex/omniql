package faker

import (
	"github.com/nebtex/omniql/commons/golang/oreflection"
	"github.com/nebtex/hybrids/golang/hybrids"
	"github.com/jmcvetta/randutil"
	"fmt"
)

func (j *Json) fakeStruct(path string, out map[string]interface{}, fieldName string, ot oreflection.OType) (err error) {

	var i hybrids.FieldNumber

	structObject := map[string]interface{}{}

	for i = 0; i < ot.FieldCount(); i++ {
		structFieldName, structFieldType, _ := ot.LookupFields().ByNumber(i)
		err = j.fakeScalar(path+"."+structFieldName, structObject, structFieldName, structFieldType.HybridType())
		if err != nil {
			return
		}
	}

	out[fieldName] = structObject
	return
}


func (j *Json) fakeVectorStruct(path string, out map[string]interface{}, fieldName string, ot oreflection.OType) (err error) {
	var choice randutil.Choice
	var vLen int
	var field hybrids.FieldNumber

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

	r := make([]map[string]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		structObject := map[string]interface{}{}

		for field = 0; field < ot.FieldCount(); field++ {
			structFieldName, structFieldType, _ := ot.Items().LookupFields().ByNumber(field)
			err = j.fakeScalar(path+fmt.Sprintf("[%d].%s", int(field), structFieldName), structObject, structFieldName, structFieldType.HybridType())
			if err != nil {
				return
			}
		}

		r[i] = structObject
	}

	out[fieldName] = r

	return
}
