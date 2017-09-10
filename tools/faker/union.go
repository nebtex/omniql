package faker

import (
	"github.com/nebtex/omniql/commons/golang/oreflection"
	"github.com/nebtex/hybrids/golang/hybrids"
	"go.uber.org/zap"
	"github.com/jmcvetta/randutil"
)

func (j *Json) fakeUnion(path string, out map[string]interface{}, fieldName string, ot oreflection.OType) (err error) {
	var unionName string
	var tableType oreflection.OType
	var choice randutil.Choice

	// generate nil or vector
	nullUnion := randutil.Choice{10, int(0)}
	newUnion := randutil.Choice{100, int(1)}

	choice, err = randutil.WeightedChoice([]randutil.Choice{nullUnion, newUnion})

	if err != nil {
		return err
	}

	item, _ := choice.Item.(int)

	if item == 0 {
		out[fieldName] = nil
		return
	}

	maxUnion := ot.Union().Items().Len()

	value, err := randutil.IntRange(0, maxUnion)

	if err != nil {
		return err
	}

	unionName, tableType, _ = ot.LookupTableOnUnion().ByPositionToCamelCase(hybrids.UnionKind(value))

	union := map[string]interface{}{}

	err = j.fakeTable(path+"."+unionName, union, unionName, tableType)
	if err != nil {
		return
	}

	out[fieldName] = union

	return

}

func (j *Json) fakeVectorUnion(path string, out map[string]interface{}, fieldName string, ot oreflection.OType) (err error) {
	var itemOt oreflection.OType
	var ok bool
	var enumName string

	vu, ok := tr.VectorUnion(fn)

	if !ok {
		//undefined
		return
	}

	if vu == nil {
		out[fieldName] = nil
		return
	}

	r := make([]map[string]interface{}, 0, vu.Len())

	for i := 0; i < vu.Len(); i++ {
		kind, table, err := vu.Get(i)

		if err != nil {

			err = &EncoderError{
				Path:        path,
				Application: e.application,
				OmniqlType:  ot.Id(),
				HybridType:  hybrids.Union,
				ErrorMsg:    err.Error(),
			}
			return
		}

		if table == nil {
			r = append(r, nil)
		}

		enumName, itemOt, ok = ot.LookupTableOnUnion().ByPositionToCamelCase(kind)

		if !ok {
			//schema not found, leave without error
			//probably is an older version and don't have the reflection of this table
			e.zap.Warn("Failed to parse a table, its schema was not found, probably is a good time to update this application",
				zap.Uint16("Table position in the union", uint16(kind)),
				zap.String("Union type", ot.Id()),
				zap.String("Path", path),
				zap.String("Application", e.application))
			r = append(r, nil)
			continue
		}
		fieldValue := map[string]interface{}{}

		err = e.encodeTable(path+"."+fieldName, fieldValue, enumName, itemOt, table)

		if err != nil {
			err = &EncoderError{
				Path:        path,
				Application: e.application,
				OmniqlType:  ot.Id(),
				HybridType:  hybrids.Union,
				ErrorMsg:    err.Error(),
			}
			return
		}
		r = append(r, fieldValue)

	}
	return
}
