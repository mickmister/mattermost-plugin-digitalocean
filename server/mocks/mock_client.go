// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/phillipahereza/mattermost-plugin-digitalocean/server/client (interfaces: DigitalOceanService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
	client "github.com/phillipahereza/mattermost-plugin-digitalocean/server/client"
	reflect "reflect"
)

// MockDigitalOceanService is a mock of DigitalOceanService interface
type MockDigitalOceanService struct {
	ctrl     *gomock.Controller
	recorder *MockDigitalOceanServiceMockRecorder
}

// MockDigitalOceanServiceMockRecorder is the mock recorder for MockDigitalOceanService
type MockDigitalOceanServiceMockRecorder struct {
	mock *MockDigitalOceanService
}

// NewMockDigitalOceanService creates a new mock instance
func NewMockDigitalOceanService(ctrl *gomock.Controller) *MockDigitalOceanService {
	mock := &MockDigitalOceanService{ctrl: ctrl}
	mock.recorder = &MockDigitalOceanServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDigitalOceanService) EXPECT() *MockDigitalOceanServiceMockRecorder {
	return m.recorder
}

// CreateDatabaseUser mocks base method
func (m *MockDigitalOceanService) CreateDatabaseUser(arg0 context.Context, arg1 string, arg2 *godo.DatabaseCreateUserRequest) (*godo.DatabaseUser, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDatabaseUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(*godo.DatabaseUser)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateDatabaseUser indicates an expected call of CreateDatabaseUser
func (mr *MockDigitalOceanServiceMockRecorder) CreateDatabaseUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDatabaseUser", reflect.TypeOf((*MockDigitalOceanService)(nil).CreateDatabaseUser), arg0, arg1, arg2)
}

// CreateDroplet mocks base method
func (m *MockDigitalOceanService) CreateDroplet(arg0 context.Context, arg1 *godo.DropletCreateRequest) (*godo.Droplet, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDroplet", arg0, arg1)
	ret0, _ := ret[0].(*godo.Droplet)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateDroplet indicates an expected call of CreateDroplet
func (mr *MockDigitalOceanServiceMockRecorder) CreateDroplet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDroplet", reflect.TypeOf((*MockDigitalOceanService)(nil).CreateDroplet), arg0, arg1)
}

// CreateSSHKey mocks base method
func (m *MockDigitalOceanService) CreateSSHKey(arg0 context.Context, arg1 *godo.KeyCreateRequest) (*godo.Key, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSSHKey", arg0, arg1)
	ret0, _ := ret[0].(*godo.Key)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateSSHKey indicates an expected call of CreateSSHKey
func (mr *MockDigitalOceanServiceMockRecorder) CreateSSHKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSSHKey", reflect.TypeOf((*MockDigitalOceanService)(nil).CreateSSHKey), arg0, arg1)
}

// DeleteDatabaseClusterUser mocks base method
func (m *MockDigitalOceanService) DeleteDatabaseClusterUser(arg0 context.Context, arg1, arg2 string) (*godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDatabaseClusterUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(*godo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDatabaseClusterUser indicates an expected call of DeleteDatabaseClusterUser
func (mr *MockDigitalOceanServiceMockRecorder) DeleteDatabaseClusterUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDatabaseClusterUser", reflect.TypeOf((*MockDigitalOceanService)(nil).DeleteDatabaseClusterUser), arg0, arg1, arg2)
}

// DeleteSSHKeyByID mocks base method
func (m *MockDigitalOceanService) DeleteSSHKeyByID(arg0 context.Context, arg1 int) (*godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSSHKeyByID", arg0, arg1)
	ret0, _ := ret[0].(*godo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSSHKeyByID indicates an expected call of DeleteSSHKeyByID
func (mr *MockDigitalOceanServiceMockRecorder) DeleteSSHKeyByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSSHKeyByID", reflect.TypeOf((*MockDigitalOceanService)(nil).DeleteSSHKeyByID), arg0, arg1)
}

// GetKubeConfig mocks base method
func (m *MockDigitalOceanService) GetKubeConfig(arg0 context.Context, arg1 string) (*godo.KubernetesClusterConfig, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKubeConfig", arg0, arg1)
	ret0, _ := ret[0].(*godo.KubernetesClusterConfig)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetKubeConfig indicates an expected call of GetKubeConfig
func (mr *MockDigitalOceanServiceMockRecorder) GetKubeConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKubeConfig", reflect.TypeOf((*MockDigitalOceanService)(nil).GetKubeConfig), arg0, arg1)
}

// GetKubernetesClusterUpgrades mocks base method
func (m *MockDigitalOceanService) GetKubernetesClusterUpgrades(arg0 context.Context, arg1 string) ([]*godo.KubernetesVersion, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKubernetesClusterUpgrades", arg0, arg1)
	ret0, _ := ret[0].([]*godo.KubernetesVersion)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetKubernetesClusterUpgrades indicates an expected call of GetKubernetesClusterUpgrades
func (mr *MockDigitalOceanServiceMockRecorder) GetKubernetesClusterUpgrades(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKubernetesClusterUpgrades", reflect.TypeOf((*MockDigitalOceanService)(nil).GetKubernetesClusterUpgrades), arg0, arg1)
}

// GetSSHKeyByID mocks base method
func (m *MockDigitalOceanService) GetSSHKeyByID(arg0 context.Context, arg1 int) (*godo.Key, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSSHKeyByID", arg0, arg1)
	ret0, _ := ret[0].(*godo.Key)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSSHKeyByID indicates an expected call of GetSSHKeyByID
func (mr *MockDigitalOceanServiceMockRecorder) GetSSHKeyByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSSHKeyByID", reflect.TypeOf((*MockDigitalOceanService)(nil).GetSSHKeyByID), arg0, arg1)
}

// ListDatabaseClusterBackups mocks base method
func (m *MockDigitalOceanService) ListDatabaseClusterBackups(arg0 context.Context, arg1 string, arg2 *godo.ListOptions) ([]godo.DatabaseBackup, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDatabaseClusterBackups", arg0, arg1, arg2)
	ret0, _ := ret[0].([]godo.DatabaseBackup)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListDatabaseClusterBackups indicates an expected call of ListDatabaseClusterBackups
func (mr *MockDigitalOceanServiceMockRecorder) ListDatabaseClusterBackups(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDatabaseClusterBackups", reflect.TypeOf((*MockDigitalOceanService)(nil).ListDatabaseClusterBackups), arg0, arg1, arg2)
}

// ListDatabaseClusterUsers mocks base method
func (m *MockDigitalOceanService) ListDatabaseClusterUsers(arg0 context.Context, arg1 string, arg2 *godo.ListOptions) ([]godo.DatabaseUser, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDatabaseClusterUsers", arg0, arg1, arg2)
	ret0, _ := ret[0].([]godo.DatabaseUser)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListDatabaseClusterUsers indicates an expected call of ListDatabaseClusterUsers
func (mr *MockDigitalOceanServiceMockRecorder) ListDatabaseClusterUsers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDatabaseClusterUsers", reflect.TypeOf((*MockDigitalOceanService)(nil).ListDatabaseClusterUsers), arg0, arg1, arg2)
}

// ListDatabaseClusters mocks base method
func (m *MockDigitalOceanService) ListDatabaseClusters(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Database, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDatabaseClusters", arg0, arg1)
	ret0, _ := ret[0].([]godo.Database)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListDatabaseClusters indicates an expected call of ListDatabaseClusters
func (mr *MockDigitalOceanServiceMockRecorder) ListDatabaseClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDatabaseClusters", reflect.TypeOf((*MockDigitalOceanService)(nil).ListDatabaseClusters), arg0, arg1)
}

// ListDatabasesInCluster mocks base method
func (m *MockDigitalOceanService) ListDatabasesInCluster(arg0 context.Context, arg1 string, arg2 *godo.ListOptions) ([]godo.DatabaseDB, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDatabasesInCluster", arg0, arg1, arg2)
	ret0, _ := ret[0].([]godo.DatabaseDB)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListDatabasesInCluster indicates an expected call of ListDatabasesInCluster
func (mr *MockDigitalOceanServiceMockRecorder) ListDatabasesInCluster(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDatabasesInCluster", reflect.TypeOf((*MockDigitalOceanService)(nil).ListDatabasesInCluster), arg0, arg1, arg2)
}

// ListDomains mocks base method
func (m *MockDigitalOceanService) ListDomains(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Domain, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDomains", arg0, arg1)
	ret0, _ := ret[0].([]godo.Domain)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListDomains indicates an expected call of ListDomains
func (mr *MockDigitalOceanServiceMockRecorder) ListDomains(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDomains", reflect.TypeOf((*MockDigitalOceanService)(nil).ListDomains), arg0, arg1)
}

// ListDroplets mocks base method
func (m *MockDigitalOceanService) ListDroplets(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Droplet, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDroplets", arg0, arg1)
	ret0, _ := ret[0].([]godo.Droplet)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListDroplets indicates an expected call of ListDroplets
func (mr *MockDigitalOceanServiceMockRecorder) ListDroplets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDroplets", reflect.TypeOf((*MockDigitalOceanService)(nil).ListDroplets), arg0, arg1)
}

// ListImages mocks base method
func (m *MockDigitalOceanService) ListImages(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Image, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages", arg0, arg1)
	ret0, _ := ret[0].([]godo.Image)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListImages indicates an expected call of ListImages
func (mr *MockDigitalOceanServiceMockRecorder) ListImages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockDigitalOceanService)(nil).ListImages), arg0, arg1)
}

// ListKubernetesClusterNodePools mocks base method
func (m *MockDigitalOceanService) ListKubernetesClusterNodePools(arg0 context.Context, arg1 string, arg2 *godo.ListOptions) ([]*godo.KubernetesNodePool, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKubernetesClusterNodePools", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*godo.KubernetesNodePool)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListKubernetesClusterNodePools indicates an expected call of ListKubernetesClusterNodePools
func (mr *MockDigitalOceanServiceMockRecorder) ListKubernetesClusterNodePools(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKubernetesClusterNodePools", reflect.TypeOf((*MockDigitalOceanService)(nil).ListKubernetesClusterNodePools), arg0, arg1, arg2)
}

// ListKubernetesClusterNodes mocks base method
func (m *MockDigitalOceanService) ListKubernetesClusterNodes(arg0 context.Context, arg1 string, arg2 *godo.ListOptions) ([]*client.K8sNode, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKubernetesClusterNodes", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*client.K8sNode)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListKubernetesClusterNodes indicates an expected call of ListKubernetesClusterNodes
func (mr *MockDigitalOceanServiceMockRecorder) ListKubernetesClusterNodes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKubernetesClusterNodes", reflect.TypeOf((*MockDigitalOceanService)(nil).ListKubernetesClusterNodes), arg0, arg1, arg2)
}

// ListKubernetesClusters mocks base method
func (m *MockDigitalOceanService) ListKubernetesClusters(arg0 context.Context, arg1 *godo.ListOptions) ([]*godo.KubernetesCluster, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKubernetesClusters", arg0, arg1)
	ret0, _ := ret[0].([]*godo.KubernetesCluster)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListKubernetesClusters indicates an expected call of ListKubernetesClusters
func (mr *MockDigitalOceanServiceMockRecorder) ListKubernetesClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKubernetesClusters", reflect.TypeOf((*MockDigitalOceanService)(nil).ListKubernetesClusters), arg0, arg1)
}

// ListRegions mocks base method
func (m *MockDigitalOceanService) ListRegions(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Region, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRegions", arg0, arg1)
	ret0, _ := ret[0].([]godo.Region)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListRegions indicates an expected call of ListRegions
func (mr *MockDigitalOceanServiceMockRecorder) ListRegions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRegions", reflect.TypeOf((*MockDigitalOceanService)(nil).ListRegions), arg0, arg1)
}

// ListSSHKeys mocks base method
func (m *MockDigitalOceanService) ListSSHKeys(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Key, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSSHKeys", arg0, arg1)
	ret0, _ := ret[0].([]godo.Key)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListSSHKeys indicates an expected call of ListSSHKeys
func (mr *MockDigitalOceanServiceMockRecorder) ListSSHKeys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSSHKeys", reflect.TypeOf((*MockDigitalOceanService)(nil).ListSSHKeys), arg0, arg1)
}

// ListSizes mocks base method
func (m *MockDigitalOceanService) ListSizes(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Size, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSizes", arg0, arg1)
	ret0, _ := ret[0].([]godo.Size)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListSizes indicates an expected call of ListSizes
func (mr *MockDigitalOceanServiceMockRecorder) ListSizes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSizes", reflect.TypeOf((*MockDigitalOceanService)(nil).ListSizes), arg0, arg1)
}

// PowerCycleDroplet mocks base method
func (m *MockDigitalOceanService) PowerCycleDroplet(arg0 context.Context, arg1 int) (*godo.Action, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerCycleDroplet", arg0, arg1)
	ret0, _ := ret[0].(*godo.Action)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PowerCycleDroplet indicates an expected call of PowerCycleDroplet
func (mr *MockDigitalOceanServiceMockRecorder) PowerCycleDroplet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerCycleDroplet", reflect.TypeOf((*MockDigitalOceanService)(nil).PowerCycleDroplet), arg0, arg1)
}

// RebootDroplet mocks base method
func (m *MockDigitalOceanService) RebootDroplet(arg0 context.Context, arg1 int) (*godo.Action, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RebootDroplet", arg0, arg1)
	ret0, _ := ret[0].(*godo.Action)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RebootDroplet indicates an expected call of RebootDroplet
func (mr *MockDigitalOceanServiceMockRecorder) RebootDroplet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RebootDroplet", reflect.TypeOf((*MockDigitalOceanService)(nil).RebootDroplet), arg0, arg1)
}

// RenameDroplet mocks base method
func (m *MockDigitalOceanService) RenameDroplet(arg0 context.Context, arg1 int, arg2 string) (*godo.Action, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameDroplet", arg0, arg1, arg2)
	ret0, _ := ret[0].(*godo.Action)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RenameDroplet indicates an expected call of RenameDroplet
func (mr *MockDigitalOceanServiceMockRecorder) RenameDroplet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameDroplet", reflect.TypeOf((*MockDigitalOceanService)(nil).RenameDroplet), arg0, arg1, arg2)
}

// ShutdownDroplet mocks base method
func (m *MockDigitalOceanService) ShutdownDroplet(arg0 context.Context, arg1 int) (*godo.Action, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShutdownDroplet", arg0, arg1)
	ret0, _ := ret[0].(*godo.Action)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ShutdownDroplet indicates an expected call of ShutdownDroplet
func (mr *MockDigitalOceanServiceMockRecorder) ShutdownDroplet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutdownDroplet", reflect.TypeOf((*MockDigitalOceanService)(nil).ShutdownDroplet), arg0, arg1)
}

// UpgradeKubernetesCluster mocks base method
func (m *MockDigitalOceanService) UpgradeKubernetesCluster(arg0 context.Context, arg1 string, arg2 *godo.KubernetesClusterUpgradeRequest) (*godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeKubernetesCluster", arg0, arg1, arg2)
	ret0, _ := ret[0].(*godo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeKubernetesCluster indicates an expected call of UpgradeKubernetesCluster
func (mr *MockDigitalOceanServiceMockRecorder) UpgradeKubernetesCluster(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeKubernetesCluster", reflect.TypeOf((*MockDigitalOceanService)(nil).UpgradeKubernetesCluster), arg0, arg1, arg2)
}
