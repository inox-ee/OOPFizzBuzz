// Code generated by MockGen. DO NOT EDIT.
// Source: number_converter.go
//
// Generated by this command:
//
//	mockgen -source number_converter.go -destination number_converter_mock.go -package core
//

// Package core is a generated GoMock package.
package core

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockN2SReplacer is a mock of N2SReplacer interface.
type MockN2SReplacer struct {
	ctrl     *gomock.Controller
	recorder *MockN2SReplacerMockRecorder
}

// MockN2SReplacerMockRecorder is the mock recorder for MockN2SReplacer.
type MockN2SReplacerMockRecorder struct {
	mock *MockN2SReplacer
}

// NewMockN2SReplacer creates a new mock instance.
func NewMockN2SReplacer(ctrl *gomock.Controller) *MockN2SReplacer {
	mock := &MockN2SReplacer{ctrl: ctrl}
	mock.recorder = &MockN2SReplacerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockN2SReplacer) EXPECT() *MockN2SReplacerMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockN2SReplacer) Apply(carry string, number int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", carry, number)
	ret0, _ := ret[0].(string)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockN2SReplacerMockRecorder) Apply(carry, number any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockN2SReplacer)(nil).Apply), carry, number)
}

// Match mocks base method.
func (m *MockN2SReplacer) Match(carry string, number int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Match", carry, number)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Match indicates an expected call of Match.
func (mr *MockN2SReplacerMockRecorder) Match(carry, number any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Match", reflect.TypeOf((*MockN2SReplacer)(nil).Match), carry, number)
}
