package faker

import "github.com/nebtex/omnibuff/commons/golang/tools/oreflection"


import (
"github.com/nebtex/omnibuff/tools/golang/tools/oreflection"
"github.com/nebtex/hybrids/golang/hybrids"
)

func (e *Encoder) encodeEnum(out map[string]interface{}, fieldName string, ot oreflection.OType, fn hybrids.FieldNumber, tr hybrids.ScalarReader) {
	var enum string

	switch ot.HybridType() {

	case hybrids.Uint8:
		v, ok := tr.Uint8(fn)
		if !ok {
			// value is not defined
			return
		}
		enum, ok = ot.LookupEnumeration().ByUint8ToCamelCase(v)
		if ok {
			//write the enum as string
			out[fieldName] = enum
		} else {
			//write the enum as json number
			out[fieldName] = float64(v)
		}
		return

	case hybrids.Uint16:
		v, ok := tr.Uint16(fn)
		if !ok {
			// value is not defined
			return
		}
		enum, ok = ot.LookupEnumeration().ByUint16ToCamelCase(v)
		if ok {
			//write the enum as string
			out[fieldName] = enum
		} else {
			//write the enum as json number
			out[fieldName] = float64(v)
		}
		return

	case hybrids.Uint32:
		v, ok := tr.Uint32(fn)
		if !ok {
			// value is not defined
			return
		}
		enum, ok = ot.LookupEnumeration().ByUint32ToCamelCase(v)
		if ok {
			//write the enum as string
			out[fieldName] = enum
		} else {
			//write the enum as json number
			out[fieldName] = float64(v)
		}
		return

	}
}

func (e *Encoder) encodeVectorEnum(out map[string]interface{}, fieldName string, ot oreflection.OType, fn hybrids.FieldNumber, vrb hybrids.VectorScalarReaderAccessor) {
	var enum string
	var r []interface{}

	switch ot.Items().HybridType() {

	case hybrids.Uint8:

		vu, ok := vrb.VectorUint8(fn)
		if !ok {
			//vector undefined
			return
		}
		if vu.Len() == 0 {
			//vector is null
			out[fieldName] = nil
			return
		}
		r = make([]interface{}, 0, vu.Len())

		for i := 0; i < vu.Len(); i++ {
			enum, ok = ot.LookupEnumeration().ByUint8ToCamelCase(vu.Get(i))
			if ok {
				//write the enum as string
				r = append(r, enum)
			} else {
				//write the enum as json number
				r = append(r, float64(vu.Get(i)))
			}
		}

		out[fieldName] = r
		return

	case hybrids.Uint16:
		vu, ok := vrb.VectorUint16(fn)
		if !ok {
			//vector undefined
			return
		}
		if vu.Len() == 0 {
			//vector is null
			out[fieldName] = nil
			return
		}
		r = make([]interface{}, 0, vu.Len())

		for i := 0; i < vu.Len(); i++ {
			enum, ok = ot.LookupEnumeration().ByUint16ToCamelCase(vu.Get(i))
			if ok {
				//write the enum as string
				r = append(r, enum)
			} else {
				//write the enum as json number
				r = append(r, float64(vu.Get(i)))
			}
		}

		out[fieldName] = r
		return

	case hybrids.Uint32:
		vu, ok := vrb.VectorUint32(fn)
		if !ok {
			//vector undefined
			return
		}
		if vu.Len() == 0 {
			//vector is null
			out[fieldName] = nil
			return
		}
		r = make([]interface{}, 0, vu.Len())

		for i := 0; i < vu.Len(); i++ {
			enum, ok = ot.LookupEnumeration().ByUint32ToCamelCase(vu.Get(i))
			if ok {
				//write the enum as string
				r = append(r, enum)
			} else {
				//write the enum as json number
				r = append(r, float64(vu.Get(i)))
			}
		}

		out[fieldName] = r
		return
	}
}
