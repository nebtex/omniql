// Code generated, DO NOT EDIT.
package faker

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"strconv"
	"github.com/nebtex/omniql/commons/golang/oreflection"
	"fmt"
)

func (j *Json) fakeBoolean(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var v bool

	v, err = j.fieldGen.Boolean(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldType.Field().Name()] = v

	return
}
	
func (j *Json) fakeVectorBoolean(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var shouldNil bool
	var vLen int
	var v bool

	shouldNil, err = j.fieldGen.ShouldBeNil(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
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
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Boolean(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Field().Items().HybridType(),
				OmniqlType:  fieldType.Id(),
				Package:     fieldType.Package(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, v)

	}

	out[fieldType.Field().Name()] = r

	return
}


func (j *Json) fakeInt8(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var v int8

	v, err = j.fieldGen.Int8(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldType.Field().Name()] = v

	return
}
	
func (j *Json) fakeVectorInt8(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var shouldNil bool
	var vLen int
	var v int8

	shouldNil, err = j.fieldGen.ShouldBeNil(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
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
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Int8(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Field().Items().HybridType(),
				OmniqlType:  fieldType.Id(),
				Package:     fieldType.Package(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldType.Field().Name()] = r

	return
}


func (j *Json) fakeUint8(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var v uint8

	v, err = j.fieldGen.Uint8(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldType.Field().Name()] = v

	return
}
	
func (j *Json) fakeVectorUint8(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var shouldNil bool
	var vLen int
	var v uint8

	shouldNil, err = j.fieldGen.ShouldBeNil(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
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
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Uint8(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Field().Items().HybridType(),
				OmniqlType:  fieldType.Id(),
				Package:     fieldType.Package(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldType.Field().Name()] = r

	return
}


func (j *Json) fakeInt16(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var v int16

	v, err = j.fieldGen.Int16(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldType.Field().Name()] = v

	return
}
	
func (j *Json) fakeVectorInt16(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var shouldNil bool
	var vLen int
	var v int16

	shouldNil, err = j.fieldGen.ShouldBeNil(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
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
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Int16(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Field().Items().HybridType(),
				OmniqlType:  fieldType.Id(),
				Package:     fieldType.Package(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldType.Field().Name()] = r

	return
}


func (j *Json) fakeUint16(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var v uint16

	v, err = j.fieldGen.Uint16(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldType.Field().Name()] = v

	return
}
	
func (j *Json) fakeVectorUint16(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var shouldNil bool
	var vLen int
	var v uint16

	shouldNil, err = j.fieldGen.ShouldBeNil(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
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
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Uint16(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Field().Items().HybridType(),
				OmniqlType:  fieldType.Id(),
				Package:     fieldType.Package(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldType.Field().Name()] = r

	return
}


func (j *Json) fakeInt32(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var v int32

	v, err = j.fieldGen.Int32(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldType.Field().Name()] = v

	return
}
	
func (j *Json) fakeVectorInt32(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var shouldNil bool
	var vLen int
	var v int32

	shouldNil, err = j.fieldGen.ShouldBeNil(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
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
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Int32(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Field().Items().HybridType(),
				OmniqlType:  fieldType.Id(),
				Package:     fieldType.Package(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldType.Field().Name()] = r

	return
}


func (j *Json) fakeUint32(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var v uint32

	v, err = j.fieldGen.Uint32(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldType.Field().Name()] = v

	return
}
	
func (j *Json) fakeVectorUint32(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var shouldNil bool
	var vLen int
	var v uint32

	shouldNil, err = j.fieldGen.ShouldBeNil(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
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
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Uint32(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Field().Items().HybridType(),
				OmniqlType:  fieldType.Id(),
				Package:     fieldType.Package(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldType.Field().Name()] = r

	return
}


func (j *Json) fakeInt64(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var v int64

	v, err = j.fieldGen.Int64(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldType.Field().Name()] = v

	return
}
	
func (j *Json) fakeVectorInt64(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var shouldNil bool
	var vLen int
	var v int64

	shouldNil, err = j.fieldGen.ShouldBeNil(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
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
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Int64(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Field().Items().HybridType(),
				OmniqlType:  fieldType.Id(),
				Package:     fieldType.Package(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, strconv.FormatInt(int64(v), 10))

	}

	out[fieldType.Field().Name()] = r

	return
}


func (j *Json) fakeUint64(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var v uint64

	v, err = j.fieldGen.Uint64(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldType.Field().Name()] = v

	return
}
	
func (j *Json) fakeVectorUint64(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var shouldNil bool
	var vLen int
	var v uint64

	shouldNil, err = j.fieldGen.ShouldBeNil(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
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
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Uint64(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Field().Items().HybridType(),
				OmniqlType:  fieldType.Id(),
				Package:     fieldType.Package(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, strconv.FormatUint(uint64(v), 10))

	}

	out[fieldType.Field().Name()] = r

	return
}


func (j *Json) fakeFloat32(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var v float32

	v, err = j.fieldGen.Float32(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldType.Field().Name()] = v

	return
}
	
func (j *Json) fakeVectorFloat32(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var shouldNil bool
	var vLen int
	var v float32

	shouldNil, err = j.fieldGen.ShouldBeNil(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
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
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Float32(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Field().Items().HybridType(),
				OmniqlType:  fieldType.Id(),
				Package:     fieldType.Package(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldType.Field().Name()] = r

	return
}


func (j *Json) fakeFloat64(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var v float64

	v, err = j.fieldGen.Float64(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldType.Field().Name()] = v

	return
}
	
func (j *Json) fakeVectorFloat64(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var shouldNil bool
	var vLen int
	var v float64

	shouldNil, err = j.fieldGen.ShouldBeNil(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
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
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.Float64(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Field().Items().HybridType(),
				OmniqlType:  fieldType.Id(),
				Package:     fieldType.Package(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldType.Field().Name()] = r

	return
}



func (j *Json) fakeScalar(path string, out map[string]interface{}, fieldType oreflection.OType)(err error){

	switch fieldType.Field().HybridType(){

	case hybrids.Boolean:
		err = j.fakeBoolean(path, out, fieldType)

	case hybrids.Int8:
		err = j.fakeInt8(path, out, fieldType)

	case hybrids.Uint8:
		err = j.fakeUint8(path, out, fieldType)

	case hybrids.Int16:
		err = j.fakeInt16(path, out, fieldType)

	case hybrids.Uint16:
		err = j.fakeUint16(path, out, fieldType)

	case hybrids.Int32:
		err = j.fakeInt32(path, out, fieldType)

	case hybrids.Uint32:
		err = j.fakeUint32(path, out, fieldType)

	case hybrids.Int64:
		err = j.fakeInt64(path, out, fieldType)

	case hybrids.Uint64:
		err = j.fakeUint64(path, out, fieldType)

	case hybrids.Float32:
		err = j.fakeFloat32(path, out, fieldType)

	case hybrids.Float64:
		err = j.fakeFloat64(path, out, fieldType)

	default:
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    "path not recognized as  scalar",
		}
    }
	return
}


func (j *Json) fakeVectorScalar(path string, out map[string]interface{}, fieldType oreflection.OType)(err error){

	switch fieldType.Field().HybridType{

	case hybrids.VectorBoolean:
		err = j.fakeVectorBoolean(path, out, fieldType)

	case hybrids.VectorInt8:
		err = j.fakeVectorInt8(path, out, fieldType)

	case hybrids.VectorUint8:
		err = j.fakeVectorUint8(path, out, fieldType)

	case hybrids.VectorInt16:
		err = j.fakeVectorInt16(path, out, fieldType)

	case hybrids.VectorUint16:
		err = j.fakeVectorUint16(path, out, fieldType)

	case hybrids.VectorInt32:
		err = j.fakeVectorInt32(path, out, fieldType)

	case hybrids.VectorUint32:
		err = j.fakeVectorUint32(path, out, fieldType)

	case hybrids.VectorInt64:
		err = j.fakeVectorInt64(path, out, fieldType)

	case hybrids.VectorUint64:
		err = j.fakeVectorUint64(path, out, fieldType)

	case hybrids.VectorFloat32:
		err = j.fakeVectorFloat32(path, out, fieldType)

	case hybrids.VectorFloat64:
		err = j.fakeVectorFloat64(path, out, fieldType)

	default:
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package: 	 fieldType.Package(),
			ErrorMsg:    "path not recognized as vector of scalar",
		}
	}
	return
}


