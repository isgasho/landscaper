// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/landscaper/pkg/landscaper/registry (interfaces: Registry)

// Package mock_registry is a generated GoMock package.
package mock_registry

import (
	context "context"
	reflect "reflect"

	v2 "github.com/gardener/component-spec/bindings-go/apis/v2"
	gomock "github.com/golang/mock/gomock"
	afero "github.com/spf13/afero"

	v1alpha1 "github.com/gardener/landscaper/pkg/apis/core/v1alpha1"
)

// MockRegistry is a mock of Registry interface
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// GetBlueprint mocks base method
func (m *MockRegistry) GetBlueprint(arg0 context.Context, arg1 v2.Resource) (*v1alpha1.Blueprint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlueprint", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.Blueprint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlueprint indicates an expected call of GetBlueprint
func (mr *MockRegistryMockRecorder) GetBlueprint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlueprint", reflect.TypeOf((*MockRegistry)(nil).GetBlueprint), arg0, arg1)
}

// GetContent mocks base method
func (m *MockRegistry) GetContent(ctx context.Context, ref v2.Resource, fs afero.Fs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContent", arg0, arg1)
	ret0, _ := ret[0].(afero.Fs)
	ret1, _ := ret[1].(error)
	return ret1
}

// GetContent indicates an expected call of GetContent
func (mr *MockRegistryMockRecorder) GetContent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContent", reflect.TypeOf((*MockRegistry)(nil).GetContent), arg0, arg1)
}
