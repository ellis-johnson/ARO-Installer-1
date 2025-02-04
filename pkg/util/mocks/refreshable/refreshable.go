// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/refreshable (interfaces: Authorizer)

// Package mock_refreshable is a generated GoMock package.
package mock_refreshable

import (
	reflect "reflect"

	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthorizer is a mock of Authorizer interface.
type MockAuthorizer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizerMockRecorder
}

// MockAuthorizerMockRecorder is the mock recorder for MockAuthorizer.
type MockAuthorizerMockRecorder struct {
	mock *MockAuthorizer
}

// NewMockAuthorizer creates a new mock instance.
func NewMockAuthorizer(ctrl *gomock.Controller) *MockAuthorizer {
	mock := &MockAuthorizer{ctrl: ctrl}
	mock.recorder = &MockAuthorizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizer) EXPECT() *MockAuthorizerMockRecorder {
	return m.recorder
}

// Rebuild mocks base method.
func (m *MockAuthorizer) Rebuild() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rebuild")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rebuild indicates an expected call of Rebuild.
func (mr *MockAuthorizerMockRecorder) Rebuild() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rebuild", reflect.TypeOf((*MockAuthorizer)(nil).Rebuild))
}

// WithAuthorization mocks base method.
func (m *MockAuthorizer) WithAuthorization() autorest.PrepareDecorator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithAuthorization")
	ret0, _ := ret[0].(autorest.PrepareDecorator)
	return ret0
}

// WithAuthorization indicates an expected call of WithAuthorization.
func (mr *MockAuthorizerMockRecorder) WithAuthorization() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithAuthorization", reflect.TypeOf((*MockAuthorizer)(nil).WithAuthorization))
}
