// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/AmazingTalker/bevis-chang/pkg/dao (interfaces: RecordDAO)

// Package daomock is a generated GoMock package.
package daomock

import (
	context "context"
	reflect "reflect"

	dao "github.com/AmazingTalker/bevis-chang/pkg/dao"
	daokit "github.com/AmazingTalker/go-rpc-kit/daokit"
	gomock "github.com/golang/mock/gomock"
)

// MockRecordDAO is a mock of RecordDAO interface.
type MockRecordDAO struct {
	ctrl     *gomock.Controller
	recorder *MockRecordDAOMockRecorder
}

// MockRecordDAOMockRecorder is the mock recorder for MockRecordDAO.
type MockRecordDAOMockRecorder struct {
	mock *MockRecordDAO
}

// NewMockRecordDAO creates a new mock instance.
func NewMockRecordDAO(ctrl *gomock.Controller) *MockRecordDAO {
	mock := &MockRecordDAO{ctrl: ctrl}
	mock.recorder = &MockRecordDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRecordDAO) EXPECT() *MockRecordDAOMockRecorder {
	return m.recorder
}

// CreateRecord mocks base method.
func (m *MockRecordDAO) CreateRecord(arg0 context.Context, arg1 *dao.Record, arg2 ...daokit.Enrich) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateRecord", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRecord indicates an expected call of CreateRecord.
func (mr *MockRecordDAOMockRecorder) CreateRecord(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecord", reflect.TypeOf((*MockRecordDAO)(nil).CreateRecord), varargs...)
}

// GetRecord mocks base method.
func (m *MockRecordDAO) GetRecord(arg0 context.Context, arg1 string) (*dao.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecord", arg0, arg1)
	ret0, _ := ret[0].(*dao.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecord indicates an expected call of GetRecord.
func (mr *MockRecordDAOMockRecorder) GetRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecord", reflect.TypeOf((*MockRecordDAO)(nil).GetRecord), arg0, arg1)
}

// ListRecords mocks base method.
func (m *MockRecordDAO) ListRecords(arg0 context.Context, arg1 dao.ListRecordsOpt) ([]dao.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRecords", arg0, arg1)
	ret0, _ := ret[0].([]dao.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRecords indicates an expected call of ListRecords.
func (mr *MockRecordDAOMockRecorder) ListRecords(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRecords", reflect.TypeOf((*MockRecordDAO)(nil).ListRecords), arg0, arg1)
}
