package kkk

import (
	"fmt"
	"github.com/nebtex/delta/go"
)

func getFieldNumber(buf []byte, offset delta.OffsetUint32) delta.FieldNumber {
	return delta.FieldNumber(buf[offset])
}

func getValueOffset(buf []byte, offset delta.OffsetUint32) delta.OffsetUint16 {
	return delta.GetOffsetUint16(buf, offset+1)
}

//LookupField search a field in the bTable using binary search, if field is not found, the ok variable will be false
//if valueOffset==0 and ok==true, this should be interpreted as a deletion message
func LookupField(buf []byte, offset delta.OffsetUint32, number delta.FieldNumber) (valueOffset delta.OffsetUint16, ok bool) {
	var pOffset delta.OffsetUint32
	var pivot uint16
	var max uint16
	var min uint16
	var pn delta.FieldNumber

	//get the field count

	n := FieldCount(buf, offset)

	//fields are sorted in ascending way, so fist compare with field locate in the center of the bTable
	pivot = uint16(n) / 2

	max = uint16(n) - 1
	min = 0

	//check if field is in range, leave other case
	if getFieldNumber(buf, offset+1) > number || getFieldNumber(buf, offset+1+delta.OffsetUint32(max*3)) < number {
		return
	}

	for {

		pOffset = delta.OffsetUint32(pivot*3) + offset + 1
		pn = getFieldNumber(buf, pOffset)
		if pn == number {
			//eureka we found the field
			valueOffset = getValueOffset(buf, pOffset)
			ok = true
			return
		}

		if pivot <= min || pivot >= max {
			//field is not defined in this vtable, :(
			break
		}

		if pn > number {
			max = pivot
			pivot -= (pivot-min)/2 | 1
		}

		if pn < number {
			//calculate next pivot
			min = pivot
			pivot += (max-pivot)/2 | 1
		}
	}
	return
}

//Size calculated the total size on bytes of the btable, this function should never return 0
//if a btable exist it should always have almost a field
func Size(buf []byte, offset delta.OffsetUint32) (size int) {
	size = FieldCount(buf, offset)*3 + 1
	return
}

//FieldCount, the first byte of the BTable contains the number of fields defined, each field takes 3 bytes
//
//the first byte contains the field number, the second byte contain the offset where the value of the field is store,
//if the second //byte is 0 this should be interpreted as a delete message over the field
func FieldCount(buf []byte, offset delta.OffsetUint32) (fc int) {
	fc = int(delta.GetUint8(buf, offset))
	return
}

//GetFieldAt return a field located at any position of the btable array, use this for iterate over the fields
//after i > n, fn and valueOffset will always return 0
func GetFieldAt(buf []byte, offset delta.OffsetUint32, i int) (fn delta.FieldNumber, valueOffset delta.OffsetUint16) {
	var iOffset delta.OffsetUint32
	var n int
	n = FieldCount(buf, offset)

	if i > n {
		panic("btable: index out of range")
	}

	iOffset = delta.OffsetUint32(i*3) + offset + 1
	fn = getFieldNumber(buf, iOffset)
	valueOffset = getValueOffset(buf, iOffset)
	return

}

//Validate the bTable
//
//first: check the size values is right, does not overflow the buffer
//
//second: check that the fields are sorted in ascending order
func Validate(buf []byte, offset delta.OffsetUint32) (err error) {
	var i, size int
	var min, newMin delta.FieldNumber
	size = Size(buf, offset)
	if (size + int(offset)) > len(buf) {
		err = fmt.Errorf("btable: overflows the buffer")
		return
	}
	fnc := FieldCount(buf, offset)
	for i = 0; i < fnc; i++ {
		newMin, _ = GetFieldAt(buf, offset, i)
		if min == 0 && newMin == 0 {
			continue
		}

		if min != 0 {
			if newMin < min {
				err = fmt.Errorf("btable: field numbers are not sorted in acesding order")
				return
			}
		}
		min = newMin
	}
	return
}