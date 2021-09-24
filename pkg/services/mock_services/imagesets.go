// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/imagesets.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockImageSetsServiceInterface is a mock of ImageSetsServiceInterface interface
type MockImageSetsServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockImageSetsServiceInterfaceMockRecorder
}

// MockImageSetsServiceInterfaceMockRecorder is the mock recorder for MockImageSetsServiceInterface
type MockImageSetsServiceInterfaceMockRecorder struct {
	mock *MockImageSetsServiceInterface
}

// NewMockImageSetsServiceInterface creates a new mock instance
func NewMockImageSetsServiceInterface(ctrl *gomock.Controller) *MockImageSetsServiceInterface {
	mock := &MockImageSetsServiceInterface{ctrl: ctrl}
	mock.recorder = &MockImageSetsServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImageSetsServiceInterface) EXPECT() *MockImageSetsServiceInterfaceMockRecorder {
	return m.recorder
}

// ListAllImageSets mocks base method
func (m *MockImageSetsServiceInterface) ListAllImageSets(w http.ResponseWriter, r *http.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllImageSets", w, r)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListAllImageSets indicates an expected call of ListAllImageSets
func (mr *MockImageSetsServiceInterfaceMockRecorder) ListAllImageSets(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllImageSets", reflect.TypeOf((*MockImageSetsServiceInterface)(nil).ListAllImageSets), w, r)
}

// GetImageSetsByID mocks base method
func (m *MockImageSetsServiceInterface) GetImageSetsByID(w http.ResponseWriter, r *http.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageSetsByID", w, r)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetImageSetsByID indicates an expected call of GetImageSetsByID
func (mr *MockImageSetsServiceInterfaceMockRecorder) GetImageSetsByID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageSetsByID", reflect.TypeOf((*MockImageSetsServiceInterface)(nil).GetImageSetsByID), w, r)
}
