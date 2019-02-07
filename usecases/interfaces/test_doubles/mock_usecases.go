// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/int128/gradleupdate/usecases/interfaces (interfaces: GetBadge,GetBadgeError,GetRepository,GetRepositoryError,SendUpdate,SendUpdateError,BatchSendUpdates,SendPullRequest)

// Package usecases is a generated GoMock package.
package usecases

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	git "github.com/int128/gradleupdate/domain/git"
	gradleupdate "github.com/int128/gradleupdate/domain/gradleupdate"
	interfaces "github.com/int128/gradleupdate/usecases/interfaces"
	reflect "reflect"
)

// MockGetBadge is a mock of GetBadge interface
type MockGetBadge struct {
	ctrl     *gomock.Controller
	recorder *MockGetBadgeMockRecorder
}

// MockGetBadgeMockRecorder is the mock recorder for MockGetBadge
type MockGetBadgeMockRecorder struct {
	mock *MockGetBadge
}

// NewMockGetBadge creates a new mock instance
func NewMockGetBadge(ctrl *gomock.Controller) *MockGetBadge {
	mock := &MockGetBadge{ctrl: ctrl}
	mock.recorder = &MockGetBadgeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetBadge) EXPECT() *MockGetBadgeMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockGetBadge) Do(arg0 context.Context, arg1 git.RepositoryID) (*interfaces.GetBadgeResponse, error) {
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(*interfaces.GetBadgeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do
func (mr *MockGetBadgeMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockGetBadge)(nil).Do), arg0, arg1)
}

// MockGetBadgeError is a mock of GetBadgeError interface
type MockGetBadgeError struct {
	ctrl     *gomock.Controller
	recorder *MockGetBadgeErrorMockRecorder
}

// MockGetBadgeErrorMockRecorder is the mock recorder for MockGetBadgeError
type MockGetBadgeErrorMockRecorder struct {
	mock *MockGetBadgeError
}

// NewMockGetBadgeError creates a new mock instance
func NewMockGetBadgeError(ctrl *gomock.Controller) *MockGetBadgeError {
	mock := &MockGetBadgeError{ctrl: ctrl}
	mock.recorder = &MockGetBadgeErrorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetBadgeError) EXPECT() *MockGetBadgeErrorMockRecorder {
	return m.recorder
}

// Error mocks base method
func (m *MockGetBadgeError) Error() string {
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error
func (mr *MockGetBadgeErrorMockRecorder) Error() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockGetBadgeError)(nil).Error))
}

// NoGradleVersion mocks base method
func (m *MockGetBadgeError) NoGradleVersion() bool {
	ret := m.ctrl.Call(m, "NoGradleVersion")
	ret0, _ := ret[0].(bool)
	return ret0
}

// NoGradleVersion indicates an expected call of NoGradleVersion
func (mr *MockGetBadgeErrorMockRecorder) NoGradleVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoGradleVersion", reflect.TypeOf((*MockGetBadgeError)(nil).NoGradleVersion))
}

// MockGetRepository is a mock of GetRepository interface
type MockGetRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGetRepositoryMockRecorder
}

// MockGetRepositoryMockRecorder is the mock recorder for MockGetRepository
type MockGetRepositoryMockRecorder struct {
	mock *MockGetRepository
}

// NewMockGetRepository creates a new mock instance
func NewMockGetRepository(ctrl *gomock.Controller) *MockGetRepository {
	mock := &MockGetRepository{ctrl: ctrl}
	mock.recorder = &MockGetRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetRepository) EXPECT() *MockGetRepositoryMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockGetRepository) Do(arg0 context.Context, arg1 git.RepositoryID) (*interfaces.GetRepositoryResponse, error) {
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(*interfaces.GetRepositoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do
func (mr *MockGetRepositoryMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockGetRepository)(nil).Do), arg0, arg1)
}

// MockGetRepositoryError is a mock of GetRepositoryError interface
type MockGetRepositoryError struct {
	ctrl     *gomock.Controller
	recorder *MockGetRepositoryErrorMockRecorder
}

// MockGetRepositoryErrorMockRecorder is the mock recorder for MockGetRepositoryError
type MockGetRepositoryErrorMockRecorder struct {
	mock *MockGetRepositoryError
}

// NewMockGetRepositoryError creates a new mock instance
func NewMockGetRepositoryError(ctrl *gomock.Controller) *MockGetRepositoryError {
	mock := &MockGetRepositoryError{ctrl: ctrl}
	mock.recorder = &MockGetRepositoryErrorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetRepositoryError) EXPECT() *MockGetRepositoryErrorMockRecorder {
	return m.recorder
}

// Error mocks base method
func (m *MockGetRepositoryError) Error() string {
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error
func (mr *MockGetRepositoryErrorMockRecorder) Error() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockGetRepositoryError)(nil).Error))
}

// NoSuchRepository mocks base method
func (m *MockGetRepositoryError) NoSuchRepository() bool {
	ret := m.ctrl.Call(m, "NoSuchRepository")
	ret0, _ := ret[0].(bool)
	return ret0
}

// NoSuchRepository indicates an expected call of NoSuchRepository
func (mr *MockGetRepositoryErrorMockRecorder) NoSuchRepository() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoSuchRepository", reflect.TypeOf((*MockGetRepositoryError)(nil).NoSuchRepository))
}

// MockSendUpdate is a mock of SendUpdate interface
type MockSendUpdate struct {
	ctrl     *gomock.Controller
	recorder *MockSendUpdateMockRecorder
}

// MockSendUpdateMockRecorder is the mock recorder for MockSendUpdate
type MockSendUpdateMockRecorder struct {
	mock *MockSendUpdate
}

// NewMockSendUpdate creates a new mock instance
func NewMockSendUpdate(ctrl *gomock.Controller) *MockSendUpdate {
	mock := &MockSendUpdate{ctrl: ctrl}
	mock.recorder = &MockSendUpdateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSendUpdate) EXPECT() *MockSendUpdateMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockSendUpdate) Do(arg0 context.Context, arg1 git.RepositoryID) error {
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do
func (mr *MockSendUpdateMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockSendUpdate)(nil).Do), arg0, arg1)
}

// MockSendUpdateError is a mock of SendUpdateError interface
type MockSendUpdateError struct {
	ctrl     *gomock.Controller
	recorder *MockSendUpdateErrorMockRecorder
}

// MockSendUpdateErrorMockRecorder is the mock recorder for MockSendUpdateError
type MockSendUpdateErrorMockRecorder struct {
	mock *MockSendUpdateError
}

// NewMockSendUpdateError creates a new mock instance
func NewMockSendUpdateError(ctrl *gomock.Controller) *MockSendUpdateError {
	mock := &MockSendUpdateError{ctrl: ctrl}
	mock.recorder = &MockSendUpdateErrorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSendUpdateError) EXPECT() *MockSendUpdateErrorMockRecorder {
	return m.recorder
}

// Error mocks base method
func (m *MockSendUpdateError) Error() string {
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error
func (mr *MockSendUpdateErrorMockRecorder) Error() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockSendUpdateError)(nil).Error))
}

// PreconditionViolation mocks base method
func (m *MockSendUpdateError) PreconditionViolation() gradleupdate.PreconditionViolation {
	ret := m.ctrl.Call(m, "PreconditionViolation")
	ret0, _ := ret[0].(gradleupdate.PreconditionViolation)
	return ret0
}

// PreconditionViolation indicates an expected call of PreconditionViolation
func (mr *MockSendUpdateErrorMockRecorder) PreconditionViolation() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreconditionViolation", reflect.TypeOf((*MockSendUpdateError)(nil).PreconditionViolation))
}

// MockBatchSendUpdates is a mock of BatchSendUpdates interface
type MockBatchSendUpdates struct {
	ctrl     *gomock.Controller
	recorder *MockBatchSendUpdatesMockRecorder
}

// MockBatchSendUpdatesMockRecorder is the mock recorder for MockBatchSendUpdates
type MockBatchSendUpdatesMockRecorder struct {
	mock *MockBatchSendUpdates
}

// NewMockBatchSendUpdates creates a new mock instance
func NewMockBatchSendUpdates(ctrl *gomock.Controller) *MockBatchSendUpdates {
	mock := &MockBatchSendUpdates{ctrl: ctrl}
	mock.recorder = &MockBatchSendUpdatesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBatchSendUpdates) EXPECT() *MockBatchSendUpdatesMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockBatchSendUpdates) Do(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Do", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do
func (mr *MockBatchSendUpdatesMockRecorder) Do(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockBatchSendUpdates)(nil).Do), arg0)
}

// MockSendPullRequest is a mock of SendPullRequest interface
type MockSendPullRequest struct {
	ctrl     *gomock.Controller
	recorder *MockSendPullRequestMockRecorder
}

// MockSendPullRequestMockRecorder is the mock recorder for MockSendPullRequest
type MockSendPullRequestMockRecorder struct {
	mock *MockSendPullRequest
}

// NewMockSendPullRequest creates a new mock instance
func NewMockSendPullRequest(ctrl *gomock.Controller) *MockSendPullRequest {
	mock := &MockSendPullRequest{ctrl: ctrl}
	mock.recorder = &MockSendPullRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSendPullRequest) EXPECT() *MockSendPullRequestMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockSendPullRequest) Do(arg0 context.Context, arg1 interfaces.SendPullRequestRequest) error {
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do
func (mr *MockSendPullRequestMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockSendPullRequest)(nil).Do), arg0, arg1)
}
