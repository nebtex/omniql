// Code generated, DO NOT EDIT.
package faker

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	fmocks "github.com/nebtex/omniql/tools/faker/mocks"
	rmocks "github.com/nebtex/omniql/commons/golang/oreflection/mocks"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
)


func Test_FakeBoolean(t *testing.T) {
	Convey("Test_FakeBoolean", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			fg.On("Boolean", "test.field", otype, fieldNumber).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeBoolean("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, true)
			fg.AssertCalled(t, "Boolean", "test.field", otype, fieldNumber)

		})

		Convey("Should return error if the generator returns an error", func() {

			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}
			otype.On("Id").Return("Table/Test")
			otype.On("Application").Return("io.test.app")

			fg.On("Boolean", "test.field", otype, fieldNumber).Return(true, fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeBoolean("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Boolean)
			So(ef.OmniqlType, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorBoolean(t *testing.T) {
	Convey("Test_FakeVectorBoolean", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorBoolean("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorBoolean)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorBoolean("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorBoolean("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorBoolean)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Boolean", "test.field", otype, fieldNumber).Return(true, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorBoolean("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Boolean)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Boolean", "test.field", otype, fieldNumber).Return(true, nil)

			out := map[string]interface{}{}

			err := f.fakeVectorBoolean("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, true)
		})
	})
}



func Test_FakeInt8(t *testing.T) {
	Convey("Test_FakeInt8", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			fg.On("Int8", "test.field", otype, fieldNumber).Return(int8(-60), nil)
			out := map[string]interface{}{}

			err := f.fakeInt8("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, float64(-60))
			fg.AssertCalled(t, "Int8", "test.field", otype, fieldNumber)

		})

		Convey("Should return error if the generator returns an error", func() {

			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}
			otype.On("Id").Return("Table/Test")
			otype.On("Application").Return("io.test.app")

			fg.On("Int8", "test.field", otype, fieldNumber).Return(int8(-60), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeInt8("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Int8)
			So(ef.OmniqlType, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorInt8(t *testing.T) {
	Convey("Test_FakeVectorInt8", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorInt8("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorInt8)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorInt8("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorInt8("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorInt8)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Int8", "test.field", otype, fieldNumber).Return(int8(-60), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorInt8("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Int8)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Int8", "test.field", otype, fieldNumber).Return(int8(-60), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorInt8("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, float64(-60))
		})
	})
}



func Test_FakeUint8(t *testing.T) {
	Convey("Test_FakeUint8", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			fg.On("Uint8", "test.field", otype, fieldNumber).Return(uint8(87), nil)
			out := map[string]interface{}{}

			err := f.fakeUint8("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, float64(87))
			fg.AssertCalled(t, "Uint8", "test.field", otype, fieldNumber)

		})

		Convey("Should return error if the generator returns an error", func() {

			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}
			otype.On("Id").Return("Table/Test")
			otype.On("Application").Return("io.test.app")

			fg.On("Uint8", "test.field", otype, fieldNumber).Return(uint8(87), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeUint8("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Uint8)
			So(ef.OmniqlType, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorUint8(t *testing.T) {
	Convey("Test_FakeVectorUint8", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorUint8("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorUint8)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorUint8("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorUint8("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorUint8)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Uint8", "test.field", otype, fieldNumber).Return(uint8(87), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorUint8("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Uint8)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Uint8", "test.field", otype, fieldNumber).Return(uint8(87), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorUint8("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, float64(87))
		})
	})
}



func Test_FakeInt16(t *testing.T) {
	Convey("Test_FakeInt16", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			fg.On("Int16", "test.field", otype, fieldNumber).Return(int16(-500), nil)
			out := map[string]interface{}{}

			err := f.fakeInt16("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, float64(-500))
			fg.AssertCalled(t, "Int16", "test.field", otype, fieldNumber)

		})

		Convey("Should return error if the generator returns an error", func() {

			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}
			otype.On("Id").Return("Table/Test")
			otype.On("Application").Return("io.test.app")

			fg.On("Int16", "test.field", otype, fieldNumber).Return(int16(-500), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeInt16("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Int16)
			So(ef.OmniqlType, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorInt16(t *testing.T) {
	Convey("Test_FakeVectorInt16", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorInt16("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorInt16)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorInt16("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorInt16("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorInt16)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Int16", "test.field", otype, fieldNumber).Return(int16(-500), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorInt16("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Int16)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Int16", "test.field", otype, fieldNumber).Return(int16(-500), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorInt16("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, float64(-500))
		})
	})
}



func Test_FakeUint16(t *testing.T) {
	Convey("Test_FakeUint16", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			fg.On("Uint16", "test.field", otype, fieldNumber).Return(uint16(458), nil)
			out := map[string]interface{}{}

			err := f.fakeUint16("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, float64(458))
			fg.AssertCalled(t, "Uint16", "test.field", otype, fieldNumber)

		})

		Convey("Should return error if the generator returns an error", func() {

			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}
			otype.On("Id").Return("Table/Test")
			otype.On("Application").Return("io.test.app")

			fg.On("Uint16", "test.field", otype, fieldNumber).Return(uint16(458), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeUint16("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Uint16)
			So(ef.OmniqlType, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorUint16(t *testing.T) {
	Convey("Test_FakeVectorUint16", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorUint16("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorUint16)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorUint16("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorUint16("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorUint16)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Uint16", "test.field", otype, fieldNumber).Return(uint16(458), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorUint16("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Uint16)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Uint16", "test.field", otype, fieldNumber).Return(uint16(458), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorUint16("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, float64(458))
		})
	})
}



func Test_FakeInt32(t *testing.T) {
	Convey("Test_FakeInt32", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			fg.On("Int32", "test.field", otype, fieldNumber).Return(int32(-60000), nil)
			out := map[string]interface{}{}

			err := f.fakeInt32("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, float64(-60000))
			fg.AssertCalled(t, "Int32", "test.field", otype, fieldNumber)

		})

		Convey("Should return error if the generator returns an error", func() {

			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}
			otype.On("Id").Return("Table/Test")
			otype.On("Application").Return("io.test.app")

			fg.On("Int32", "test.field", otype, fieldNumber).Return(int32(-60000), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeInt32("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Int32)
			So(ef.OmniqlType, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorInt32(t *testing.T) {
	Convey("Test_FakeVectorInt32", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorInt32("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorInt32)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorInt32("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorInt32("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorInt32)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Int32", "test.field", otype, fieldNumber).Return(int32(-60000), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorInt32("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Int32)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Int32", "test.field", otype, fieldNumber).Return(int32(-60000), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorInt32("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, float64(-60000))
		})
	})
}



func Test_FakeUint32(t *testing.T) {
	Convey("Test_FakeUint32", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			fg.On("Uint32", "test.field", otype, fieldNumber).Return(uint32(23568778), nil)
			out := map[string]interface{}{}

			err := f.fakeUint32("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, float64(23568778))
			fg.AssertCalled(t, "Uint32", "test.field", otype, fieldNumber)

		})

		Convey("Should return error if the generator returns an error", func() {

			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}
			otype.On("Id").Return("Table/Test")
			otype.On("Application").Return("io.test.app")

			fg.On("Uint32", "test.field", otype, fieldNumber).Return(uint32(23568778), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeUint32("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Uint32)
			So(ef.OmniqlType, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorUint32(t *testing.T) {
	Convey("Test_FakeVectorUint32", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorUint32("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorUint32)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorUint32("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorUint32("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorUint32)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Uint32", "test.field", otype, fieldNumber).Return(uint32(23568778), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorUint32("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Uint32)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Uint32", "test.field", otype, fieldNumber).Return(uint32(23568778), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorUint32("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, float64(23568778))
		})
	})
}



func Test_FakeInt64(t *testing.T) {
	Convey("Test_FakeInt64", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			fg.On("Int64", "test.field", otype, fieldNumber).Return(int64(-54545788), nil)
			out := map[string]interface{}{}

			err := f.fakeInt64("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, "-54545788")
			fg.AssertCalled(t, "Int64", "test.field", otype, fieldNumber)

		})

		Convey("Should return error if the generator returns an error", func() {

			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}
			otype.On("Id").Return("Table/Test")
			otype.On("Application").Return("io.test.app")

			fg.On("Int64", "test.field", otype, fieldNumber).Return(int64(-54545788), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeInt64("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Int64)
			So(ef.OmniqlType, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorInt64(t *testing.T) {
	Convey("Test_FakeVectorInt64", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorInt64("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorInt64)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorInt64("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorInt64("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorInt64)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Int64", "test.field", otype, fieldNumber).Return(int64(-54545788), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorInt64("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Int64)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Int64", "test.field", otype, fieldNumber).Return(int64(-54545788), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorInt64("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, "-54545788")
		})
	})
}



func Test_FakeUint64(t *testing.T) {
	Convey("Test_FakeUint64", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			fg.On("Uint64", "test.field", otype, fieldNumber).Return(uint64(123587), nil)
			out := map[string]interface{}{}

			err := f.fakeUint64("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, "123587")
			fg.AssertCalled(t, "Uint64", "test.field", otype, fieldNumber)

		})

		Convey("Should return error if the generator returns an error", func() {

			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}
			otype.On("Id").Return("Table/Test")
			otype.On("Application").Return("io.test.app")

			fg.On("Uint64", "test.field", otype, fieldNumber).Return(uint64(123587), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeUint64("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Uint64)
			So(ef.OmniqlType, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorUint64(t *testing.T) {
	Convey("Test_FakeVectorUint64", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorUint64("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorUint64)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorUint64("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorUint64("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorUint64)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Uint64", "test.field", otype, fieldNumber).Return(uint64(123587), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorUint64("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Uint64)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Uint64", "test.field", otype, fieldNumber).Return(uint64(123587), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorUint64("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, "123587")
		})
	})
}



func Test_FakeFloat32(t *testing.T) {
	Convey("Test_FakeFloat32", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			fg.On("Float32", "test.field", otype, fieldNumber).Return(float32(25.50), nil)
			out := map[string]interface{}{}

			err := f.fakeFloat32("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, float64(25.50))
			fg.AssertCalled(t, "Float32", "test.field", otype, fieldNumber)

		})

		Convey("Should return error if the generator returns an error", func() {

			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}
			otype.On("Id").Return("Table/Test")
			otype.On("Application").Return("io.test.app")

			fg.On("Float32", "test.field", otype, fieldNumber).Return(float32(25.50), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeFloat32("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Float32)
			So(ef.OmniqlType, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorFloat32(t *testing.T) {
	Convey("Test_FakeVectorFloat32", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorFloat32("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorFloat32)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorFloat32("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorFloat32("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorFloat32)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Float32", "test.field", otype, fieldNumber).Return(float32(25.50), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorFloat32("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Float32)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Float32", "test.field", otype, fieldNumber).Return(float32(25.50), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorFloat32("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, float64(25.50))
		})
	})
}



func Test_FakeFloat64(t *testing.T) {
	Convey("Test_FakeFloat64", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			fg.On("Float64", "test.field", otype, fieldNumber).Return(float64(-356.4545), nil)
			out := map[string]interface{}{}

			err := f.fakeFloat64("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, float64(-356.4545))
			fg.AssertCalled(t, "Float64", "test.field", otype, fieldNumber)

		})

		Convey("Should return error if the generator returns an error", func() {

			fieldNumber := hybrids.FieldNumber(2)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}
			otype.On("Id").Return("Table/Test")
			otype.On("Application").Return("io.test.app")

			fg.On("Float64", "test.field", otype, fieldNumber).Return(float64(-356.4545), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeFloat64("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Float64)
			So(ef.OmniqlType, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorFloat64(t *testing.T) {
	Convey("Test_FakeVectorFloat64", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorFloat64("test.field", out, "field", otype, fieldNumber)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			ef := err.(*Error)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorFloat64)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorFloat64("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorFloat64("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.VectorFloat64)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Float64", "test.field", otype, fieldNumber).Return(float64(-356.4545), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorFloat64("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype, fieldNumber)
			fg.AssertCalled(t, "VectorLen", "test.field", otype, fieldNumber)
			So(ef.Application, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Float64)
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fieldNumber := hybrids.FieldNumber(3)
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			otype.On("Id").Return("Struct/Test")
			otype.On("Application").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype, fieldNumber).Return(false, nil)
			fg.On("VectorLen", "test.field", otype, fieldNumber).Return(10, nil)
			fg.On("Float64", "test.field", otype, fieldNumber).Return(float64(-356.4545), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorFloat64("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, float64(-356.4545))
		})
	})
}

