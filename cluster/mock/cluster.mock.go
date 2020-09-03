// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/cluster (interfaces: Cluster)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	types "github.com/libopenstorage/gossip/types"
	api "github.com/libopenstorage/openstorage/api"
	cluster "github.com/libopenstorage/openstorage/cluster"
	osdconfig "github.com/libopenstorage/openstorage/osdconfig"
	clusterdomain "github.com/libopenstorage/openstorage/pkg/clusterdomain"
	schedpolicy "github.com/libopenstorage/openstorage/schedpolicy"
)

// MockCluster is a mock of Cluster interface.
type MockCluster struct {
	ctrl     *gomock.Controller
	recorder *MockClusterMockRecorder
}

// MockClusterMockRecorder is the mock recorder for MockCluster.
type MockClusterMockRecorder struct {
	mock *MockCluster
}

// NewMockCluster creates a new mock instance.
func NewMockCluster(ctrl *gomock.Controller) *MockCluster {
	mock := &MockCluster{ctrl: ctrl}
	mock.recorder = &MockClusterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCluster) EXPECT() *MockClusterMockRecorder {
	return m.recorder
}

// AddEventListener mocks base method.
func (m *MockCluster) AddEventListener(arg0 cluster.ClusterListener) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEventListener", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventListener indicates an expected call of AddEventListener.
func (mr *MockClusterMockRecorder) AddEventListener(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventListener", reflect.TypeOf((*MockCluster)(nil).AddEventListener), arg0)
}

// ClusterNotifyClusterDomainsUpdate mocks base method.
func (m *MockCluster) ClusterNotifyClusterDomainsUpdate(arg0 types.ClusterDomainsActiveMap) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterNotifyClusterDomainsUpdate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClusterNotifyClusterDomainsUpdate indicates an expected call of ClusterNotifyClusterDomainsUpdate.
func (mr *MockClusterMockRecorder) ClusterNotifyClusterDomainsUpdate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterNotifyClusterDomainsUpdate", reflect.TypeOf((*MockCluster)(nil).ClusterNotifyClusterDomainsUpdate), arg0)
}

// ClusterNotifyNodeDown mocks base method.
func (m *MockCluster) ClusterNotifyNodeDown(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterNotifyNodeDown", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClusterNotifyNodeDown indicates an expected call of ClusterNotifyNodeDown.
func (mr *MockClusterMockRecorder) ClusterNotifyNodeDown(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterNotifyNodeDown", reflect.TypeOf((*MockCluster)(nil).ClusterNotifyNodeDown), arg0)
}

// CreatePair mocks base method.
func (m *MockCluster) CreatePair(arg0 *api.ClusterPairCreateRequest) (*api.ClusterPairCreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePair", arg0)
	ret0, _ := ret[0].(*api.ClusterPairCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePair indicates an expected call of CreatePair.
func (mr *MockClusterMockRecorder) CreatePair(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePair", reflect.TypeOf((*MockCluster)(nil).CreatePair), arg0)
}

// DeleteDomain mocks base method.
func (m *MockCluster) DeleteDomain(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDomain", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDomain indicates an expected call of DeleteDomain.
func (mr *MockClusterMockRecorder) DeleteDomain(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDomain", reflect.TypeOf((*MockCluster)(nil).DeleteDomain), arg0)
}

// DeleteNodeConf mocks base method.
func (m *MockCluster) DeleteNodeConf(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNodeConf", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNodeConf indicates an expected call of DeleteNodeConf.
func (mr *MockClusterMockRecorder) DeleteNodeConf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNodeConf", reflect.TypeOf((*MockCluster)(nil).DeleteNodeConf), arg0)
}

// DeletePair mocks base method.
func (m *MockCluster) DeletePair(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePair", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePair indicates an expected call of DeletePair.
func (mr *MockClusterMockRecorder) DeletePair(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePair", reflect.TypeOf((*MockCluster)(nil).DeletePair), arg0)
}

// DisableUpdates mocks base method.
func (m *MockCluster) DisableUpdates() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableUpdates")
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableUpdates indicates an expected call of DisableUpdates.
func (mr *MockClusterMockRecorder) DisableUpdates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableUpdates", reflect.TypeOf((*MockCluster)(nil).DisableUpdates))
}

// EnableUpdates mocks base method.
func (m *MockCluster) EnableUpdates() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableUpdates")
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableUpdates indicates an expected call of EnableUpdates.
func (mr *MockClusterMockRecorder) EnableUpdates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableUpdates", reflect.TypeOf((*MockCluster)(nil).EnableUpdates))
}

// Enumerate mocks base method.
func (m *MockCluster) Enumerate() (api.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enumerate")
	ret0, _ := ret[0].(api.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enumerate indicates an expected call of Enumerate.
func (mr *MockClusterMockRecorder) Enumerate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enumerate", reflect.TypeOf((*MockCluster)(nil).Enumerate))
}

// EnumerateAlerts mocks base method.
func (m *MockCluster) EnumerateAlerts(arg0, arg1 time.Time, arg2 api.ResourceType) (*api.Alerts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnumerateAlerts", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.Alerts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnumerateAlerts indicates an expected call of EnumerateAlerts.
func (mr *MockClusterMockRecorder) EnumerateAlerts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumerateAlerts", reflect.TypeOf((*MockCluster)(nil).EnumerateAlerts), arg0, arg1, arg2)
}

// EnumerateDomains mocks base method.
func (m *MockCluster) EnumerateDomains() ([]*clusterdomain.ClusterDomainInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnumerateDomains")
	ret0, _ := ret[0].([]*clusterdomain.ClusterDomainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnumerateDomains indicates an expected call of EnumerateDomains.
func (mr *MockClusterMockRecorder) EnumerateDomains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumerateDomains", reflect.TypeOf((*MockCluster)(nil).EnumerateDomains))
}

// EnumerateJobs mocks base method.
func (m *MockCluster) EnumerateJobs(arg0 context.Context, arg1 *api.SdkEnumerateJobsRequest) (*api.SdkEnumerateJobsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnumerateJobs", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkEnumerateJobsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnumerateJobs indicates an expected call of EnumerateJobs.
func (mr *MockClusterMockRecorder) EnumerateJobs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumerateJobs", reflect.TypeOf((*MockCluster)(nil).EnumerateJobs), arg0, arg1)
}

// EnumerateNodeConf mocks base method.
func (m *MockCluster) EnumerateNodeConf() (*osdconfig.NodesConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnumerateNodeConf")
	ret0, _ := ret[0].(*osdconfig.NodesConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnumerateNodeConf indicates an expected call of EnumerateNodeConf.
func (mr *MockClusterMockRecorder) EnumerateNodeConf() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumerateNodeConf", reflect.TypeOf((*MockCluster)(nil).EnumerateNodeConf))
}

// EnumeratePairs mocks base method.
func (m *MockCluster) EnumeratePairs() (*api.ClusterPairsEnumerateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnumeratePairs")
	ret0, _ := ret[0].(*api.ClusterPairsEnumerateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnumeratePairs indicates an expected call of EnumeratePairs.
func (mr *MockClusterMockRecorder) EnumeratePairs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumeratePairs", reflect.TypeOf((*MockCluster)(nil).EnumeratePairs))
}

// EnumerateRebalanceJobs mocks base method.
func (m *MockCluster) EnumerateRebalanceJobs(arg0 context.Context, arg1 *api.SdkEnumerateRebalanceJobsRequest) (*api.SdkEnumerateRebalanceJobsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnumerateRebalanceJobs", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkEnumerateRebalanceJobsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnumerateRebalanceJobs indicates an expected call of EnumerateRebalanceJobs.
func (mr *MockClusterMockRecorder) EnumerateRebalanceJobs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumerateRebalanceJobs", reflect.TypeOf((*MockCluster)(nil).EnumerateRebalanceJobs), arg0, arg1)
}

// EraseAlert mocks base method.
func (m *MockCluster) EraseAlert(arg0 api.ResourceType, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EraseAlert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EraseAlert indicates an expected call of EraseAlert.
func (mr *MockClusterMockRecorder) EraseAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EraseAlert", reflect.TypeOf((*MockCluster)(nil).EraseAlert), arg0, arg1)
}

// GetClusterConf mocks base method.
func (m *MockCluster) GetClusterConf() (*osdconfig.ClusterConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterConf")
	ret0, _ := ret[0].(*osdconfig.ClusterConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterConf indicates an expected call of GetClusterConf.
func (mr *MockClusterMockRecorder) GetClusterConf() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterConf", reflect.TypeOf((*MockCluster)(nil).GetClusterConf))
}

// GetData mocks base method.
func (m *MockCluster) GetData() (map[string]*api.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData")
	ret0, _ := ret[0].(map[string]*api.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetData indicates an expected call of GetData.
func (mr *MockClusterMockRecorder) GetData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockCluster)(nil).GetData))
}

// GetGossipState mocks base method.
func (m *MockCluster) GetGossipState() *cluster.ClusterState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGossipState")
	ret0, _ := ret[0].(*cluster.ClusterState)
	return ret0
}

// GetGossipState indicates an expected call of GetGossipState.
func (mr *MockClusterMockRecorder) GetGossipState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGossipState", reflect.TypeOf((*MockCluster)(nil).GetGossipState))
}

// GetJobStatus mocks base method.
func (m *MockCluster) GetJobStatus(arg0 context.Context, arg1 *api.SdkGetJobStatusRequest) (*api.SdkGetJobStatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobStatus", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkGetJobStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobStatus indicates an expected call of GetJobStatus.
func (mr *MockClusterMockRecorder) GetJobStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobStatus", reflect.TypeOf((*MockCluster)(nil).GetJobStatus), arg0, arg1)
}

// GetNodeConf mocks base method.
func (m *MockCluster) GetNodeConf(arg0 string) (*osdconfig.NodeConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeConf", arg0)
	ret0, _ := ret[0].(*osdconfig.NodeConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeConf indicates an expected call of GetNodeConf.
func (mr *MockClusterMockRecorder) GetNodeConf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeConf", reflect.TypeOf((*MockCluster)(nil).GetNodeConf), arg0)
}

// GetNodeIdFromIp mocks base method.
func (m *MockCluster) GetNodeIdFromIp(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeIdFromIp", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeIdFromIp indicates an expected call of GetNodeIdFromIp.
func (mr *MockClusterMockRecorder) GetNodeIdFromIp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeIdFromIp", reflect.TypeOf((*MockCluster)(nil).GetNodeIdFromIp), arg0)
}

// GetPair mocks base method.
func (m *MockCluster) GetPair(arg0 string) (*api.ClusterPairGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPair", arg0)
	ret0, _ := ret[0].(*api.ClusterPairGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPair indicates an expected call of GetPair.
func (mr *MockClusterMockRecorder) GetPair(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPair", reflect.TypeOf((*MockCluster)(nil).GetPair), arg0)
}

// GetPairToken mocks base method.
func (m *MockCluster) GetPairToken(arg0 bool) (*api.ClusterPairTokenGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPairToken", arg0)
	ret0, _ := ret[0].(*api.ClusterPairTokenGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPairToken indicates an expected call of GetPairToken.
func (mr *MockClusterMockRecorder) GetPairToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPairToken", reflect.TypeOf((*MockCluster)(nil).GetPairToken), arg0)
}

// GetRebalanceJobStatus mocks base method.
func (m *MockCluster) GetRebalanceJobStatus(arg0 context.Context, arg1 *api.SdkGetRebalanceJobStatusRequest) (*api.SdkGetRebalanceJobStatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRebalanceJobStatus", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkGetRebalanceJobStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRebalanceJobStatus indicates an expected call of GetRebalanceJobStatus.
func (mr *MockClusterMockRecorder) GetRebalanceJobStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRebalanceJobStatus", reflect.TypeOf((*MockCluster)(nil).GetRebalanceJobStatus), arg0, arg1)
}

// GetSelfDomain mocks base method.
func (m *MockCluster) GetSelfDomain() (*clusterdomain.ClusterDomainInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelfDomain")
	ret0, _ := ret[0].(*clusterdomain.ClusterDomainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSelfDomain indicates an expected call of GetSelfDomain.
func (mr *MockClusterMockRecorder) GetSelfDomain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelfDomain", reflect.TypeOf((*MockCluster)(nil).GetSelfDomain))
}

// Inspect mocks base method.
func (m *MockCluster) Inspect(arg0 string) (api.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Inspect", arg0)
	ret0, _ := ret[0].(api.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inspect indicates an expected call of Inspect.
func (mr *MockClusterMockRecorder) Inspect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inspect", reflect.TypeOf((*MockCluster)(nil).Inspect), arg0)
}

// InspectDomain mocks base method.
func (m *MockCluster) InspectDomain(arg0 string) (*clusterdomain.ClusterDomainInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectDomain", arg0)
	ret0, _ := ret[0].(*clusterdomain.ClusterDomainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectDomain indicates an expected call of InspectDomain.
func (mr *MockClusterMockRecorder) InspectDomain(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectDomain", reflect.TypeOf((*MockCluster)(nil).InspectDomain), arg0)
}

// NodeRemoveDone mocks base method.
func (m *MockCluster) NodeRemoveDone(arg0 string, arg1 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NodeRemoveDone", arg0, arg1)
}

// NodeRemoveDone indicates an expected call of NodeRemoveDone.
func (mr *MockClusterMockRecorder) NodeRemoveDone(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeRemoveDone", reflect.TypeOf((*MockCluster)(nil).NodeRemoveDone), arg0, arg1)
}

// NodeStatus mocks base method.
func (m *MockCluster) NodeStatus() (api.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeStatus")
	ret0, _ := ret[0].(api.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeStatus indicates an expected call of NodeStatus.
func (mr *MockClusterMockRecorder) NodeStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeStatus", reflect.TypeOf((*MockCluster)(nil).NodeStatus))
}

// ObjectStoreCreate mocks base method.
func (m *MockCluster) ObjectStoreCreate(arg0 string) (*api.ObjectstoreInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectStoreCreate", arg0)
	ret0, _ := ret[0].(*api.ObjectstoreInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectStoreCreate indicates an expected call of ObjectStoreCreate.
func (mr *MockClusterMockRecorder) ObjectStoreCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStoreCreate", reflect.TypeOf((*MockCluster)(nil).ObjectStoreCreate), arg0)
}

// ObjectStoreDelete mocks base method.
func (m *MockCluster) ObjectStoreDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectStoreDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ObjectStoreDelete indicates an expected call of ObjectStoreDelete.
func (mr *MockClusterMockRecorder) ObjectStoreDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStoreDelete", reflect.TypeOf((*MockCluster)(nil).ObjectStoreDelete), arg0)
}

// ObjectStoreInspect mocks base method.
func (m *MockCluster) ObjectStoreInspect(arg0 string) (*api.ObjectstoreInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectStoreInspect", arg0)
	ret0, _ := ret[0].(*api.ObjectstoreInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectStoreInspect indicates an expected call of ObjectStoreInspect.
func (mr *MockClusterMockRecorder) ObjectStoreInspect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStoreInspect", reflect.TypeOf((*MockCluster)(nil).ObjectStoreInspect), arg0)
}

// ObjectStoreUpdate mocks base method.
func (m *MockCluster) ObjectStoreUpdate(arg0 string, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectStoreUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ObjectStoreUpdate indicates an expected call of ObjectStoreUpdate.
func (mr *MockClusterMockRecorder) ObjectStoreUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStoreUpdate", reflect.TypeOf((*MockCluster)(nil).ObjectStoreUpdate), arg0, arg1)
}

// PeerStatus mocks base method.
func (m *MockCluster) PeerStatus(arg0 string) (map[string]api.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeerStatus", arg0)
	ret0, _ := ret[0].(map[string]api.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeerStatus indicates an expected call of PeerStatus.
func (mr *MockClusterMockRecorder) PeerStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerStatus", reflect.TypeOf((*MockCluster)(nil).PeerStatus), arg0)
}

// ProcessPairRequest mocks base method.
func (m *MockCluster) ProcessPairRequest(arg0 *api.ClusterPairProcessRequest) (*api.ClusterPairProcessResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessPairRequest", arg0)
	ret0, _ := ret[0].(*api.ClusterPairProcessResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessPairRequest indicates an expected call of ProcessPairRequest.
func (mr *MockClusterMockRecorder) ProcessPairRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessPairRequest", reflect.TypeOf((*MockCluster)(nil).ProcessPairRequest), arg0)
}

// Rebalance mocks base method.
func (m *MockCluster) Rebalance(arg0 context.Context, arg1 *api.SdkStorageRebalanceRequest) (*api.SdkStorageRebalanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rebalance", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkStorageRebalanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rebalance indicates an expected call of Rebalance.
func (mr *MockClusterMockRecorder) Rebalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rebalance", reflect.TypeOf((*MockCluster)(nil).Rebalance), arg0, arg1)
}

// RefreshPair mocks base method.
func (m *MockCluster) RefreshPair(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshPair", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshPair indicates an expected call of RefreshPair.
func (mr *MockClusterMockRecorder) RefreshPair(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshPair", reflect.TypeOf((*MockCluster)(nil).RefreshPair), arg0)
}

// Remove mocks base method.
func (m *MockCluster) Remove(arg0 []api.Node, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockClusterMockRecorder) Remove(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockCluster)(nil).Remove), arg0, arg1)
}

// RemoveVolumeAttachments mocks base method.
func (m *MockCluster) RemoveVolumeAttachments(arg0 context.Context, arg1 *api.SdkNodeRemoveVolumeAttachmentsRequest) (*api.SdkJobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveVolumeAttachments", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkJobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveVolumeAttachments indicates an expected call of RemoveVolumeAttachments.
func (mr *MockClusterMockRecorder) RemoveVolumeAttachments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveVolumeAttachments", reflect.TypeOf((*MockCluster)(nil).RemoveVolumeAttachments), arg0, arg1)
}

// Resize mocks base method.
func (m *MockCluster) Resize(arg0 context.Context, arg1 *api.SdkStoragePoolResizeRequest) (*api.SdkStoragePoolResizeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resize", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkStoragePoolResizeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resize indicates an expected call of Resize.
func (mr *MockClusterMockRecorder) Resize(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resize", reflect.TypeOf((*MockCluster)(nil).Resize), arg0, arg1)
}

// SchedPolicyCreate mocks base method.
func (m *MockCluster) SchedPolicyCreate(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchedPolicyCreate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SchedPolicyCreate indicates an expected call of SchedPolicyCreate.
func (mr *MockClusterMockRecorder) SchedPolicyCreate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedPolicyCreate", reflect.TypeOf((*MockCluster)(nil).SchedPolicyCreate), arg0, arg1)
}

// SchedPolicyDelete mocks base method.
func (m *MockCluster) SchedPolicyDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchedPolicyDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SchedPolicyDelete indicates an expected call of SchedPolicyDelete.
func (mr *MockClusterMockRecorder) SchedPolicyDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedPolicyDelete", reflect.TypeOf((*MockCluster)(nil).SchedPolicyDelete), arg0)
}

// SchedPolicyEnumerate mocks base method.
func (m *MockCluster) SchedPolicyEnumerate() ([]*schedpolicy.SchedPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchedPolicyEnumerate")
	ret0, _ := ret[0].([]*schedpolicy.SchedPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchedPolicyEnumerate indicates an expected call of SchedPolicyEnumerate.
func (mr *MockClusterMockRecorder) SchedPolicyEnumerate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedPolicyEnumerate", reflect.TypeOf((*MockCluster)(nil).SchedPolicyEnumerate))
}

// SchedPolicyGet mocks base method.
func (m *MockCluster) SchedPolicyGet(arg0 string) (*schedpolicy.SchedPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchedPolicyGet", arg0)
	ret0, _ := ret[0].(*schedpolicy.SchedPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchedPolicyGet indicates an expected call of SchedPolicyGet.
func (mr *MockClusterMockRecorder) SchedPolicyGet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedPolicyGet", reflect.TypeOf((*MockCluster)(nil).SchedPolicyGet), arg0)
}

// SchedPolicyUpdate mocks base method.
func (m *MockCluster) SchedPolicyUpdate(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchedPolicyUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SchedPolicyUpdate indicates an expected call of SchedPolicyUpdate.
func (mr *MockClusterMockRecorder) SchedPolicyUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedPolicyUpdate", reflect.TypeOf((*MockCluster)(nil).SchedPolicyUpdate), arg0, arg1)
}

// SecretCheckLogin mocks base method.
func (m *MockCluster) SecretCheckLogin() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretCheckLogin")
	ret0, _ := ret[0].(error)
	return ret0
}

// SecretCheckLogin indicates an expected call of SecretCheckLogin.
func (mr *MockClusterMockRecorder) SecretCheckLogin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretCheckLogin", reflect.TypeOf((*MockCluster)(nil).SecretCheckLogin))
}

// SecretGet mocks base method.
func (m *MockCluster) SecretGet(arg0 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretGet", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecretGet indicates an expected call of SecretGet.
func (mr *MockClusterMockRecorder) SecretGet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretGet", reflect.TypeOf((*MockCluster)(nil).SecretGet), arg0)
}

// SecretGetDefaultSecretKey mocks base method.
func (m *MockCluster) SecretGetDefaultSecretKey() (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretGetDefaultSecretKey")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecretGetDefaultSecretKey indicates an expected call of SecretGetDefaultSecretKey.
func (mr *MockClusterMockRecorder) SecretGetDefaultSecretKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretGetDefaultSecretKey", reflect.TypeOf((*MockCluster)(nil).SecretGetDefaultSecretKey))
}

// SecretLogin mocks base method.
func (m *MockCluster) SecretLogin(arg0 string, arg1 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretLogin", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SecretLogin indicates an expected call of SecretLogin.
func (mr *MockClusterMockRecorder) SecretLogin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretLogin", reflect.TypeOf((*MockCluster)(nil).SecretLogin), arg0, arg1)
}

// SecretSet mocks base method.
func (m *MockCluster) SecretSet(arg0 string, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretSet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SecretSet indicates an expected call of SecretSet.
func (mr *MockClusterMockRecorder) SecretSet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretSet", reflect.TypeOf((*MockCluster)(nil).SecretSet), arg0, arg1)
}

// SecretSetDefaultSecretKey mocks base method.
func (m *MockCluster) SecretSetDefaultSecretKey(arg0 string, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretSetDefaultSecretKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SecretSetDefaultSecretKey indicates an expected call of SecretSetDefaultSecretKey.
func (mr *MockClusterMockRecorder) SecretSetDefaultSecretKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretSetDefaultSecretKey", reflect.TypeOf((*MockCluster)(nil).SecretSetDefaultSecretKey), arg0, arg1)
}

// SetClusterConf mocks base method.
func (m *MockCluster) SetClusterConf(arg0 *osdconfig.ClusterConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetClusterConf", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetClusterConf indicates an expected call of SetClusterConf.
func (mr *MockClusterMockRecorder) SetClusterConf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClusterConf", reflect.TypeOf((*MockCluster)(nil).SetClusterConf), arg0)
}

// SetNodeConf mocks base method.
func (m *MockCluster) SetNodeConf(arg0 *osdconfig.NodeConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNodeConf", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetNodeConf indicates an expected call of SetNodeConf.
func (mr *MockClusterMockRecorder) SetNodeConf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNodeConf", reflect.TypeOf((*MockCluster)(nil).SetNodeConf), arg0)
}

// SetSize mocks base method.
func (m *MockCluster) SetSize(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSize indicates an expected call of SetSize.
func (mr *MockClusterMockRecorder) SetSize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSize", reflect.TypeOf((*MockCluster)(nil).SetSize), arg0)
}

// Shutdown mocks base method.
func (m *MockCluster) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockClusterMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockCluster)(nil).Shutdown))
}

// Start mocks base method.
func (m *MockCluster) Start(arg0 bool, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockClusterMockRecorder) Start(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockCluster)(nil).Start), arg0, arg1, arg2)
}

// StartWithConfiguration mocks base method.
func (m *MockCluster) StartWithConfiguration(arg0 bool, arg1 string, arg2 []string, arg3 string, arg4 *cluster.ClusterServerConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartWithConfiguration", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartWithConfiguration indicates an expected call of StartWithConfiguration.
func (mr *MockClusterMockRecorder) StartWithConfiguration(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWithConfiguration", reflect.TypeOf((*MockCluster)(nil).StartWithConfiguration), arg0, arg1, arg2, arg3, arg4)
}

// UpdateData mocks base method.
func (m *MockCluster) UpdateData(arg0 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateData", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateData indicates an expected call of UpdateData.
func (mr *MockClusterMockRecorder) UpdateData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateData", reflect.TypeOf((*MockCluster)(nil).UpdateData), arg0)
}

// UpdateDomainState mocks base method.
func (m *MockCluster) UpdateDomainState(arg0 string, arg1 types.ClusterDomainState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDomainState", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDomainState indicates an expected call of UpdateDomainState.
func (mr *MockClusterMockRecorder) UpdateDomainState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDomainState", reflect.TypeOf((*MockCluster)(nil).UpdateDomainState), arg0, arg1)
}

// UpdateJobState mocks base method.
func (m *MockCluster) UpdateJobState(arg0 context.Context, arg1 *api.SdkUpdateJobRequest) (*api.SdkUpdateJobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJobState", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkUpdateJobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateJobState indicates an expected call of UpdateJobState.
func (mr *MockClusterMockRecorder) UpdateJobState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJobState", reflect.TypeOf((*MockCluster)(nil).UpdateJobState), arg0, arg1)
}

// UpdateLabels mocks base method.
func (m *MockCluster) UpdateLabels(arg0 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLabels", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLabels indicates an expected call of UpdateLabels.
func (mr *MockClusterMockRecorder) UpdateLabels(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLabels", reflect.TypeOf((*MockCluster)(nil).UpdateLabels), arg0)
}

// UpdateRebalanceJobState mocks base method.
func (m *MockCluster) UpdateRebalanceJobState(arg0 context.Context, arg1 *api.SdkUpdateRebalanceJobRequest) (*api.SdkUpdateRebalanceJobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRebalanceJobState", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkUpdateRebalanceJobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRebalanceJobState indicates an expected call of UpdateRebalanceJobState.
func (mr *MockClusterMockRecorder) UpdateRebalanceJobState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRebalanceJobState", reflect.TypeOf((*MockCluster)(nil).UpdateRebalanceJobState), arg0, arg1)
}

// UpdateSchedulerNodeName mocks base method.
func (m *MockCluster) UpdateSchedulerNodeName(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSchedulerNodeName", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSchedulerNodeName indicates an expected call of UpdateSchedulerNodeName.
func (mr *MockClusterMockRecorder) UpdateSchedulerNodeName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSchedulerNodeName", reflect.TypeOf((*MockCluster)(nil).UpdateSchedulerNodeName), arg0)
}

// Uuid mocks base method.
func (m *MockCluster) Uuid() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uuid")
	ret0, _ := ret[0].(string)
	return ret0
}

// Uuid indicates an expected call of Uuid.
func (mr *MockClusterMockRecorder) Uuid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uuid", reflect.TypeOf((*MockCluster)(nil).Uuid))
}

// ValidatePair mocks base method.
func (m *MockCluster) ValidatePair(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatePair", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidatePair indicates an expected call of ValidatePair.
func (mr *MockClusterMockRecorder) ValidatePair(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatePair", reflect.TypeOf((*MockCluster)(nil).ValidatePair), arg0)
}
