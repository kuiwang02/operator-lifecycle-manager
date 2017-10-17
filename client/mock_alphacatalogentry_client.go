// Automatically generated by MockGen. DO NOT EDIT!
// Source: client/alphacatalogentry_client.go

package client

import (
	v1alpha1 "github.com/coreos-inc/alm/apis/alphacatalogentry/v1alpha1"
	gomock "github.com/golang/mock/gomock"
)

// Mock of AlphaCatalogEntryInterface interface
type MockAlphaCatalogEntryInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockAlphaCatalogEntryInterfaceRecorder
}

// Recorder for MockAlphaCatalogEntryInterface (not exported)
type _MockAlphaCatalogEntryInterfaceRecorder struct {
	mock *MockAlphaCatalogEntryInterface
}

func NewMockAlphaCatalogEntryInterface(ctrl *gomock.Controller) *MockAlphaCatalogEntryInterface {
	mock := &MockAlphaCatalogEntryInterface{ctrl: ctrl}
	mock.recorder = &_MockAlphaCatalogEntryInterfaceRecorder{mock}
	return mock
}

func (_m *MockAlphaCatalogEntryInterface) EXPECT() *_MockAlphaCatalogEntryInterfaceRecorder {
	return _m.recorder
}

func (_m *MockAlphaCatalogEntryInterface) UpdateEntry(csv *v1alpha1.AlphaCatalogEntry) (*v1alpha1.AlphaCatalogEntry, error) {
	ret := _m.ctrl.Call(_m, "UpdateEntry", csv)
	ret0, _ := ret[0].(*v1alpha1.AlphaCatalogEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAlphaCatalogEntryInterfaceRecorder) UpdateEntry(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateEntry", arg0)
}
