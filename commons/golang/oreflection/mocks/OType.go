// Code generated by mockery v1.0.0
package mocks

import corev1 "github.com/nebtex/omniql/pkg/next/corev1"
import mock "github.com/stretchr/testify/mock"
import oreflection "github.com/nebtex/omniql/commons/golang/oreflection"

// OType is an autogenerated mock type for the OType type
type OType struct {
	mock.Mock
}

// Enumeration provides a mock function with given fields:
func (_m *OType) Enumeration() oreflection.Enumeration {
	ret := _m.Called()

	var r0 oreflection.Enumeration
	if rf, ok := ret.Get(0).(func() oreflection.Enumeration); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oreflection.Enumeration)
		}
	}

	return r0
}

// Field provides a mock function with given fields:
func (_m *OType) Field() oreflection.Field {
	ret := _m.Called()

	var r0 oreflection.Field
	if rf, ok := ret.Get(0).(func() oreflection.Field); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oreflection.Field)
		}
	}

	return r0
}

// Id provides a mock function with given fields:
func (_m *OType) Id() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Kind provides a mock function with given fields:
func (_m *OType) Kind() corev1.SchemaTypes {
	ret := _m.Called()

	var r0 corev1.SchemaTypes
	if rf, ok := ret.Get(0).(func() corev1.SchemaTypes); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(corev1.SchemaTypes)
	}

	return r0
}

// Package provides a mock function with given fields:
func (_m *OType) Package() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Resource provides a mock function with given fields:
func (_m *OType) Resource() oreflection.Resource {
	ret := _m.Called()

	var r0 oreflection.Resource
	if rf, ok := ret.Get(0).(func() oreflection.Resource); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oreflection.Resource)
		}
	}

	return r0
}

// Table provides a mock function with given fields:
func (_m *OType) Table() oreflection.Table {
	ret := _m.Called()

	var r0 oreflection.Table
	if rf, ok := ret.Get(0).(func() oreflection.Table); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oreflection.Table)
		}
	}

	return r0
}

// Union provides a mock function with given fields:
func (_m *OType) Union() oreflection.Union {
	ret := _m.Called()

	var r0 oreflection.Union
	if rf, ok := ret.Get(0).(func() oreflection.Union); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oreflection.Union)
		}
	}

	return r0
}
