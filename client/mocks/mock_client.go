// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jannickfahlbusch/owntracks-go/client (interfaces: Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	types "github.com/jannickfahlbusch/owntracks-go/types"
	reflect "reflect"
	time "time"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Devices mocks base method
func (m *MockClient) Devices(arg0 context.Context, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Devices", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Devices indicates an expected call of Devices
func (mr *MockClientMockRecorder) Devices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Devices", reflect.TypeOf((*MockClient)(nil).Devices), arg0, arg1)
}

// Locations mocks base method
func (m *MockClient) Locations(arg0 context.Context, arg1, arg2 string, arg3, arg4 time.Time) (*types.LocationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Locations", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*types.LocationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Locations indicates an expected call of Locations
func (mr *MockClientMockRecorder) Locations(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Locations", reflect.TypeOf((*MockClient)(nil).Locations), arg0, arg1, arg2, arg3, arg4)
}

// Publish mocks base method
func (m *MockClient) Publish(arg0 context.Context, arg1, arg2 string, arg3 *types.Location) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish
func (mr *MockClientMockRecorder) Publish(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockClient)(nil).Publish), arg0, arg1, arg2, arg3)
}

// Users mocks base method
func (m *MockClient) Users(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Users", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Users indicates an expected call of Users
func (mr *MockClientMockRecorder) Users(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockClient)(nil).Users), arg0)
}

// Version mocks base method
func (m *MockClient) Version(arg0 context.Context) (*types.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", arg0)
	ret0, _ := ret[0].(*types.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (mr *MockClientMockRecorder) Version(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockClient)(nil).Version), arg0)
}
