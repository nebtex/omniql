package faker

import "github.com/nebtex/omniql/commons/golang/oreflection"
import (
	b64 "encoding/base64"
	"fmt"
)

func (j *Json) fakeString(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var v string

	v, err = j.fieldGen.String(path, fieldType)

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

	out[fieldType.Field().Name()] = v

	return
}

func (j *Json) fakeByte(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var v []byte

	v, err = j.fieldGen.Byte(path, fieldType)

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

	out[fieldType.Field().Name()] = b64.StdEncoding.EncodeToString(v)

	return
}

func (j *Json) fakeTheResourceID(path string, out map[string]interface{}, fieldType oreflection.OType, resource oreflection.Resource) (err error) {
	var v []byte

	v, err = j.fieldGen.ResourceID(path, fieldType, resource.ResourceIDType())

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

	out[fieldType.Field().Name()] = b64.StdEncoding.EncodeToString(v)

	return
}

func (j *Json) fakeVectorString(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var shouldNil bool
	var vLen int
	var v string

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

	//generate vector len
	vLen, err = j.fieldGen.VectorLen(path, fieldType)

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

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err = j.fieldGen.String(path, fieldType)
		if err != nil {
			err = &Error{
				Path:       path + fmt.Sprintf("[%d]", i),
				HybridType: fieldType.Field().Items().HybridType(),
				OmniqlType: fieldType.Id(),
				Package:    fieldType.Package(),
				ErrorMsg:   err.Error(),
			}
			return
		}
		r = append(r, v)

	}

	out[fieldType.Field().Name()] = r

	return
}


func (j *Json) fakeVectorByte(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var shouldNil bool
	var vLen int
	var v []byte

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

	//generate vector len
	vLen, err = j.fieldGen.VectorLen(path, fieldType)

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

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err = j.fieldGen.Byte(path, fieldType)
		if err != nil {
			err = &Error{
				Path:       path + fmt.Sprintf("[%d]", i),
				HybridType: fieldType.Field().Items().HybridType(),
				OmniqlType: fieldType.Id(),
				Package:    fieldType.Package(),
				ErrorMsg:   err.Error(),
			}
			return
		}
		r = append(r, v)

	}

	out[fieldType.Field().Name()] = b64.StdEncoding.EncodeToString(v)

	return
}


func (j *Json) fakeVectorResourceID(path string, out map[string]interface{}, fieldType oreflection.OType, resource oreflection.Resource) (err error) {
	var shouldNil bool
	var vLen int
	var v []byte

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

	//generate vector len
	vLen, err = j.fieldGen.VectorLen(path, fieldType)

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

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err = j.fieldGen.ResourceID(path, fieldType, resource.ResourceIDType())
		if err != nil {
			err = &Error{
				Path:       path + fmt.Sprintf("[%d]", i),
				HybridType: fieldType.Field().Items().HybridType(),
				OmniqlType: fieldType.Id(),
				Package:    fieldType.Package(),
				ErrorMsg:   err.Error(),
			}
			return
		}
		r = append(r, v)

	}

	out[fieldType.Field().Name()] = b64.StdEncoding.EncodeToString(v)

	return
}
