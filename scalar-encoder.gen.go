// Code generated, DO NOT EDIT.
package faker

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"strconv"
	"github.com/jmcvetta/randutil"
)

func (j *Json) getBoolean() (v bool, err error) {
	var r int
	r, err = randutil.IntRange(hybrids.MinBoolean, hybrids.MaxBoolean)
	if err != nil {
		return
	}
	v = iToBool(r)
	return
}

func (j *Json) fakeBoolean(path string, out map[string]interface{}, fieldName string) (err error) {
	var v bool
	v, err = j.getBoolean()
	if err != nil {
		return
	}
	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorBoolean(out map[string]interface{}, fieldName string) (err error) {
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

	for i := 0; i < vLen; i++ {
		r[i] = j.getBoolean()
	}

	out[fieldName] = r

	return
}



func (j *Json) getInt8() (v int8, err error) {
	var r int
	r, err = randutil.IntRange(hybrids.MinInt8, hybrids.MaxInt8)
	if err != nil {
		return
	}
	v = iToBool(r)
	return
}

func (j *Json) fakeInt8(path string, out map[string]interface{}, fieldName string) (err error) {
	var v int8
	v, err = j.getInt8()
	if err != nil {
		return
	}
	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorInt8(out map[string]interface{}, fieldName string) (err error) {
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

	for i := 0; i < vLen; i++ {
		r[i] = j.getInt8()
	}

	out[fieldName] = r

	return
}



func (j *Json) getUint8() (v uint8, err error) {
	var r int
	r, err = randutil.IntRange(hybrids.MinUint8, hybrids.MaxUint8)
	if err != nil {
		return
	}
	v = iToBool(r)
	return
}

func (j *Json) fakeUint8(path string, out map[string]interface{}, fieldName string) (err error) {
	var v uint8
	v, err = j.getUint8()
	if err != nil {
		return
	}
	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorUint8(out map[string]interface{}, fieldName string) (err error) {
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

	for i := 0; i < vLen; i++ {
		r[i] = j.getUint8()
	}

	out[fieldName] = r

	return
}



func (j *Json) getInt16() (v int16, err error) {
	var r int
	r, err = randutil.IntRange(hybrids.MinInt16, hybrids.MaxInt16)
	if err != nil {
		return
	}
	v = iToBool(r)
	return
}

func (j *Json) fakeInt16(path string, out map[string]interface{}, fieldName string) (err error) {
	var v int16
	v, err = j.getInt16()
	if err != nil {
		return
	}
	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorInt16(out map[string]interface{}, fieldName string) (err error) {
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

	for i := 0; i < vLen; i++ {
		r[i] = j.getInt16()
	}

	out[fieldName] = r

	return
}



func (j *Json) getUint16() (v uint16, err error) {
	var r int
	r, err = randutil.IntRange(hybrids.MinUint16, hybrids.MaxUint16)
	if err != nil {
		return
	}
	v = iToBool(r)
	return
}

func (j *Json) fakeUint16(path string, out map[string]interface{}, fieldName string) (err error) {
	var v uint16
	v, err = j.getUint16()
	if err != nil {
		return
	}
	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorUint16(out map[string]interface{}, fieldName string) (err error) {
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

	for i := 0; i < vLen; i++ {
		r[i] = j.getUint16()
	}

	out[fieldName] = r

	return
}



func (j *Json) getInt32() (v int32, err error) {
	var r int
	r, err = randutil.IntRange(hybrids.MinInt32, hybrids.MaxInt32)
	if err != nil {
		return
	}
	v = iToBool(r)
	return
}

func (j *Json) fakeInt32(path string, out map[string]interface{}, fieldName string) (err error) {
	var v int32
	v, err = j.getInt32()
	if err != nil {
		return
	}
	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorInt32(out map[string]interface{}, fieldName string) (err error) {
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

	for i := 0; i < vLen; i++ {
		r[i] = j.getInt32()
	}

	out[fieldName] = r

	return
}



func (j *Json) getUint32() (v uint32, err error) {
	var r int
	r, err = randutil.IntRange(hybrids.MinUint32, hybrids.MaxUint32)
	if err != nil {
		return
	}
	v = iToBool(r)
	return
}

func (j *Json) fakeUint32(path string, out map[string]interface{}, fieldName string) (err error) {
	var v uint32
	v, err = j.getUint32()
	if err != nil {
		return
	}
	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorUint32(out map[string]interface{}, fieldName string) (err error) {
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

	for i := 0; i < vLen; i++ {
		r[i] = j.getUint32()
	}

	out[fieldName] = r

	return
}



func (j *Json) getInt64() (v int64, err error) {
	var r int
	r, err = randutil.IntRange(hybrids.MinInt64, hybrids.MaxInt64)
	if err != nil {
		return
	}
	v = iToBool(r)
	return
}

func (j *Json) fakeInt64(path string, out map[string]interface{}, fieldName string) (err error) {
	var v int64
	v, err = j.getInt64()
	if err != nil {
		return
	}
	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorInt64(out map[string]interface{}, fieldName string) (err error) {
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

	for i := 0; i < vLen; i++ {
		r[i] = j.getInt64()
	}

	out[fieldName] = r

	return
}



func (j *Json) getUint64() (v uint64, err error) {
	var r int
	r, err = randutil.IntRange(hybrids.MinUint64, hybrids.MaxUint64)
	if err != nil {
		return
	}
	v = iToBool(r)
	return
}

func (j *Json) fakeUint64(path string, out map[string]interface{}, fieldName string) (err error) {
	var v uint64
	v, err = j.getUint64()
	if err != nil {
		return
	}
	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorUint64(out map[string]interface{}, fieldName string) (err error) {
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

	for i := 0; i < vLen; i++ {
		r[i] = j.getUint64()
	}

	out[fieldName] = r

	return
}



func (j *Json) getFloat32() (v float32, err error) {
	var r int
	r, err = randutil.IntRange(hybrids.MinFloat32, hybrids.MaxFloat32)
	if err != nil {
		return
	}
	v = iToBool(r)
	return
}

func (j *Json) fakeFloat32(path string, out map[string]interface{}, fieldName string) (err error) {
	var v float32
	v, err = j.getFloat32()
	if err != nil {
		return
	}
	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorFloat32(out map[string]interface{}, fieldName string) (err error) {
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

	for i := 0; i < vLen; i++ {
		r[i] = j.getFloat32()
	}

	out[fieldName] = r

	return
}



func (j *Json) getFloat64() (v float64, err error) {
	var r int
	r, err = randutil.IntRange(hybrids.MinFloat64, hybrids.MaxFloat64)
	if err != nil {
		return
	}
	v = iToBool(r)
	return
}

func (j *Json) fakeFloat64(path string, out map[string]interface{}, fieldName string) (err error) {
	var v float64
	v, err = j.getFloat64()
	if err != nil {
		return
	}
	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorFloat64(out map[string]interface{}, fieldName string) (err error) {
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

	for i := 0; i < vLen; i++ {
		r[i] = j.getFloat64()
	}

	out[fieldName] = r

	return
}




func (j *Json) fakeScalar(path string, out map[string]interface{}, fieldName string, fieldType hybrids.Types){
	switch fieldType{

	case hybrids.Boolean:
		j.fakeBoolean(path, out, fieldName)

	case hybrids.Int8:
		j.fakeInt8(path, out, fieldName)

	case hybrids.Uint8:
		j.fakeUint8(path, out, fieldName)

	case hybrids.Int16:
		j.fakeInt16(path, out, fieldName)

	case hybrids.Uint16:
		j.fakeUint16(path, out, fieldName)

	case hybrids.Int32:
		j.fakeInt32(path, out, fieldName)

	case hybrids.Uint32:
		j.fakeUint32(path, out, fieldName)

	case hybrids.Int64:
		j.fakeInt64(path, out, fieldName)

	case hybrids.Uint64:
		j.fakeUint64(path, out, fieldName)

	case hybrids.Float32:
		j.fakeFloat32(path, out, fieldName)

	case hybrids.Float64:
		j.fakeFloat64(path, out, fieldName)

    }
	return
}
