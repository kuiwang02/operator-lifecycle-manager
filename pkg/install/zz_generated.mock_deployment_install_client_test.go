// Automatically generated by MockGen. DO NOT EDIT!
// Source: pkg/client/deployment_install_client.go

package install

import (
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v1beta10 "k8s.io/api/extensions/v1beta1"
	v1beta1 "k8s.io/api/rbac/v1beta1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Mock of InstallStrategyDeploymentInterface interface
type MockInstallStrategyDeploymentInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockInstallStrategyDeploymentInterfaceRecorder
}

// Recorder for MockInstallStrategyDeploymentInterface (not exported)
type _MockInstallStrategyDeploymentInterfaceRecorder struct {
	mock *MockInstallStrategyDeploymentInterface
}

func NewMockInstallStrategyDeploymentInterface(ctrl *gomock.Controller) *MockInstallStrategyDeploymentInterface {
	mock := &MockInstallStrategyDeploymentInterface{ctrl: ctrl}
	mock.recorder = &_MockInstallStrategyDeploymentInterfaceRecorder{mock}
	return mock
}

func (_m *MockInstallStrategyDeploymentInterface) EXPECT() *_MockInstallStrategyDeploymentInterfaceRecorder {
	return _m.recorder
}

func (_m *MockInstallStrategyDeploymentInterface) CreateRole(role *v1beta1.Role) (*v1beta1.Role, error) {
	ret := _m.ctrl.Call(_m, "CreateRole", role)
	ret0, _ := ret[0].(*v1beta1.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInstallStrategyDeploymentInterfaceRecorder) CreateRole(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateRole", arg0)
}

func (_m *MockInstallStrategyDeploymentInterface) CreateRoleBinding(roleBinding *v1beta1.RoleBinding) (*v1beta1.RoleBinding, error) {
	ret := _m.ctrl.Call(_m, "CreateRoleBinding", roleBinding)
	ret0, _ := ret[0].(*v1beta1.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInstallStrategyDeploymentInterfaceRecorder) CreateRoleBinding(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateRoleBinding", arg0)
}

func (_m *MockInstallStrategyDeploymentInterface) EnsureServiceAccount(serviceAccount *v1.ServiceAccount) (*v1.ServiceAccount, error) {
	ret := _m.ctrl.Call(_m, "EnsureServiceAccount", serviceAccount)
	ret0, _ := ret[0].(*v1.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInstallStrategyDeploymentInterfaceRecorder) EnsureServiceAccount(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EnsureServiceAccount", arg0)
}

func (_m *MockInstallStrategyDeploymentInterface) CreateDeployment(deployment *v1beta10.Deployment) (*v1beta10.Deployment, error) {
	ret := _m.ctrl.Call(_m, "CreateDeployment", deployment)
	ret0, _ := ret[0].(*v1beta10.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInstallStrategyDeploymentInterfaceRecorder) CreateDeployment(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateDeployment", arg0)
}

func (_m *MockInstallStrategyDeploymentInterface) CreateOrUpdateDeployment(deployment *v1beta10.Deployment) (*v1beta10.Deployment, error) {
	ret := _m.ctrl.Call(_m, "CreateOrUpdateDeployment", deployment)
	ret0, _ := ret[0].(*v1beta10.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInstallStrategyDeploymentInterfaceRecorder) CreateOrUpdateDeployment(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateOrUpdateDeployment", arg0)
}

func (_m *MockInstallStrategyDeploymentInterface) DeleteDeployment(name string) error {
	ret := _m.ctrl.Call(_m, "DeleteDeployment", name)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockInstallStrategyDeploymentInterfaceRecorder) DeleteDeployment(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteDeployment", arg0)
}

func (_m *MockInstallStrategyDeploymentInterface) GetServiceAccountByName(serviceAccountName string) (*v1.ServiceAccount, error) {
	ret := _m.ctrl.Call(_m, "GetServiceAccountByName", serviceAccountName)
	ret0, _ := ret[0].(*v1.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInstallStrategyDeploymentInterfaceRecorder) GetServiceAccountByName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetServiceAccountByName", arg0)
}

func (_m *MockInstallStrategyDeploymentInterface) GetOwnedDeployments(owner v10.ObjectMeta) (*v1beta10.DeploymentList, error) {
	ret := _m.ctrl.Call(_m, "GetOwnedDeployments", owner)
	ret0, _ := ret[0].(*v1beta10.DeploymentList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInstallStrategyDeploymentInterfaceRecorder) GetOwnedDeployments(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOwnedDeployments", arg0)
}
