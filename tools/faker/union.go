package faker

import (
	"github.com/nebtex/omniql/commons/golang/oreflection"
	"github.com/nebtex/hybrids/golang/hybrids"
	"go.uber.org/zap"
)

func (j *Json) fakeUnion(path string, out map[string]interface{}, fieldName string, ot oreflection.OType) (err error) {
	ot.

}

func (e *Encoder) encodeVectorUnion(path string, out map[string]interface{}, fieldName string, ot oreflection.OType, fn hybrids.FieldNumber, tr hybrids.TableReader) (err error) {
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
