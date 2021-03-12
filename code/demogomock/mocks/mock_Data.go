// Code generated by MockGen. DO NOT EDIT.
// Source: spike-testing-golang/code/demogomock (interfaces: Data)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockData is a mock of Data interface
type MockData struct {
	ctrl     *gomock.Controller
	recorder *MockDataMockRecorder
}

// MockDataMockRecorder is the mock recorder for MockData
type MockDataMockRecorder struct {
	mock *MockData
}

// NewMockData creates a new mock instance
func NewMockData(ctrl *gomock.Controller) *MockData {
	mock := &MockData{ctrl: ctrl}
	mock.recorder = &MockDataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockData) EXPECT() *MockDataMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockData) Get(arg0 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDataMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockData)(nil).Get), arg0)
}

// Send mocks base method
func (m *MockData) Send(arg0 string, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockDataMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockData)(nil).Send), arg0, arg1)
}