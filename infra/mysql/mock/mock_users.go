// Code generated by MockGen. DO NOT EDIT.
// Source: users.go

// Package mock_mysql is a generated GoMock package.
package mock_mysql

import (
	gomock "github.com/golang/mock/gomock"
	mysql "github.com/midnight-trigger/todo/infra/mysql"
	reflect "reflect"
)

// MockIUsers is a mock of IUsers interface
type MockIUsers struct {
	ctrl     *gomock.Controller
	recorder *MockIUsersMockRecorder
}

// MockIUsersMockRecorder is the mock recorder for MockIUsers
type MockIUsersMockRecorder struct {
	mock *MockIUsers
}

// NewMockIUsers creates a new mock instance
func NewMockIUsers(ctrl *gomock.Controller) *MockIUsers {
	mock := &MockIUsers{ctrl: ctrl}
	mock.recorder = &MockIUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUsers) EXPECT() *MockIUsersMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockIUsers) Create(user *mysql.Users) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockIUsersMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUsers)(nil).Create), user)
}
