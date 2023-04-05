// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-azure-devops/server/store (interfaces: KVStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	serializers "github.com/mattermost/mattermost-plugin-azure-devops/server/serializers"
	store "github.com/mattermost/mattermost-plugin-azure-devops/server/store"
	reflect "reflect"
)

// MockKVStore is a mock of KVStore interface
type MockKVStore struct {
	ctrl     *gomock.Controller
	recorder *MockKVStoreMockRecorder
}

// MockKVStoreMockRecorder is the mock recorder for MockKVStore
type MockKVStoreMockRecorder struct {
	mock *MockKVStore
}

// NewMockKVStore creates a new mock instance
func NewMockKVStore(ctrl *gomock.Controller) *MockKVStore {
	mock := &MockKVStore{ctrl: ctrl}
	mock.recorder = &MockKVStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKVStore) EXPECT() *MockKVStoreMockRecorder {
	return m.recorder
}

// DeleteProject mocks base method
func (m *MockKVStore) DeleteProject(arg0 *serializers.ProjectDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProject indicates an expected call of DeleteProject
func (mr *MockKVStoreMockRecorder) DeleteProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockKVStore)(nil).DeleteProject), arg0)
}

// DeleteSubscription mocks base method
func (m *MockKVStore) DeleteSubscription(arg0 *serializers.SubscriptionDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubscription", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubscription indicates an expected call of DeleteSubscription
func (mr *MockKVStoreMockRecorder) DeleteSubscription(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubscription", reflect.TypeOf((*MockKVStore)(nil).DeleteSubscription), arg0)
}

// DeleteUser mocks base method
func (m *MockKVStore) DeleteUser(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockKVStoreMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockKVStore)(nil).DeleteUser), arg0)
}

// DeleteUserTokenOnEncryptionSecretChange mocks base method
func (m *MockKVStore) DeleteUserTokenOnEncryptionSecretChange() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserTokenOnEncryptionSecretChange")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserTokenOnEncryptionSecretChange indicates an expected call of DeleteUserTokenOnEncryptionSecretChange
func (mr *MockKVStoreMockRecorder) DeleteUserTokenOnEncryptionSecretChange() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserTokenOnEncryptionSecretChange", reflect.TypeOf((*MockKVStore)(nil).DeleteUserTokenOnEncryptionSecretChange))
}

// GetAllProjects mocks base method
func (m *MockKVStore) GetAllProjects(arg0 string) ([]serializers.ProjectDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProjects", arg0)
	ret0, _ := ret[0].([]serializers.ProjectDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProjects indicates an expected call of GetAllProjects
func (mr *MockKVStoreMockRecorder) GetAllProjects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProjects", reflect.TypeOf((*MockKVStore)(nil).GetAllProjects), arg0)
}

// GetAllSubscriptions mocks base method
func (m *MockKVStore) GetAllSubscriptions(arg0 string) ([]*serializers.SubscriptionDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSubscriptions", arg0)
	ret0, _ := ret[0].([]*serializers.SubscriptionDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSubscriptions indicates an expected call of GetAllSubscriptions
func (mr *MockKVStoreMockRecorder) GetAllSubscriptions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSubscriptions", reflect.TypeOf((*MockKVStore)(nil).GetAllSubscriptions), arg0)
}

// GetProject mocks base method
func (m *MockKVStore) GetProject() (*store.ProjectList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject")
	ret0, _ := ret[0].(*store.ProjectList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject
func (mr *MockKVStoreMockRecorder) GetProject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockKVStore)(nil).GetProject))
}

// GetSubscriptionList mocks base method
func (m *MockKVStore) GetSubscriptionList() (*store.SubscriptionList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptionList")
	ret0, _ := ret[0].(*store.SubscriptionList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptionList indicates an expected call of GetSubscriptionList
func (mr *MockKVStoreMockRecorder) GetSubscriptionList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionList", reflect.TypeOf((*MockKVStore)(nil).GetSubscriptionList))
}

// LoadAzureDevopsUserIDFromMattermostUser mocks base method
func (m *MockKVStore) LoadAzureDevopsUserIDFromMattermostUser(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadAzureDevopsUserIDFromMattermostUser", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadAzureDevopsUserIDFromMattermostUser indicates an expected call of LoadAzureDevopsUserIDFromMattermostUser
func (mr *MockKVStoreMockRecorder) LoadAzureDevopsUserIDFromMattermostUser(arg0 string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadAzureDevopsUserIDFromMattermostUser", reflect.TypeOf((*MockKVStore)(nil).LoadAzureDevopsUserIDFromMattermostUser), arg0)
}

// LoadAzureDevopsUserDetails mocks base method
func (m *MockKVStore) LoadAzureDevopsUserDetails(arg0 string) (*serializers.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadAzureDevopsUserDetails", arg0)
	ret0, _ := ret[0].(*serializers.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadAzureDevopsUserDetails indicates an expected call of LoadAzureDevopsUserDetails
func (mr *MockKVStoreMockRecorder) LoadAzureDevopsUserDetails(arg0 string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadAzureDevopsUserDetails", reflect.TypeOf((*MockKVStore)(nil).LoadAzureDevopsUserDetails), arg0)
}

// StoreOAuthState mocks base method
func (m *MockKVStore) StoreOAuthState(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreOAuthState", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreOAuthState indicates an expected call of StoreOAuthState
func (mr *MockKVStoreMockRecorder) StoreOAuthState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreOAuthState", reflect.TypeOf((*MockKVStore)(nil).StoreOAuthState), arg0, arg1)
}

// StoreProject mocks base method
func (m *MockKVStore) StoreProject(arg0 *serializers.ProjectDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreProject", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreProject indicates an expected call of StoreProject
func (mr *MockKVStoreMockRecorder) StoreProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreProject", reflect.TypeOf((*MockKVStore)(nil).StoreProject), arg0)
}

// StoreSubscription mocks base method
func (m *MockKVStore) StoreSubscription(arg0 *serializers.SubscriptionDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreSubscription", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreSubscription indicates an expected call of StoreSubscription
func (mr *MockKVStoreMockRecorder) StoreSubscription(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreSubscription", reflect.TypeOf((*MockKVStore)(nil).StoreSubscription), arg0)
}

// StoreAzureDevopsUserDetailsWithMattermostUserID mocks base method
func (m *MockKVStore) StoreAzureDevopsUserDetailsWithMattermostUserID(arg0 *serializers.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreAzureDevopsUserDetailsWithMattermostUserID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreAzureDevopsUserDetailsWithMattermostUserID indicates an expected call of StoreAzureDevopsUserDetailsWithMattermostUserID
func (mr *MockKVStoreMockRecorder) StoreAzureDevopsUserDetailsWithMattermostUserID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreAzureDevopsUserDetailsWithMattermostUserID", reflect.TypeOf((*MockKVStore)(nil).StoreAzureDevopsUserDetailsWithMattermostUserID), arg0)
}

// VerifyOAuthState mocks base method
func (m *MockKVStore) VerifyOAuthState(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyOAuthState", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyOAuthState indicates an expected call of VerifyOAuthState
func (mr *MockKVStoreMockRecorder) VerifyOAuthState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyOAuthState", reflect.TypeOf((*MockKVStore)(nil).VerifyOAuthState), arg0, arg1)
}

// StoreSubscriptionAndChannelIDMap mocks base method
func (m *MockKVStore) StoreSubscriptionAndChannelIDMap(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreSubscriptionAndChannelIDMap", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreSubscriptionAndChannelIDMap indicates an expected call of StoreSubscriptionAndChannelIDMap
func (mr *MockKVStoreMockRecorder) StoreSubscriptionAndChannelIDMap(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreSubscriptionAndChannelIDMap", reflect.TypeOf((*MockKVStore)(nil).StoreSubscriptionAndChannelIDMap), arg0, arg1, arg2)
}

// GetSubscriptionAndChannelIDMap mocks base method
func (m *MockKVStore) GetSubscriptionAndChannelIDMap(arg0 string) (*store.SubscriptionWebhookSecretAndChannelMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptionAndChannelIDMap", arg0)
	ret0, _ := ret[0].(*store.SubscriptionWebhookSecretAndChannelMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptionAndChannelIDMap indicates an expected call of GetSubscriptionAndChannelIDMap
func (mr *MockKVStoreMockRecorder) GetSubscriptionAndChannelIDMap(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionAndChannelIDMap", reflect.TypeOf((*MockKVStore)(nil).GetSubscriptionAndChannelIDMap), arg0)
}

// DeleteSubscriptionAndChannelIDMap mocks base method
func (m *MockKVStore) DeleteSubscriptionAndChannelIDMap(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubscriptionAndChannelIDMap", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubscriptionAndChannelIDMap indicates an expected call of DeleteSubscriptionAndChannelIDMap
func (mr *MockKVStoreMockRecorder) DeleteSubscriptionAndChannelIDMap(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubscriptionAndChannelIDMap", reflect.TypeOf((*MockKVStore)(nil).DeleteSubscriptionAndChannelIDMap), arg0)
}
