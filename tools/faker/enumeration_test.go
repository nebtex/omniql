package faker

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	fmocks "github.com/nebtex/omniql/tools/faker/mocks"
	rmocks "github.com/nebtex/omniql/commons/golang/oreflection/mocks"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
)

func Test_FakeEnumeration(t *testing.T) {
	Convey("Test_FakeEnumeration", t, func() {

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
			otype.On("Package").Return("io.test.app")

			fg.On("Boolean", "test.field", otype, fieldNumber).Return(true, fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeBoolean("test.field", out, "field", otype, fieldNumber)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.Package, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Boolean)
			So(ef.OmniqlType, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})
	})
}
