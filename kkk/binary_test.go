package kkk

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
	"github.com/nebtex/delta/go"
	"github.com/google/flatbuffers/go"
)

func Test_IterateOverBTable(t *testing.T) {
	table := []struct {
		bTable                []byte
		fieldResult           []delta.FieldNumber
		offsetResult          []delta.OffsetUint16
		sizeShouldError       bool
		getFieldAtShouldError bool
		name                  string
	}{
		{
			bTable:       []byte{7, 0, 25, 0, 1, 30, 0, 5, 50, 0, 20, 60, 0, 22, 70, 0, 55, 90, 0, 60, 40, 0},
			name:         "bTable without error, all should be ok",
			fieldResult:  []delta.FieldNumber{0, 1, 5, 20, 22, 55, 60},
			offsetResult: []delta.OffsetUint16{25, 30, 50, 60, 70, 90, 40},
		},
	}

	Convey("Test Btable methods", t, func(c C) {
		Convey("Iterate over the bTable", func(c C) {
			for _, ti := range table {
				Convey(fmt.Sprintf("Test: %s", ti.name), func() {
					fieldsCount := FieldCount(ti.bTable, 0)

					for i := 0; i < fieldsCount; i++ {
						fn, bo := GetFieldAt(ti.bTable, 0, i)
						So(fn, ShouldEqual, ti.fieldResult[i])
						So(bo, ShouldEqual, ti.offsetResult[i])
					}

				})
			}
		})
		Convey("if index > size valueOffset should panic ", func(c C) {
			defer func() {
				r := recover()
				c.So(r, ShouldNotBeNil)
			}()
			_, vo := GetFieldAt(table[0].bTable, 0, 8)
			So(vo, ShouldEqual, 0)

		})
		Convey("lookup every field of the bTable ", func(c C) {
			ti := table[0]
			for i := 0; i < len(ti.fieldResult); i++ {
				bo, ok := LookupField(ti.bTable, 0, ti.fieldResult[i])
				t.Log(i, ti.fieldResult[i], bo, ok)

				So(ok, ShouldEqual, true)
				So(bo, ShouldEqual, ti.offsetResult[i])
			}

		})

		Convey("Should return false when the field not exists ", func(c C) {
			ti := table[0]
			bo, ok := LookupField(ti.bTable, 0, 125)
			So(ok, ShouldEqual, false)
			So(bo, ShouldEqual, 0)
			bo, ok = LookupField(ti.bTable, 0, 7)
			So(ok, ShouldEqual, false)
			So(bo, ShouldEqual, 0)
		})

	})
}

func BenchmarkSample(b *testing.B) {
	mem := make([]byte, 254*3+1)
	flatbuffers.WriteUint8(mem, 254)
	for i := 0; i < 254; i++ {
		flatbuffers.WriteUint8(mem[1+i*3:], uint8(i))
		flatbuffers.WriteUint16(mem[1+i*3:], uint16(i))

	}
	fmt.Println("47777", FieldCount(mem, 0))

	fmt.Println("5555", FieldCount(mem, 0))
	LookupField(mem, 0, 0)

	b.ResetTimer()
	for i := 0; i < 254; i++ {
		LookupField(mem, 0, delta.FieldNumber(i))

	}
}

/*
func Test_BTable_Validation(t *testing.T) {

    Convey("all ok should not return error", t, func(c C) {
        bTableBinary := []byte{7, 0, 25, 0, 1, 30, 0, 5, 50, 0, 20, 60, 0, 22, 70, 0, 55, 90, 0, 60, 40, 0}
        b:= bTable(0)
        err := b.Validate(r)
        t.Log(err)
        So(err, ShouldBeNil)
    })
    Convey("unsorted fields should return error", t, func(c C) {
        bTableBinary := []byte{7, 0, 25, 0, 1, 70, 0, 30, 50, 0, 20, 60, 0, 22, 70, 0, 55, 90, 0, 60, 40, 0}
        b:= bTable(0)
        err := b.Validate(r)
        t.Log(err)
        So(err, ShouldNotBeNil)
    })

    Convey("bad size should return error", t, func(c C) {
        bTableBinary := []byte{70, 0, 25, 0, 1, 30, 0, 5, 50, 0, 20, 60, 0, 22, 70, 0, 55, 90, 0, 60, 40, 0}
        b:= bTable(0)
        err := b.Validate( r)
        t.Log(err)
        So(err, ShouldNotBeNil)
    })

}*/
