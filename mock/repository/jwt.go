// Code generated by MockGen. DO NOT EDIT.
// Source: jwt.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	entity "github.com/akubi0w1/golang-sample/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockJWT is a mock of JWT interface.
type MockJWT struct {
	ctrl     *gomock.Controller
	recorder *MockJWTMockRecorder
}

// MockJWTMockRecorder is the mock recorder for MockJWT.
type MockJWTMockRecorder struct {
	mock *MockJWT
}

// NewMockJWT creates a new mock instance.
func NewMockJWT(ctrl *gomock.Controller) *MockJWT {
	mock := &MockJWT{ctrl: ctrl}
	mock.recorder = &MockJWTMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJWT) EXPECT() *MockJWTMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockJWT) Generate(claims entity.Claims) (entity.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", claims)
	ret0, _ := ret[0].(entity.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockJWTMockRecorder) Generate(claims interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockJWT)(nil).Generate), claims)
}

// Parse mocks base method.
func (m *MockJWT) Parse(token entity.Token) (entity.Claims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", token)
	ret0, _ := ret[0].(entity.Claims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockJWTMockRecorder) Parse(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockJWT)(nil).Parse), token)
}
