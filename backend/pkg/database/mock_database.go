// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openclarity/apiclarity/backend/pkg/database (interfaces: Database)

// Package database is a generated GoMock package.
package database

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// APIEventsAnnotationsTable mocks base method.
func (m *MockDatabase) APIEventsAnnotationsTable() APIEventAnnotationTable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIEventsAnnotationsTable")
	ret0, _ := ret[0].(APIEventAnnotationTable)
	return ret0
}

// APIEventsAnnotationsTable indicates an expected call of APIEventsAnnotationsTable.
func (mr *MockDatabaseMockRecorder) APIEventsAnnotationsTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIEventsAnnotationsTable", reflect.TypeOf((*MockDatabase)(nil).APIEventsAnnotationsTable))
}

// APIEventsTable mocks base method.
func (m *MockDatabase) APIEventsTable() APIEventsTable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIEventsTable")
	ret0, _ := ret[0].(APIEventsTable)
	return ret0
}

// APIEventsTable indicates an expected call of APIEventsTable.
func (mr *MockDatabaseMockRecorder) APIEventsTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIEventsTable", reflect.TypeOf((*MockDatabase)(nil).APIEventsTable))
}

// APIGatewaysTable mocks base method.
func (m *MockDatabase) APIGatewaysTable() APIGatewaysTable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIGatewaysTable")
	ret0, _ := ret[0].(APIGatewaysTable)
	return ret0
}

// APIGatewaysTable indicates an expected call of APIGatewaysTable.
func (mr *MockDatabaseMockRecorder) APIGatewaysTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIGatewaysTable", reflect.TypeOf((*MockDatabase)(nil).APIGatewaysTable))
}

// APIInfoAnnotationsTable mocks base method.
func (m *MockDatabase) APIInfoAnnotationsTable() APIAnnotationsTable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIInfoAnnotationsTable")
	ret0, _ := ret[0].(APIAnnotationsTable)
	return ret0
}

// APIInfoAnnotationsTable indicates an expected call of APIInfoAnnotationsTable.
func (mr *MockDatabaseMockRecorder) APIInfoAnnotationsTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIInfoAnnotationsTable", reflect.TypeOf((*MockDatabase)(nil).APIInfoAnnotationsTable))
}

// APIInventoryTable mocks base method.
func (m *MockDatabase) APIInventoryTable() APIInventoryTable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIInventoryTable")
	ret0, _ := ret[0].(APIInventoryTable)
	return ret0
}

// APIInventoryTable indicates an expected call of APIInventoryTable.
func (mr *MockDatabaseMockRecorder) APIInventoryTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIInventoryTable", reflect.TypeOf((*MockDatabase)(nil).APIInventoryTable))
}

// LabelsTable mocks base method.
func (m *MockDatabase) LabelsTable() LabelsTable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LabelsTable")
	ret0, _ := ret[0].(LabelsTable)
	return ret0
}

// LabelsTable indicates an expected call of LabelsTable.
func (mr *MockDatabaseMockRecorder) LabelsTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LabelsTable", reflect.TypeOf((*MockDatabase)(nil).LabelsTable))
}

// ReviewTable mocks base method.
func (m *MockDatabase) ReviewTable() ReviewTable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReviewTable")
	ret0, _ := ret[0].(ReviewTable)
	return ret0
}

// ReviewTable indicates an expected call of ReviewTable.
func (mr *MockDatabaseMockRecorder) ReviewTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReviewTable", reflect.TypeOf((*MockDatabase)(nil).ReviewTable))
}
