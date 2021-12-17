// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vitelabs/go-vite/consensus (interfaces: LinkedArray)

// Package consensus is a generated GoMock package.
package consensus

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"

	types "github.com/vitelabs/go-vite/v2/common/types"
	cdb "github.com/vitelabs/go-vite/v2/ledger/consensus/cdb"
)

// MockLinkedArray is a mock of LinkedArray interface
type MockLinkedArray struct {
	ctrl     *gomock.Controller
	recorder *MockLinkedArrayMockRecorder
}

// MockLinkedArrayMockRecorder is the mock recorder for MockLinkedArray
type MockLinkedArrayMockRecorder struct {
	mock *MockLinkedArray
}

// NewMockLinkedArray creates a new mock instance
func NewMockLinkedArray(ctrl *gomock.Controller) *MockLinkedArray {
	mock := &MockLinkedArray{ctrl: ctrl}
	mock.recorder = &MockLinkedArrayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLinkedArray) EXPECT() *MockLinkedArrayMockRecorder {
	return m.recorder
}

// GetByIndex mocks base method
func (m *MockLinkedArray) GetByIndex(arg0 uint64) (*cdb.Point, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIndex", arg0)
	ret0, _ := ret[0].(*cdb.Point)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIndex indicates an expected call of GetByIndex
func (mr *MockLinkedArrayMockRecorder) GetByIndex(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIndex", reflect.TypeOf((*MockLinkedArray)(nil).GetByIndex), arg0)
}

// GetByIndexWithProof mocks base method
func (m *MockLinkedArray) GetByIndexWithProof(arg0 uint64, arg1 types.Hash) (*cdb.Point, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIndexWithProof", arg0, arg1)
	ret0, _ := ret[0].(*cdb.Point)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIndexWithProof indicates an expected call of GetByIndexWithProof
func (mr *MockLinkedArrayMockRecorder) GetByIndexWithProof(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIndexWithProof", reflect.TypeOf((*MockLinkedArray)(nil).GetByIndexWithProof), arg0, arg1)
}

// Index2Time mocks base method
func (m *MockLinkedArray) Index2Time(arg0 uint64) (time.Time, time.Time) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index2Time", arg0)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(time.Time)
	return ret0, ret1
}

// Index2Time indicates an expected call of Index2Time
func (mr *MockLinkedArrayMockRecorder) Index2Time(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index2Time", reflect.TypeOf((*MockLinkedArray)(nil).Index2Time), arg0)
}

// Time2Index mocks base method
func (m *MockLinkedArray) Time2Index(arg0 time.Time) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Time2Index", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Time2Index indicates an expected call of Time2Index
func (mr *MockLinkedArrayMockRecorder) Time2Index(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Time2Index", reflect.TypeOf((*MockLinkedArray)(nil).Time2Index), arg0)
}
