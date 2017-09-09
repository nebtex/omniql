// Code generated, DO NOT EDIT.
package faker

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"strconv"
	"github.com/jmcvetta/randutil"
)

func (j *Json) getBoolean() (r bool, err error) {
	var v int

	v, err = randutil.IntRange(hybrids.MinBoolean, hybrids.MaxBoolean)

	if err != nil {
		return
	}

	r = iToBool(v)

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
	
func (j *Json) fakeVectorBoolean(path string, out map[string]interface{}, fieldName string) (err error) {
	var choice randutil.Choice
	var vLen int
	var v bool

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
		v, err =  j.getBoolean()
		if err != nil {
	        return err
	    }
		r[i] = v
	}

	out[fieldName] = r

	return
}


func (j *Json) getInt8() (r float64, err error) {
	var v int

	v, err = randutil.IntRange(hybrids.MinInt8, hybrids.MaxInt8)

	if err != nil {
		return
	}

	r = float64(v)

	return
}

func (j *Json) fakeInt8(path string, out map[string]interface{}, fieldName string) (err error) {
	var v float64

	v, err = j.getInt8()

	if err != nil {
		return
	}

	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorInt8(path string, out map[string]interface{}, fieldName string) (err error) {
	var choice randutil.Choice
	var vLen int
	var v float64

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
		v, err =  j.getInt8()
		if err != nil {
	        return err
	    }
		r[i] = v
	}

	out[fieldName] = r

	return
}


func (j *Json) getUint8() (r float64, err error) {
	var v int

	v, err = randutil.IntRange(hybrids.MinUint8, hybrids.MaxUint8)

	if err != nil {
		return
	}

	r = float64(v)

	return
}

func (j *Json) fakeUint8(path string, out map[string]interface{}, fieldName string) (err error) {
	var v float64

	v, err = j.getUint8()

	if err != nil {
		return
	}

	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorUint8(path string, out map[string]interface{}, fieldName string) (err error) {
	var choice randutil.Choice
	var vLen int
	var v float64

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
		v, err =  j.getUint8()
		if err != nil {
	        return err
	    }
		r[i] = v
	}

	out[fieldName] = r

	return
}


func (j *Json) getInt16() (r float64, err error) {
	var v int

	v, err = randutil.IntRange(hybrids.MinInt16, hybrids.MaxInt16)

	if err != nil {
		return
	}

	r = float64(v)

	return
}

func (j *Json) fakeInt16(path string, out map[string]interface{}, fieldName string) (err error) {
	var v float64

	v, err = j.getInt16()

	if err != nil {
		return
	}

	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorInt16(path string, out map[string]interface{}, fieldName string) (err error) {
	var choice randutil.Choice
	var vLen int
	var v float64

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
		v, err =  j.getInt16()
		if err != nil {
	        return err
	    }
		r[i] = v
	}

	out[fieldName] = r

	return
}


func (j *Json) getUint16() (r float64, err error) {
	var v int

	v, err = randutil.IntRange(hybrids.MinUint16, hybrids.MaxUint16)

	if err != nil {
		return
	}

	r = float64(v)

	return
}

func (j *Json) fakeUint16(path string, out map[string]interface{}, fieldName string) (err error) {
	var v float64

	v, err = j.getUint16()

	if err != nil {
		return
	}

	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorUint16(path string, out map[string]interface{}, fieldName string) (err error) {
	var choice randutil.Choice
	var vLen int
	var v float64

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
		v, err =  j.getUint16()
		if err != nil {
	        return err
	    }
		r[i] = v
	}

	out[fieldName] = r

	return
}


func (j *Json) getInt32() (r float64, err error) {
	var v int

	v, err = randutil.IntRange(hybrids.MinInt32, hybrids.MaxInt32)

	if err != nil {
		return
	}

	r = float64(v)

	return
}

func (j *Json) fakeInt32(path string, out map[string]interface{}, fieldName string) (err error) {
	var v float64

	v, err = j.getInt32()

	if err != nil {
		return
	}

	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorInt32(path string, out map[string]interface{}, fieldName string) (err error) {
	var choice randutil.Choice
	var vLen int
	var v float64

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
		v, err =  j.getInt32()
		if err != nil {
	        return err
	    }
		r[i] = v
	}

	out[fieldName] = r

	return
}


func (j *Json) getUint32() (r float64, err error) {
	var v int

	v, err = randutil.IntRange(hybrids.MinUint32, hybrids.MaxUint32)

	if err != nil {
		return
	}

	r = float64(v)

	return
}

func (j *Json) fakeUint32(path string, out map[string]interface{}, fieldName string) (err error) {
	var v float64

	v, err = j.getUint32()

	if err != nil {
		return
	}

	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorUint32(path string, out map[string]interface{}, fieldName string) (err error) {
	var choice randutil.Choice
	var vLen int
	var v float64

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
		v, err =  j.getUint32()
		if err != nil {
	        return err
	    }
		r[i] = v
	}

	out[fieldName] = r

	return
}


func (j *Json) getInt64() (r string, err error) {
	var v int

	v, err = randutil.IntRange(-200000000, -200000000)

	if err != nil {
		return
	}

	r = strconv.FormatInt(int64(v), 10)

	return
}

func (j *Json) fakeInt64(path string, out map[string]interface{}, fieldName string) (err error) {
	var v string

	v, err = j.getInt64()

	if err != nil {
		return
	}

	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorInt64(path string, out map[string]interface{}, fieldName string) (err error) {
	var choice randutil.Choice
	var vLen int
	var v string

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
		v, err =  j.getInt64()
		if err != nil {
	        return err
	    }
		r[i] = v
	}

	out[fieldName] = r

	return
}


func (j *Json) getUint64() (r string, err error) {
	var v int

	v, err = randutil.IntRange(0, 0)

	if err != nil {
		return
	}

	r = strconv.FormatUint(uint64(v), 10)

	return
}

func (j *Json) fakeUint64(path string, out map[string]interface{}, fieldName string) (err error) {
	var v string

	v, err = j.getUint64()

	if err != nil {
		return
	}

	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorUint64(path string, out map[string]interface{}, fieldName string) (err error) {
	var choice randutil.Choice
	var vLen int
	var v string

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
		v, err =  j.getUint64()
		if err != nil {
	        return err
	    }
		r[i] = v
	}

	out[fieldName] = r

	return
}


func (j *Json) getFloat32() (r float64, err error) {
	var v int

	v, err = randutil.IntRange(-158, -158)

	if err != nil {
		return
	}

	r = float64(v)

	return
}

func (j *Json) fakeFloat32(path string, out map[string]interface{}, fieldName string) (err error) {
	var v float64

	v, err = j.getFloat32()

	if err != nil {
		return
	}

	out[fieldName] = v

	return
}
	
func (j *Json) fakeVectorFloat32(path string, out map[string]interface{}, fieldName string) (err error) {
	var choice randutil.Choice
	var vLen int
	var v float64

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
		v, err =  j.getFloat32()
		if err != nil {
	        return err
	    }
		r[i] = v
	}

	out[fieldName] = r

	return
}


func (j *Json) getFloat64() (r float64, err error) {
	var v int

	v, err = randutil.IntRange(-3500, -3500)

	if err != nil {
		return
	}

	r = float64(v)

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
	
func (j *Json) fakeVectorFloat64(path string, out map[string]interface{}, fieldName string) (err error) {
	var choice randutil.Choice
	var vLen int
	var v float64

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
		v, err =  j.getFloat64()
		if err != nil {
	        return err
	    }
		r[i] = v
	}

	out[fieldName] = r

	return
}



func (j *Json) fakeScalar(path string, out map[string]interface{}, fieldName string, fieldType hybrids.Types)(err error){
	switch fieldType{

	case hybrids.Boolean:
		err = j.fakeBoolean(path, out, fieldName)
	

	case hybrids.Int8:
		err = j.fakeInt8(path, out, fieldName)
	

	case hybrids.Uint8:
		err = j.fakeUint8(path, out, fieldName)
	

	case hybrids.Int16:
		err = j.fakeInt16(path, out, fieldName)
	

	case hybrids.Uint16:
		err = j.fakeUint16(path, out, fieldName)
	

	case hybrids.Int32:
		err = j.fakeInt32(path, out, fieldName)
	

	case hybrids.Uint32:
		err = j.fakeUint32(path, out, fieldName)
	

	case hybrids.Int64:
		err = j.fakeInt64(path, out, fieldName)
	

	case hybrids.Uint64:
		err = j.fakeUint64(path, out, fieldName)
	

	case hybrids.Float32:
		err = j.fakeFloat32(path, out, fieldName)
	

	case hybrids.Float64:
		err = j.fakeFloat64(path, out, fieldName)
	

    }
	return
}


func (j *Json) fakeVectorScalar(path string, out map[string]interface{}, fieldName string, fieldType hybrids.Types)(err error){
	switch fieldType{

	case hybrids.VectorBoolean:
		err = j.fakeVectorBoolean(path, out, fieldName)


	case hybrids.VectorInt8:
		err = j.fakeVectorInt8(path, out, fieldName)


	case hybrids.VectorUint8:
		err = j.fakeVectorUint8(path, out, fieldName)


	case hybrids.VectorInt16:
		err = j.fakeVectorInt16(path, out, fieldName)


	case hybrids.VectorUint16:
		err = j.fakeVectorUint16(path, out, fieldName)


	case hybrids.VectorInt32:
		err = j.fakeVectorInt32(path, out, fieldName)


	case hybrids.VectorUint32:
		err = j.fakeVectorUint32(path, out, fieldName)


	case hybrids.VectorInt64:
		err = j.fakeVectorInt64(path, out, fieldName)


	case hybrids.VectorUint64:
		err = j.fakeVectorUint64(path, out, fieldName)


	case hybrids.VectorFloat32:
		err = j.fakeVectorFloat32(path, out, fieldName)


	case hybrids.VectorFloat64:
		err = j.fakeVectorFloat64(path, out, fieldName)


	}
	return
}


