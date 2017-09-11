package faker

import (
	"github.com/nebtex/omniql/commons/golang/oreflection"
	"github.com/nebtex/hybrids/golang/hybrids"
)

func (j *Json) fakeTable(path string, out map[string]interface{}, fieldType oreflection.OType, tableType oreflection.Table) (err error) {
	var shouldNil bool
	var i hybrids.FieldNumber

	shouldNil, err = j.fieldGen.ShouldBeNil(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.Field().HybridType(),
			OmniqlType: fieldType.Id(),
			Package:    fieldType.Package(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	if shouldNil {
		out[fieldType.Field().Name()] = nil
		return
	}

	tableObject := map[string]interface{}{}

	for i = 0; i < tableType.FieldCount(); i++ {
		shouldGenerateField, err := j.fieldGen.ShouldGenerateField(path, fieldType, i)
		if err != nil {
			err = &Error{
				Path:       path,
				HybridType: fieldType.Field().HybridType(),
				OmniqlType: fieldType.Id(),
				Package:    fieldType.Package(),
				ErrorMsg:   err.Error(),
			}
			return
		}
		if !shouldGenerateField {
			continue
		}

		childFieldName, childFieldType, _ := tableType.LookupFields().ByNumber(i)

		if childFieldType.Field().IsEnumeration() {
			enumType := childFieldType.Field().Type().Enumeration()
			err = j.fakeEnum(path+"."+childFieldName, tableObject, childFieldType, enumType)
			if err != nil {
				return
			}
			continue
		}
		if childFieldType.Field().HybridType().IsScalar() {

			err = j.fakeScalar(path+"."+childFieldName, tableObject, childFieldType)
			if err != nil {
				return
			}
			continue
		}
		if childFieldType.Field().HybridType().IsVectorScalar() {

			err = j.fakeVectorScalar(path+"."+childFieldName, tableObject, childFieldType)
			if err != nil {
				return
			}
			continue
		}

		switch childFieldType.Field().HybridType() {
		case hybrids.Struct:
			structType := childFieldType.Field().Type().Struct()
			err = j.fakeStruct(path+"."+childFieldName, tableObject, childFieldType, structType)
		case hybrids.VectorStruct:
			structType := childFieldType.Field().Items().Type().Struct()
			err = j.fakeStruct(path+"."+childFieldName, tableObject, childFieldType, structType)
		case hybrids.String:
			err = j.fakeString(path+"."+childFieldName, tableObject, childFieldType)
		case hybrids.Byte:
			err = j.fakeByte(path+"."+childFieldName, tableObject, childFieldType)
		case hybrids.ResourceID:
			resource := childFieldType.Field().Type().Resource()
			err = j.fakeTheResourceID(path+"."+childFieldName, tableObject, childFieldType, resource)
		case hybrids.VectorString:
			err = j.fakeVectorString(path+"."+childFieldName, tableObject, childFieldType)
		case hybrids.VectorByte:
			err = j.fakeVectorByte(path+"."+childFieldName, tableObject, childFieldType)
		case hybrids.VectorResourceID:
			resource := childFieldType.Field().Items().Type().Resource()
			err = j.fakeVectorResourceID(path+"."+childFieldName, tableObject, childFieldType, resource)
		case hybrids.Table:
			table := childFieldType.Field().Type().Table()
			err = j.fakeTable(path+"."+childFieldName, tableObject, childFieldType, table)
		case hybrids.Union:
			union := childFieldType.Field().Type().Union()
			err = j.fakeUnion(path+"."+childFieldName, tableObject, childFieldType, union)
		case hybrids.VectorUnion:
			union := childFieldType.Field().Items().Type().Union()
			err = j.fakeVectorUnion(path+"."+childFieldName, tableObject, childFieldType, union)
		case hybrids.VectorTable:
			table := childFieldType.Field().Items().Type().Table()
			err = j.fakeVectorTable(path+"."+childFieldName, tableObject, childFieldType, table)
		}
		if err != nil {
			return
		}
	}

	out[fieldType.Field().Name()] = tableObject
	return
}
