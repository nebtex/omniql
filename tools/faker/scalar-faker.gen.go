// Code generated, DO NOT EDIT.
package faker

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"strconv"
	"github.com/nebtex/omniql/commons/golang/oreflection"
	"fmt"
)

func (j *Json) fakeBoolean(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var v bool

	v, err = j.fieldGen.Boolean(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.Boolean,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorBoolean(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var shouldNil bool
	var vLen int
	var v bool

	shouldNil, err = j.fieldGen.ShouldBeNil(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorBoolean,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	if shouldNil {
		out[fieldName] = nil
		return
	}

	//generate vector len
	vLen, err = j.fieldGen.VectorLen(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorBoolean,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Boolean(path, parentType, fn)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  hybrids.Boolean,
				OmniqlType:  parentType.Id(),
				Application: parentType.Application(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, v)

	}

	out[fieldName] = r

	return
}


func (j *Json) fakeInt8(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var v int8

	v, err = j.fieldGen.Int8(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.Int8,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldName] = float64(v)

	return
}
	
func (j *Json) fakeVectorInt8(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var shouldNil bool
	var vLen int
	var v int8

	shouldNil, err = j.fieldGen.ShouldBeNil(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorInt8,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	if shouldNil {
		out[fieldName] = nil
		return
	}

	//generate vector len
	vLen, err = j.fieldGen.VectorLen(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorInt8,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Int8(path, parentType, fn)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  hybrids.Int8,
				OmniqlType:  parentType.Id(),
				Application: parentType.Application(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldName] = r

	return
}


func (j *Json) fakeUint8(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var v uint8

	v, err = j.fieldGen.Uint8(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.Uint8,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldName] = float64(v)

	return
}
	
func (j *Json) fakeVectorUint8(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var shouldNil bool
	var vLen int
	var v uint8

	shouldNil, err = j.fieldGen.ShouldBeNil(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorUint8,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	if shouldNil {
		out[fieldName] = nil
		return
	}

	//generate vector len
	vLen, err = j.fieldGen.VectorLen(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorUint8,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Uint8(path, parentType, fn)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  hybrids.Uint8,
				OmniqlType:  parentType.Id(),
				Application: parentType.Application(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldName] = r

	return
}


func (j *Json) fakeInt16(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var v int16

	v, err = j.fieldGen.Int16(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.Int16,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldName] = float64(v)

	return
}
	
func (j *Json) fakeVectorInt16(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var shouldNil bool
	var vLen int
	var v int16

	shouldNil, err = j.fieldGen.ShouldBeNil(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorInt16,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	if shouldNil {
		out[fieldName] = nil
		return
	}

	//generate vector len
	vLen, err = j.fieldGen.VectorLen(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorInt16,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Int16(path, parentType, fn)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  hybrids.Int16,
				OmniqlType:  parentType.Id(),
				Application: parentType.Application(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldName] = r

	return
}


func (j *Json) fakeUint16(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var v uint16

	v, err = j.fieldGen.Uint16(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.Uint16,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldName] = float64(v)

	return
}
	
func (j *Json) fakeVectorUint16(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var shouldNil bool
	var vLen int
	var v uint16

	shouldNil, err = j.fieldGen.ShouldBeNil(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorUint16,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	if shouldNil {
		out[fieldName] = nil
		return
	}

	//generate vector len
	vLen, err = j.fieldGen.VectorLen(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorUint16,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Uint16(path, parentType, fn)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  hybrids.Uint16,
				OmniqlType:  parentType.Id(),
				Application: parentType.Application(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldName] = r

	return
}


func (j *Json) fakeInt32(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var v int32

	v, err = j.fieldGen.Int32(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.Int32,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldName] = float64(v)

	return
}
	
func (j *Json) fakeVectorInt32(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var shouldNil bool
	var vLen int
	var v int32

	shouldNil, err = j.fieldGen.ShouldBeNil(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorInt32,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	if shouldNil {
		out[fieldName] = nil
		return
	}

	//generate vector len
	vLen, err = j.fieldGen.VectorLen(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorInt32,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Int32(path, parentType, fn)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  hybrids.Int32,
				OmniqlType:  parentType.Id(),
				Application: parentType.Application(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldName] = r

	return
}


func (j *Json) fakeUint32(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var v uint32

	v, err = j.fieldGen.Uint32(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.Uint32,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldName] = float64(v)

	return
}
	
func (j *Json) fakeVectorUint32(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var shouldNil bool
	var vLen int
	var v uint32

	shouldNil, err = j.fieldGen.ShouldBeNil(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorUint32,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	if shouldNil {
		out[fieldName] = nil
		return
	}

	//generate vector len
	vLen, err = j.fieldGen.VectorLen(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorUint32,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Uint32(path, parentType, fn)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  hybrids.Uint32,
				OmniqlType:  parentType.Id(),
				Application: parentType.Application(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldName] = r

	return
}


func (j *Json) fakeInt64(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var v int64

	v, err = j.fieldGen.Int64(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.Int64,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldName] = strconv.FormatInt(int64(v), 10)

	return
}
	
func (j *Json) fakeVectorInt64(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var shouldNil bool
	var vLen int
	var v int64

	shouldNil, err = j.fieldGen.ShouldBeNil(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorInt64,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	if shouldNil {
		out[fieldName] = nil
		return
	}

	//generate vector len
	vLen, err = j.fieldGen.VectorLen(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorInt64,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Int64(path, parentType, fn)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  hybrids.Int64,
				OmniqlType:  parentType.Id(),
				Application: parentType.Application(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, strconv.FormatInt(int64(v), 10))

	}

	out[fieldName] = r

	return
}


func (j *Json) fakeUint64(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var v uint64

	v, err = j.fieldGen.Uint64(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.Uint64,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldName] = strconv.FormatUint(uint64(v), 10)

	return
}
	
func (j *Json) fakeVectorUint64(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var shouldNil bool
	var vLen int
	var v uint64

	shouldNil, err = j.fieldGen.ShouldBeNil(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorUint64,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	if shouldNil {
		out[fieldName] = nil
		return
	}

	//generate vector len
	vLen, err = j.fieldGen.VectorLen(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorUint64,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Uint64(path, parentType, fn)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  hybrids.Uint64,
				OmniqlType:  parentType.Id(),
				Application: parentType.Application(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, strconv.FormatUint(uint64(v), 10))

	}

	out[fieldName] = r

	return
}


func (j *Json) fakeFloat32(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var v float32

	v, err = j.fieldGen.Float32(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.Float32,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldName] = float64(v)

	return
}
	
func (j *Json) fakeVectorFloat32(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var shouldNil bool
	var vLen int
	var v float32

	shouldNil, err = j.fieldGen.ShouldBeNil(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorFloat32,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	if shouldNil {
		out[fieldName] = nil
		return
	}

	//generate vector len
	vLen, err = j.fieldGen.VectorLen(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorFloat32,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Float32(path, parentType, fn)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  hybrids.Float32,
				OmniqlType:  parentType.Id(),
				Application: parentType.Application(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldName] = r

	return
}


func (j *Json) fakeFloat64(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var v float64

	v, err = j.fieldGen.Float64(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.Float64,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldName] = float64(v)

	return
}
	
func (j *Json) fakeVectorFloat64(path string, out map[string]interface{}, fieldName string, parentType oreflection.OType, fn hybrids.FieldNumber) (err error) {
	var shouldNil bool
	var vLen int
	var v float64

	shouldNil, err = j.fieldGen.ShouldBeNil(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorFloat64,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	if shouldNil {
		out[fieldName] = nil
		return
	}

	//generate vector len
	vLen, err = j.fieldGen.VectorLen(path, parentType, fn)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  hybrids.VectorFloat64,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Float64(path, parentType, fn)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  hybrids.Float64,
				OmniqlType:  parentType.Id(),
				Application: parentType.Application(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldName] = r

	return
}



func (j *Json) fakeScalar(path string, out map[string]interface{}, fieldName string, fieldType hybrids.Types, parentType oreflection.OType, fn hybrids.FieldNumber)(err error){

	switch fieldType{

	case hybrids.Boolean:
		err = j.fakeBoolean(path, out, fieldName, parentType, fn)

	case hybrids.Int8:
		err = j.fakeInt8(path, out, fieldName, parentType, fn)

	case hybrids.Uint8:
		err = j.fakeUint8(path, out, fieldName, parentType, fn)

	case hybrids.Int16:
		err = j.fakeInt16(path, out, fieldName, parentType, fn)

	case hybrids.Uint16:
		err = j.fakeUint16(path, out, fieldName, parentType, fn)

	case hybrids.Int32:
		err = j.fakeInt32(path, out, fieldName, parentType, fn)

	case hybrids.Uint32:
		err = j.fakeUint32(path, out, fieldName, parentType, fn)

	case hybrids.Int64:
		err = j.fakeInt64(path, out, fieldName, parentType, fn)

	case hybrids.Uint64:
		err = j.fakeUint64(path, out, fieldName, parentType, fn)

	case hybrids.Float32:
		err = j.fakeFloat32(path, out, fieldName, parentType, fn)

	case hybrids.Float64:
		err = j.fakeFloat64(path, out, fieldName, parentType, fn)

	default:
		err = &Error{
			Path:        path,
			HybridType:  fieldType,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    "path not recognized as  scalar",
		}
    }
	return
}


func (j *Json) fakeVectorScalar(path string, out map[string]interface{}, fieldName string, fieldType hybrids.Types, parentType oreflection.OType, fn hybrids.FieldNumber)(err error){

	switch fieldType{

	case hybrids.VectorBoolean:
		err = j.fakeVectorBoolean(path, out, fieldName, parentType, fn)

	case hybrids.VectorInt8:
		err = j.fakeVectorInt8(path, out, fieldName, parentType, fn)

	case hybrids.VectorUint8:
		err = j.fakeVectorUint8(path, out, fieldName, parentType, fn)

	case hybrids.VectorInt16:
		err = j.fakeVectorInt16(path, out, fieldName, parentType, fn)

	case hybrids.VectorUint16:
		err = j.fakeVectorUint16(path, out, fieldName, parentType, fn)

	case hybrids.VectorInt32:
		err = j.fakeVectorInt32(path, out, fieldName, parentType, fn)

	case hybrids.VectorUint32:
		err = j.fakeVectorUint32(path, out, fieldName, parentType, fn)

	case hybrids.VectorInt64:
		err = j.fakeVectorInt64(path, out, fieldName, parentType, fn)

	case hybrids.VectorUint64:
		err = j.fakeVectorUint64(path, out, fieldName, parentType, fn)

	case hybrids.VectorFloat32:
		err = j.fakeVectorFloat32(path, out, fieldName, parentType, fn)

	case hybrids.VectorFloat64:
		err = j.fakeVectorFloat64(path, out, fieldName, parentType, fn)

	default:
		err = &Error{
			Path:        path,
			HybridType:  fieldType,
			OmniqlType:  parentType.Id(),
			Application: parentType.Application(),
			ErrorMsg:    "path not recognized as vector of scalar",
		}
	}
	return
}


