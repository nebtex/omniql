package faker

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	fmocks "github.com/nebtex/omniql/tools/faker/fieldgen/mocks"
	rmocks "github.com/nebtex/omniql/commons/golang/oreflection/mocks"
	//	"fmt"
	"github.com/nebtex/hybrids/golang/hybrids"
)

func Test_FakeEnumeration(t *testing.T) {
	Convey("Test_FakeEnumeration", t, func() {

		/*	Convey("Should Fail the the random choice fail", func() {
				fg := &fmocks.FieldGenerator{}
				f := &Json{fieldGen: fg}
				otype := &rmocks.OType{}
				enumType := &rmocks.OType{}

				fieldMock := &rmocks.Field{}
				otype.On("Field").Return(fieldMock)
				otype.On("Id").Return("Field/test")
				otype.On("Package").Return("test.package")

				fieldMock.On("Name").Return("field")
				fieldMock.On("HybridType").Return(hybrids.Uint32)

				enumGen := &fmocks.EnumerationGenerator{}
				enumGen.On("ShouldGenerateString", "test.field", otype).Return(false, fmt.Errorf("entropy error"))
				fg.On("Enumeration").Return(enumGen)

				out := map[string]interface{}{}

				err := f.fakeEnum("test.field", out, otype, enumType)

				So(err, ShouldNotBeNil)
				enumGen.AssertCalled(t, "ShouldGenerateString", "test.field", otype)

			})*/

		Convey("if a string is selected, it should write a string in the json", func() {
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}
			enumType := &rmocks.OType{}

			fieldMock := &rmocks.Field{}
			otype.On("Field").Return(fieldMock)
			otype.On("Id").Return("Field/test")
			otype.On("Package").Return("test.package")

			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Uint32)

			enumGen := &fmocks.EnumerationGenerator{}
			enumGen.On("ShouldGenerateString", "test.field", otype).Return(true, nil)
			enumGen.On("StringEnumeration", "test.field", otype).Return("PASS", nil)

			fg.On("Enumeration").Return(enumGen)

			out := map[string]interface{}{}

			err := f.fakeEnum("test.field", out, otype, enumType)

			So(err, ShouldBeNil)
			enumGen.AssertCalled(t, "ShouldGenerateString", "test.field", otype)
			enumGen.AssertCalled(t, "StringEnumeration", "test.field", otype)
			So(out["field"], ShouldEqual, "PASS")

		})

		Convey("should write a float64 in other cases", func() {
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}
			enumType := &rmocks.OType{}

			fieldMock := &rmocks.Field{}
			otype.On("Field").Return(fieldMock)
			otype.On("Id").Return("Field/test")
			otype.On("Package").Return("test.package")

			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Uint32)

			enumGen := &fmocks.EnumerationGenerator{}
			enumGen.On("ShouldGenerateString", "test.field", otype).Return(false, nil)
			enumGen.On("StringEnumeration", "test.field", otype).Return("PASS", nil)

			fg.On("Enumeration").Return(enumGen)

			out := map[string]interface{}{}

			err := f.fakeEnum("test.field", out, otype, enumType)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, "PASS")

		})
	})
}
