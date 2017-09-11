package faker

import (
	"github.com/nebtex/omniql/commons/golang/oreflection"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
)

func (j *Json) fakeStruct(path string, out map[string]interface{}, fieldType oreflection.OType, structType oreflection.Struct) (err error) {

	var i hybrids.FieldNumber

	structObject := map[string]interface{}{}
	for i = 0; i < structType.FieldCount(); i++ {
		structFieldName, structFieldType, _ := structType.LookupFields().ByNumber(i)
		if structFieldType.Field().IsEnumeration() {
			enumType := structFieldType.Field().Type()
			err = j.fakeEnum(path+"."+structFieldName, structObject, structFieldType, enumType)
			if err != nil {
				return
			}
		} else {
			err = j.fakeScalar(path+"."+structFieldName, structObject, structFieldType)
			if err != nil {
				return
			}
		}
	}

	out[fieldType.Field().Name()] = structObject
	return
}

func (j *Json) fakeVectorStruct(path string, out map[string]interface{}, ft oreflection.OType, structType oreflection.Struct) (err error) {
	var shouldNil bool
	var vLen int
	var field hybrids.FieldNumber

	shouldNil, err = j.fieldGen.ShouldBeNil(path, ft)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: ft.Field().HybridType(),
			OmniqlType: ft.Id(),
			Package:    ft.Package(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	if shouldNil {
		out[ft.Field().Name()] = nil
		return
	}

	//generate vector len
	vLen, err = j.fieldGen.VectorLen(path, ft)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: ft.Field().HybridType(),
			OmniqlType: ft.Id(),
			Package:    ft.Package(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	r := make([]map[string]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		structObject := map[string]interface{}{}

		for field = 0; field < structType.FieldCount(); field++ {
			structFieldName, structFieldType, _ := structType.LookupFields().ByNumber(field)
			err = j.fakeScalar(path+fmt.Sprintf("[%d].%s", i, structFieldName), structObject, structFieldType)
			if err != nil {
				return
			}
		}
		r = append(r, structObject)
	}

	out[ft.Field().Name()] = r

	return
}
