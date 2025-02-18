// Code generated by mockery v2.21.4. DO NOT EDIT.

package mocks

import (
	config "github.com/KarnerTh/mermerd/config"
	mock "github.com/stretchr/testify/mock"
)

// MermerdConfig is an autogenerated mock type for the MermerdConfig type
type MermerdConfig struct {
	mock.Mock
}

// ConnectionString provides a mock function with given fields:
func (_m *MermerdConfig) ConnectionString() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ConnectionStringSuggestions provides a mock function with given fields:
func (_m *MermerdConfig) ConnectionStringSuggestions() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// Debug provides a mock function with given fields:
func (_m *MermerdConfig) Debug() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// EncloseWithMermaidBackticks provides a mock function with given fields:
func (_m *MermerdConfig) EncloseWithMermaidBackticks() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IgnoreTables provides a mock function with given fields:
func (_m *MermerdConfig) IgnoreTables() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// OmitAttributeKeys provides a mock function with given fields:
func (_m *MermerdConfig) OmitAttributeKeys() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OmitConstraintLabels provides a mock function with given fields:
func (_m *MermerdConfig) OmitConstraintLabels() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OutputFileName provides a mock function with given fields:
func (_m *MermerdConfig) OutputFileName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// OutputMode provides a mock function with given fields:
func (_m *MermerdConfig) OutputMode() config.OutputModeType {
	ret := _m.Called()

	var r0 config.OutputModeType
	if rf, ok := ret.Get(0).(func() config.OutputModeType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(config.OutputModeType)
	}

	return r0
}

// RelationshipLabels provides a mock function with given fields:
func (_m *MermerdConfig) RelationshipLabels() []config.RelationshipLabel {
	ret := _m.Called()

	var r0 []config.RelationshipLabel
	if rf, ok := ret.Get(0).(func() []config.RelationshipLabel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]config.RelationshipLabel)
		}
	}

	return r0
}

// SchemaPrefixSeparator provides a mock function with given fields:
func (_m *MermerdConfig) SchemaPrefixSeparator() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Schemas provides a mock function with given fields:
func (_m *MermerdConfig) Schemas() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// SelectedTables provides a mock function with given fields:
func (_m *MermerdConfig) SelectedTables() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// ShowAllConstraints provides a mock function with given fields:
func (_m *MermerdConfig) ShowAllConstraints() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ShowDescriptions provides a mock function with given fields:
func (_m *MermerdConfig) ShowDescriptions() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// ShowSchemaPrefix provides a mock function with given fields:
func (_m *MermerdConfig) ShowSchemaPrefix() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// UseAllSchemas provides a mock function with given fields:
func (_m *MermerdConfig) UseAllSchemas() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// UseAllTables provides a mock function with given fields:
func (_m *MermerdConfig) UseAllTables() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewMermerdConfig interface {
	mock.TestingT
	Cleanup(func())
}

// NewMermerdConfig creates a new instance of MermerdConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMermerdConfig(t mockConstructorTestingTNewMermerdConfig) *MermerdConfig {
	mock := &MermerdConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
