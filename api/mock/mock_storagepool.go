// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/api (interfaces: OpenStoragePoolServer,OpenStoragePoolClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	api "github.com/libopenstorage/openstorage/api"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockOpenStoragePoolServer is a mock of OpenStoragePoolServer interface
type MockOpenStoragePoolServer struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStoragePoolServerMockRecorder
}

// MockOpenStoragePoolServerMockRecorder is the mock recorder for MockOpenStoragePoolServer
type MockOpenStoragePoolServerMockRecorder struct {
	mock *MockOpenStoragePoolServer
}

// NewMockOpenStoragePoolServer creates a new mock instance
func NewMockOpenStoragePoolServer(ctrl *gomock.Controller) *MockOpenStoragePoolServer {
	mock := &MockOpenStoragePoolServer{ctrl: ctrl}
	mock.recorder = &MockOpenStoragePoolServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpenStoragePoolServer) EXPECT() *MockOpenStoragePoolServerMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockOpenStoragePoolServer) Delete(arg0 context.Context, arg1 *api.SdkStoragePoolDeleteRequest) (*api.SdkStoragePoolDeleteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkStoragePoolDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockOpenStoragePoolServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOpenStoragePoolServer)(nil).Delete), arg0, arg1)
}

// Resize mocks base method
func (m *MockOpenStoragePoolServer) Resize(arg0 context.Context, arg1 *api.SdkStoragePoolResizeRequest) (*api.SdkStoragePoolResizeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resize", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkStoragePoolResizeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resize indicates an expected call of Resize
func (mr *MockOpenStoragePoolServerMockRecorder) Resize(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resize", reflect.TypeOf((*MockOpenStoragePoolServer)(nil).Resize), arg0, arg1)
}

// UpdateIOPS mocks base method
func (m *MockOpenStoragePoolServer) UpdateIOPS(arg0 context.Context, arg1 *api.SdkStoragePoolUpdateIOPSRequest) (*api.SdkStoragePoolUpdateIOPSResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIOPS", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkStoragePoolUpdateIOPSResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIOPS indicates an expected call of UpdateIOPS
func (mr *MockOpenStoragePoolServerMockRecorder) UpdateIOPS(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIOPS", reflect.TypeOf((*MockOpenStoragePoolServer)(nil).UpdateIOPS), arg0, arg1)
}

// MockOpenStoragePoolClient is a mock of OpenStoragePoolClient interface
type MockOpenStoragePoolClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStoragePoolClientMockRecorder
}

// MockOpenStoragePoolClientMockRecorder is the mock recorder for MockOpenStoragePoolClient
type MockOpenStoragePoolClientMockRecorder struct {
	mock *MockOpenStoragePoolClient
}

// NewMockOpenStoragePoolClient creates a new mock instance
func NewMockOpenStoragePoolClient(ctrl *gomock.Controller) *MockOpenStoragePoolClient {
	mock := &MockOpenStoragePoolClient{ctrl: ctrl}
	mock.recorder = &MockOpenStoragePoolClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpenStoragePoolClient) EXPECT() *MockOpenStoragePoolClientMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockOpenStoragePoolClient) Delete(arg0 context.Context, arg1 *api.SdkStoragePoolDeleteRequest, arg2 ...grpc.CallOption) (*api.SdkStoragePoolDeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*api.SdkStoragePoolDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockOpenStoragePoolClientMockRecorder) Delete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOpenStoragePoolClient)(nil).Delete), varargs...)
}

// Resize mocks base method
func (m *MockOpenStoragePoolClient) Resize(arg0 context.Context, arg1 *api.SdkStoragePoolResizeRequest, arg2 ...grpc.CallOption) (*api.SdkStoragePoolResizeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Resize", varargs...)
	ret0, _ := ret[0].(*api.SdkStoragePoolResizeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resize indicates an expected call of Resize
func (mr *MockOpenStoragePoolClientMockRecorder) Resize(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resize", reflect.TypeOf((*MockOpenStoragePoolClient)(nil).Resize), varargs...)
}

// UpdateIOPS mocks base method
func (m *MockOpenStoragePoolClient) UpdateIOPS(arg0 context.Context, arg1 *api.SdkStoragePoolUpdateIOPSRequest, arg2 ...grpc.CallOption) (*api.SdkStoragePoolUpdateIOPSResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateIOPS", varargs...)
	ret0, _ := ret[0].(*api.SdkStoragePoolUpdateIOPSResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIOPS indicates an expected call of UpdateIOPS
func (mr *MockOpenStoragePoolClientMockRecorder) UpdateIOPS(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIOPS", reflect.TypeOf((*MockOpenStoragePoolClient)(nil).UpdateIOPS), varargs...)
}
