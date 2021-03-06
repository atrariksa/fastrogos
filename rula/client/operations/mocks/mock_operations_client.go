// Code generated by MockGen. DO NOT EDIT.
// Source: ./client/operations/operations_client.go

// Package mock_operations is a generated GoMock package.
package mock_operations

import (
	reflect "reflect"

	operations "github.com/atrariksa/fastrogos/rula/client/operations"
	runtime "github.com/go-openapi/runtime"
	gomock "github.com/golang/mock/gomock"
)

// MockClientService is a mock of ClientService interface.
type MockClientService struct {
	ctrl     *gomock.Controller
	recorder *MockClientServiceMockRecorder
}

// MockClientServiceMockRecorder is the mock recorder for MockClientService.
type MockClientServiceMockRecorder struct {
	mock *MockClientService
}

// NewMockClientService creates a new mock instance.
func NewMockClientService(ctrl *gomock.Controller) *MockClientService {
	mock := &MockClientService{ctrl: ctrl}
	mock.recorder = &MockClientServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientService) EXPECT() *MockClientServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockClientService) CreateUser(params *operations.CreateUserParams, opts ...operations.ClientOption) (*operations.CreateUserCreated, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateUser", varargs...)
	ret0, _ := ret[0].(*operations.CreateUserCreated)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockClientServiceMockRecorder) CreateUser(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockClientService)(nil).CreateUser), varargs...)
}

// DeleteUser mocks base method.
func (m *MockClientService) DeleteUser(params *operations.DeleteUserParams, opts ...operations.ClientOption) (*operations.DeleteUserOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteUser", varargs...)
	ret0, _ := ret[0].(*operations.DeleteUserOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockClientServiceMockRecorder) DeleteUser(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockClientService)(nil).DeleteUser), varargs...)
}

// General mocks base method.
func (m *MockClientService) General(params *operations.GeneralParams, opts ...operations.ClientOption) (*operations.GeneralOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "General", varargs...)
	ret0, _ := ret[0].(*operations.GeneralOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// General indicates an expected call of General.
func (mr *MockClientServiceMockRecorder) General(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "General", reflect.TypeOf((*MockClientService)(nil).General), varargs...)
}

// Login mocks base method.
func (m *MockClientService) Login(params *operations.LoginParams, opts ...operations.ClientOption) (*operations.LoginOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Login", varargs...)
	ret0, _ := ret[0].(*operations.LoginOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockClientServiceMockRecorder) Login(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockClientService)(nil).Login), varargs...)
}

// SetTransport mocks base method.
func (m *MockClientService) SetTransport(transport runtime.ClientTransport) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTransport", transport)
}

// SetTransport indicates an expected call of SetTransport.
func (mr *MockClientServiceMockRecorder) SetTransport(transport interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTransport", reflect.TypeOf((*MockClientService)(nil).SetTransport), transport)
}

// UpdateUser mocks base method.
func (m *MockClientService) UpdateUser(params *operations.UpdateUserParams, opts ...operations.ClientOption) (*operations.UpdateUserOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateUser", varargs...)
	ret0, _ := ret[0].(*operations.UpdateUserOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockClientServiceMockRecorder) UpdateUser(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockClientService)(nil).UpdateUser), varargs...)
}
