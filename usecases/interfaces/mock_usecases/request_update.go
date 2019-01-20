// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/int128/gradleupdate/usecases/interfaces (interfaces: RequestUpdate)

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	domain "github.com/int128/gradleupdate/domain"
	reflect "reflect"
)

// MockRequestUpdate is a mock of RequestUpdate interface
type MockRequestUpdate struct {
	ctrl     *gomock.Controller
	recorder *MockRequestUpdateMockRecorder
}

// MockRequestUpdateMockRecorder is the mock recorder for MockRequestUpdate
type MockRequestUpdateMockRecorder struct {
	mock *MockRequestUpdate
}

// NewMockRequestUpdate creates a new mock instance
func NewMockRequestUpdate(ctrl *gomock.Controller) *MockRequestUpdate {
	mock := &MockRequestUpdate{ctrl: ctrl}
	mock.recorder = &MockRequestUpdateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRequestUpdate) EXPECT() *MockRequestUpdateMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockRequestUpdate) Do(arg0 context.Context, arg1 domain.RepositoryID, arg2 string) error {
	ret := m.ctrl.Call(m, "Do", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do
func (mr *MockRequestUpdateMockRecorder) Do(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockRequestUpdate)(nil).Do), arg0, arg1, arg2)
}
