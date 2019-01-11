// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/int128/gradleupdate/gateways/interfaces (interfaces: RepositoryRepository)

// Package mock_gateways is a generated GoMock package.
package mock_gateways

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	domain "github.com/int128/gradleupdate/domain"
	reflect "reflect"
)

// MockRepositoryRepository is a mock of RepositoryRepository interface
type MockRepositoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryRepositoryMockRecorder
}

// MockRepositoryRepositoryMockRecorder is the mock recorder for MockRepositoryRepository
type MockRepositoryRepositoryMockRecorder struct {
	mock *MockRepositoryRepository
}

// NewMockRepositoryRepository creates a new mock instance
func NewMockRepositoryRepository(ctrl *gomock.Controller) *MockRepositoryRepository {
	mock := &MockRepositoryRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepositoryRepository) EXPECT() *MockRepositoryRepositoryMockRecorder {
	return m.recorder
}

// Fork mocks base method
func (m *MockRepositoryRepository) Fork(arg0 context.Context, arg1 domain.RepositoryID) (*domain.Repository, error) {
	ret := m.ctrl.Call(m, "Fork", arg0, arg1)
	ret0, _ := ret[0].(*domain.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fork indicates an expected call of Fork
func (mr *MockRepositoryRepositoryMockRecorder) Fork(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fork", reflect.TypeOf((*MockRepositoryRepository)(nil).Fork), arg0, arg1)
}

// Get mocks base method
func (m *MockRepositoryRepository) Get(arg0 context.Context, arg1 domain.RepositoryID) (*domain.Repository, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*domain.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryRepositoryMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepositoryRepository)(nil).Get), arg0, arg1)
}

// GetBranch mocks base method
func (m *MockRepositoryRepository) GetBranch(arg0 context.Context, arg1 domain.BranchID) (*domain.Branch, error) {
	ret := m.ctrl.Call(m, "GetBranch", arg0, arg1)
	ret0, _ := ret[0].(*domain.Branch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBranch indicates an expected call of GetBranch
func (mr *MockRepositoryRepositoryMockRecorder) GetBranch(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBranch", reflect.TypeOf((*MockRepositoryRepository)(nil).GetBranch), arg0, arg1)
}

// GetFileContent mocks base method
func (m *MockRepositoryRepository) GetFileContent(arg0 context.Context, arg1 domain.RepositoryID, arg2 string) (domain.FileContent, error) {
	ret := m.ctrl.Call(m, "GetFileContent", arg0, arg1, arg2)
	ret0, _ := ret[0].(domain.FileContent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileContent indicates an expected call of GetFileContent
func (mr *MockRepositoryRepositoryMockRecorder) GetFileContent(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileContent", reflect.TypeOf((*MockRepositoryRepository)(nil).GetFileContent), arg0, arg1, arg2)
}

// IsNotFoundError mocks base method
func (m *MockRepositoryRepository) IsNotFoundError(arg0 error) bool {
	ret := m.ctrl.Call(m, "IsNotFoundError", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNotFoundError indicates an expected call of IsNotFoundError
func (mr *MockRepositoryRepositoryMockRecorder) IsNotFoundError(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNotFoundError", reflect.TypeOf((*MockRepositoryRepository)(nil).IsNotFoundError), arg0)
}