// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openclarity/apiclarity/backend/pkg/database (interfaces: APIEventsTable)

// Package database is a generated GoMock package.
package database

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	models "github.com/openclarity/apiclarity/api/server/models"
	operations "github.com/openclarity/apiclarity/api/server/restapi/operations"
	spec "github.com/openclarity/speculator/pkg/spec"
)

// MockAPIEventsTable is a mock of APIEventsTable interface.
type MockAPIEventsTable struct {
	ctrl     *gomock.Controller
	recorder *MockAPIEventsTableMockRecorder
}

// MockAPIEventsTableMockRecorder is the mock recorder for MockAPIEventsTable.
type MockAPIEventsTableMockRecorder struct {
	mock *MockAPIEventsTable
}

// NewMockAPIEventsTable creates a new mock instance.
func NewMockAPIEventsTable(ctrl *gomock.Controller) *MockAPIEventsTable {
	mock := &MockAPIEventsTable{ctrl: ctrl}
	mock.recorder = &MockAPIEventsTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIEventsTable) EXPECT() *MockAPIEventsTableMockRecorder {
	return m.recorder
}

// CreateAPIEvent mocks base method.
func (m *MockAPIEventsTable) CreateAPIEvent(arg0 *APIEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateAPIEvent", arg0)
}

// CreateAPIEvent indicates an expected call of CreateAPIEvent.
func (mr *MockAPIEventsTableMockRecorder) CreateAPIEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAPIEvent", reflect.TypeOf((*MockAPIEventsTable)(nil).CreateAPIEvent), arg0)
}

// GetAPIEvent mocks base method.
func (m *MockAPIEventsTable) GetAPIEvent(arg0 uint32) (*APIEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIEvent", arg0)
	ret0, _ := ret[0].(*APIEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAPIEvent indicates an expected call of GetAPIEvent.
func (mr *MockAPIEventsTableMockRecorder) GetAPIEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIEvent", reflect.TypeOf((*MockAPIEventsTable)(nil).GetAPIEvent), arg0)
}

// GetAPIEventProvidedSpecDiff mocks base method.
func (m *MockAPIEventsTable) GetAPIEventProvidedSpecDiff(arg0 uint32) (*APIEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIEventProvidedSpecDiff", arg0)
	ret0, _ := ret[0].(*APIEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAPIEventProvidedSpecDiff indicates an expected call of GetAPIEventProvidedSpecDiff.
func (mr *MockAPIEventsTableMockRecorder) GetAPIEventProvidedSpecDiff(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIEventProvidedSpecDiff", reflect.TypeOf((*MockAPIEventsTable)(nil).GetAPIEventProvidedSpecDiff), arg0)
}

// GetAPIEventReconstructedSpecDiff mocks base method.
func (m *MockAPIEventsTable) GetAPIEventReconstructedSpecDiff(arg0 uint32) (*APIEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIEventReconstructedSpecDiff", arg0)
	ret0, _ := ret[0].(*APIEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAPIEventReconstructedSpecDiff indicates an expected call of GetAPIEventReconstructedSpecDiff.
func (mr *MockAPIEventsTableMockRecorder) GetAPIEventReconstructedSpecDiff(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIEventReconstructedSpecDiff", reflect.TypeOf((*MockAPIEventsTable)(nil).GetAPIEventReconstructedSpecDiff), arg0)
}

// GetAPIEventsAndTotal mocks base method.
func (m *MockAPIEventsTable) GetAPIEventsAndTotal(arg0 operations.GetAPIEventsParams) ([]APIEvent, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIEventsAndTotal", arg0)
	ret0, _ := ret[0].([]APIEvent)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAPIEventsAndTotal indicates an expected call of GetAPIEventsAndTotal.
func (mr *MockAPIEventsTableMockRecorder) GetAPIEventsAndTotal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIEventsAndTotal", reflect.TypeOf((*MockAPIEventsTable)(nil).GetAPIEventsAndTotal), arg0)
}

// GetAPIEventsLatestDiffs mocks base method.
func (m *MockAPIEventsTable) GetAPIEventsLatestDiffs(arg0 int) ([]APIEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIEventsLatestDiffs", arg0)
	ret0, _ := ret[0].([]APIEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAPIEventsLatestDiffs indicates an expected call of GetAPIEventsLatestDiffs.
func (mr *MockAPIEventsTableMockRecorder) GetAPIEventsLatestDiffs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIEventsLatestDiffs", reflect.TypeOf((*MockAPIEventsTable)(nil).GetAPIEventsLatestDiffs), arg0)
}

// GetAPIEventsTotal mocks base method.
func (m *MockAPIEventsTable) GetAPIEventsTotal(arg0 operations.GetAPIEventsParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIEventsTotal", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAPIEventsTotal indicates an expected call of GetAPIEventsTotal.
func (mr *MockAPIEventsTableMockRecorder) GetAPIEventsTotal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIEventsTotal", reflect.TypeOf((*MockAPIEventsTable)(nil).GetAPIEventsTotal), arg0)
}

// GetAPIEventsWithAnnotations mocks base method.
func (m *MockAPIEventsTable) GetAPIEventsWithAnnotations(arg0 context.Context, arg1 GetAPIEventsQuery) ([]*APIEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIEventsWithAnnotations", arg0, arg1)
	ret0, _ := ret[0].([]*APIEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAPIEventsWithAnnotations indicates an expected call of GetAPIEventsWithAnnotations.
func (mr *MockAPIEventsTableMockRecorder) GetAPIEventsWithAnnotations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIEventsWithAnnotations", reflect.TypeOf((*MockAPIEventsTable)(nil).GetAPIEventsWithAnnotations), arg0, arg1)
}

// GetAPIUsages mocks base method.
func (m *MockAPIEventsTable) GetAPIUsages(arg0 operations.GetAPIUsageHitCountParams) ([]*models.HitCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIUsages", arg0)
	ret0, _ := ret[0].([]*models.HitCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAPIUsages indicates an expected call of GetAPIUsages.
func (mr *MockAPIEventsTableMockRecorder) GetAPIUsages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIUsages", reflect.TypeOf((*MockAPIEventsTable)(nil).GetAPIUsages), arg0)
}

// GetDashboardAPIUsages mocks base method.
func (m *MockAPIEventsTable) GetDashboardAPIUsages(arg0, arg1 time.Time, arg2 APIUsageType) ([]*models.APIUsage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDashboardAPIUsages", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*models.APIUsage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDashboardAPIUsages indicates an expected call of GetDashboardAPIUsages.
func (mr *MockAPIEventsTableMockRecorder) GetDashboardAPIUsages(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDashboardAPIUsages", reflect.TypeOf((*MockAPIEventsTable)(nil).GetDashboardAPIUsages), arg0, arg1, arg2)
}

// GroupByAPIInfo mocks base method.
func (m *MockAPIEventsTable) GroupByAPIInfo() ([]HostGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GroupByAPIInfo")
	ret0, _ := ret[0].([]HostGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GroupByAPIInfo indicates an expected call of GroupByAPIInfo.
func (mr *MockAPIEventsTableMockRecorder) GroupByAPIInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupByAPIInfo", reflect.TypeOf((*MockAPIEventsTable)(nil).GroupByAPIInfo))
}

// SetAPIEventsReconstructedPathID mocks base method.
func (m *MockAPIEventsTable) SetAPIEventsReconstructedPathID(arg0 []*spec.ApprovedSpecReviewPathItem, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAPIEventsReconstructedPathID", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAPIEventsReconstructedPathID indicates an expected call of SetAPIEventsReconstructedPathID.
func (mr *MockAPIEventsTableMockRecorder) SetAPIEventsReconstructedPathID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAPIEventsReconstructedPathID", reflect.TypeOf((*MockAPIEventsTable)(nil).SetAPIEventsReconstructedPathID), arg0, arg1, arg2)
}

// UpdateAPIEvent mocks base method.
func (m *MockAPIEventsTable) UpdateAPIEvent(arg0 *APIEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAPIEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAPIEvent indicates an expected call of UpdateAPIEvent.
func (mr *MockAPIEventsTableMockRecorder) UpdateAPIEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAPIEvent", reflect.TypeOf((*MockAPIEventsTable)(nil).UpdateAPIEvent), arg0)
}
