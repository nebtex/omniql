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
	Convey("Should populate field with the data of FieldGenerator", t, func() {
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

	Convey("Should return error if the generator returns an error", t, func() {

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

}



func Test_FakeInt8(t *testing.T) {
	Convey("Should populate field with the data of FieldGenerator", t, func() {
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

	Convey("Should return error if the generator returns an error", t, func() {

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

}



func Test_FakeUint8(t *testing.T) {
	Convey("Should populate field with the data of FieldGenerator", t, func() {
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

	Convey("Should return error if the generator returns an error", t, func() {

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

}



func Test_FakeInt16(t *testing.T) {
	Convey("Should populate field with the data of FieldGenerator", t, func() {
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

	Convey("Should return error if the generator returns an error", t, func() {

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

}



func Test_FakeUint16(t *testing.T) {
	Convey("Should populate field with the data of FieldGenerator", t, func() {
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

	Convey("Should return error if the generator returns an error", t, func() {

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

}



func Test_FakeInt32(t *testing.T) {
	Convey("Should populate field with the data of FieldGenerator", t, func() {
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

	Convey("Should return error if the generator returns an error", t, func() {

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

}



func Test_FakeUint32(t *testing.T) {
	Convey("Should populate field with the data of FieldGenerator", t, func() {
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

	Convey("Should return error if the generator returns an error", t, func() {

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

}



func Test_FakeInt64(t *testing.T) {
	Convey("Should populate field with the data of FieldGenerator", t, func() {
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

	Convey("Should return error if the generator returns an error", t, func() {

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

}



func Test_FakeUint64(t *testing.T) {
	Convey("Should populate field with the data of FieldGenerator", t, func() {
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

	Convey("Should return error if the generator returns an error", t, func() {

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

}



func Test_FakeFloat32(t *testing.T) {
	Convey("Should populate field with the data of FieldGenerator", t, func() {
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

	Convey("Should return error if the generator returns an error", t, func() {

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

}



func Test_FakeFloat64(t *testing.T) {
	Convey("Should populate field with the data of FieldGenerator", t, func() {
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

	Convey("Should return error if the generator returns an error", t, func() {

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

}

