/*
Copyright 2022 The Fluid Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/ddc/base/engine.go

// Package base is a generated GoMock package.
package base

import (
	reflect "reflect"

	v1alpha1 "github.com/fluid-cloudnative/fluid/api/v1alpha1"
	runtime "github.com/fluid-cloudnative/fluid/pkg/runtime"
	utils "github.com/fluid-cloudnative/fluid/pkg/utils"
	gomock "github.com/golang/mock/gomock"
)

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// CheckExistenceOfPath mocks base method.
func (m *MockEngine) CheckExistenceOfPath(targetDataload v1alpha1.DataLoad) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistenceOfPath", targetDataload)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckExistenceOfPath indicates an expected call of CheckExistenceOfPath.
func (mr *MockEngineMockRecorder) CheckExistenceOfPath(targetDataload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistenceOfPath", reflect.TypeOf((*MockEngine)(nil).CheckExistenceOfPath), targetDataload)
}

// CheckRuntimeReady mocks base method.
func (m *MockEngine) CheckRuntimeReady() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRuntimeReady")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckRuntimeReady indicates an expected call of CheckRuntimeReady.
func (mr *MockEngineMockRecorder) CheckRuntimeReady() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRuntimeReady", reflect.TypeOf((*MockEngine)(nil).CheckRuntimeReady))
}

// CreateVolume mocks base method.
func (m *MockEngine) CreateVolume() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolume")
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVolume indicates an expected call of CreateVolume.
func (mr *MockEngineMockRecorder) CreateVolume() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockEngine)(nil).CreateVolume))
}

// DeleteVolume mocks base method.
func (m *MockEngine) DeleteVolume() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVolume indicates an expected call of DeleteVolume.
func (mr *MockEngineMockRecorder) DeleteVolume() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockEngine)(nil).DeleteVolume))
}

// ID mocks base method.
func (m *MockEngine) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockEngineMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockEngine)(nil).ID))
}

// LoadData mocks base method.
func (m *MockEngine) LoadData(ctx runtime.ReconcileRequestContext, targetDataload v1alpha1.DataLoad) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadData", ctx, targetDataload)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadData indicates an expected call of LoadData.
func (mr *MockEngineMockRecorder) LoadData(ctx, targetDataload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadData", reflect.TypeOf((*MockEngine)(nil).LoadData), ctx, targetDataload)
}

// Setup mocks base method.
func (m *MockEngine) Setup(ctx runtime.ReconcileRequestContext) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Setup", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Setup indicates an expected call of Setup.
func (mr *MockEngineMockRecorder) Setup(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockEngine)(nil).Setup), ctx)
}

// Shutdown mocks base method.
func (m *MockEngine) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockEngineMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockEngine)(nil).Shutdown))
}

// Sync mocks base method.
func (m *MockEngine) Sync(ctx runtime.ReconcileRequestContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync.
func (mr *MockEngineMockRecorder) Sync(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockEngine)(nil).Sync), ctx)
}

// MockDataloader is a mock of Dataloader interface.
type MockDataloader struct {
	ctrl     *gomock.Controller
	recorder *MockDataloaderMockRecorder
}

// MockDataloaderMockRecorder is the mock recorder for MockDataloader.
type MockDataloaderMockRecorder struct {
	mock *MockDataloader
}

// NewMockDataloader creates a new mock instance.
func NewMockDataloader(ctrl *gomock.Controller) *MockDataloader {
	mock := &MockDataloader{ctrl: ctrl}
	mock.recorder = &MockDataloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataloader) EXPECT() *MockDataloaderMockRecorder {
	return m.recorder
}

// CheckExistenceOfPath mocks base method.
func (m *MockDataloader) CheckExistenceOfPath(targetDataload v1alpha1.DataLoad) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistenceOfPath", targetDataload)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckExistenceOfPath indicates an expected call of CheckExistenceOfPath.
func (mr *MockDataloaderMockRecorder) CheckExistenceOfPath(targetDataload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistenceOfPath", reflect.TypeOf((*MockDataloader)(nil).CheckExistenceOfPath), targetDataload)
}

// CheckRuntimeReady mocks base method.
func (m *MockDataloader) CheckRuntimeReady() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRuntimeReady")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckRuntimeReady indicates an expected call of CheckRuntimeReady.
func (mr *MockDataloaderMockRecorder) CheckRuntimeReady() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRuntimeReady", reflect.TypeOf((*MockDataloader)(nil).CheckRuntimeReady))
}

// LoadData mocks base method.
func (m *MockDataloader) LoadData(ctx runtime.ReconcileRequestContext, targetDataload v1alpha1.DataLoad) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadData", ctx, targetDataload)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadData indicates an expected call of LoadData.
func (mr *MockDataloaderMockRecorder) LoadData(ctx, targetDataload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadData", reflect.TypeOf((*MockDataloader)(nil).LoadData), ctx, targetDataload)
}

// MockImplement is a mock of Implement interface.
type MockImplement struct {
	ctrl     *gomock.Controller
	recorder *MockImplementMockRecorder
}

// MockImplementMockRecorder is the mock recorder for MockImplement.
type MockImplementMockRecorder struct {
	mock *MockImplement
}

// NewMockImplement creates a new mock instance.
func NewMockImplement(ctrl *gomock.Controller) *MockImplement {
	mock := &MockImplement{ctrl: ctrl}
	mock.recorder = &MockImplementMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImplement) EXPECT() *MockImplementMockRecorder {
	return m.recorder
}

// AssignNodesToCache mocks base method.
func (m *MockImplement) AssignNodesToCache(desiredNum int32) (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignNodesToCache", desiredNum)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignNodesToCache indicates an expected call of AssignNodesToCache.
func (mr *MockImplementMockRecorder) AssignNodesToCache(desiredNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignNodesToCache", reflect.TypeOf((*MockImplement)(nil).AssignNodesToCache), desiredNum)
}

// BindToDataset mocks base method.
func (m *MockImplement) BindToDataset() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindToDataset")
	ret0, _ := ret[0].(error)
	return ret0
}

// BindToDataset indicates an expected call of BindToDataset.
func (mr *MockImplementMockRecorder) BindToDataset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindToDataset", reflect.TypeOf((*MockImplement)(nil).BindToDataset))
}

// CheckAndUpdateRuntimeStatus mocks base method.
func (m *MockImplement) CheckAndUpdateRuntimeStatus() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAndUpdateRuntimeStatus")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAndUpdateRuntimeStatus indicates an expected call of CheckAndUpdateRuntimeStatus.
func (mr *MockImplementMockRecorder) CheckAndUpdateRuntimeStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAndUpdateRuntimeStatus", reflect.TypeOf((*MockImplement)(nil).CheckAndUpdateRuntimeStatus))
}

// CheckExistenceOfPath mocks base method.
func (m *MockImplement) CheckExistenceOfPath(targetDataload v1alpha1.DataLoad) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistenceOfPath", targetDataload)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckExistenceOfPath indicates an expected call of CheckExistenceOfPath.
func (mr *MockImplementMockRecorder) CheckExistenceOfPath(targetDataload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistenceOfPath", reflect.TypeOf((*MockImplement)(nil).CheckExistenceOfPath), targetDataload)
}

// CheckMasterReady mocks base method.
func (m *MockImplement) CheckMasterReady() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckMasterReady")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckMasterReady indicates an expected call of CheckMasterReady.
func (mr *MockImplementMockRecorder) CheckMasterReady() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckMasterReady", reflect.TypeOf((*MockImplement)(nil).CheckMasterReady))
}

// CheckRuntimeHealthy mocks base method.
func (m *MockImplement) CheckRuntimeHealthy() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRuntimeHealthy")
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckRuntimeHealthy indicates an expected call of CheckRuntimeHealthy.
func (mr *MockImplementMockRecorder) CheckRuntimeHealthy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRuntimeHealthy", reflect.TypeOf((*MockImplement)(nil).CheckRuntimeHealthy))
}

// CheckRuntimeReady mocks base method.
func (m *MockImplement) CheckRuntimeReady() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRuntimeReady")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckRuntimeReady indicates an expected call of CheckRuntimeReady.
func (mr *MockImplementMockRecorder) CheckRuntimeReady() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRuntimeReady", reflect.TypeOf((*MockImplement)(nil).CheckRuntimeReady))
}

// CheckWorkersReady mocks base method.
func (m *MockImplement) CheckWorkersReady() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckWorkersReady")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckWorkersReady indicates an expected call of CheckWorkersReady.
func (mr *MockImplementMockRecorder) CheckWorkersReady() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckWorkersReady", reflect.TypeOf((*MockImplement)(nil).CheckWorkersReady))
}

// CreateDataLoadJob mocks base method.
func (m *MockImplement) CreateDataLoadJob(ctx runtime.ReconcileRequestContext, targetDataload v1alpha1.DataLoad) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDataLoadJob", ctx, targetDataload)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDataLoadJob indicates an expected call of CreateDataLoadJob.
func (mr *MockImplementMockRecorder) CreateDataLoadJob(ctx, targetDataload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDataLoadJob", reflect.TypeOf((*MockImplement)(nil).CreateDataLoadJob), ctx, targetDataload)
}

// CreateVolume mocks base method.
func (m *MockImplement) CreateVolume() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolume")
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVolume indicates an expected call of CreateVolume.
func (mr *MockImplementMockRecorder) CreateVolume() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockImplement)(nil).CreateVolume))
}

// DeleteVolume mocks base method.
func (m *MockImplement) DeleteVolume() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVolume indicates an expected call of DeleteVolume.
func (mr *MockImplementMockRecorder) DeleteVolume() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockImplement)(nil).DeleteVolume))
}

// FreeStorageBytes mocks base method.
func (m *MockImplement) FreeStorageBytes() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FreeStorageBytes")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FreeStorageBytes indicates an expected call of FreeStorageBytes.
func (mr *MockImplementMockRecorder) FreeStorageBytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FreeStorageBytes", reflect.TypeOf((*MockImplement)(nil).FreeStorageBytes))
}

// PrepareUFS mocks base method.
func (m *MockImplement) PrepareUFS() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareUFS")
	ret0, _ := ret[0].(error)
	return ret0
}

// PrepareUFS indicates an expected call of PrepareUFS.
func (mr *MockImplementMockRecorder) PrepareUFS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareUFS", reflect.TypeOf((*MockImplement)(nil).PrepareUFS))
}

// SetupMaster mocks base method.
func (m *MockImplement) SetupMaster() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupMaster")
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupMaster indicates an expected call of SetupMaster.
func (mr *MockImplementMockRecorder) SetupMaster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupMaster", reflect.TypeOf((*MockImplement)(nil).SetupMaster))
}

// SetupWorkers mocks base method.
func (m *MockImplement) SetupWorkers() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupWorkers")
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupWorkers indicates an expected call of SetupWorkers.
func (mr *MockImplementMockRecorder) SetupWorkers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupWorkers", reflect.TypeOf((*MockImplement)(nil).SetupWorkers))
}

// ShouldCheckUFS mocks base method.
func (m *MockImplement) ShouldCheckUFS() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldCheckUFS")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShouldCheckUFS indicates an expected call of ShouldCheckUFS.
func (mr *MockImplementMockRecorder) ShouldCheckUFS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldCheckUFS", reflect.TypeOf((*MockImplement)(nil).ShouldCheckUFS))
}

// ShouldSetupMaster mocks base method.
func (m *MockImplement) ShouldSetupMaster() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldSetupMaster")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShouldSetupMaster indicates an expected call of ShouldSetupMaster.
func (mr *MockImplementMockRecorder) ShouldSetupMaster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldSetupMaster", reflect.TypeOf((*MockImplement)(nil).ShouldSetupMaster))
}

// ShouldSetupWorkers mocks base method.
func (m *MockImplement) ShouldSetupWorkers() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldSetupWorkers")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShouldSetupWorkers indicates an expected call of ShouldSetupWorkers.
func (mr *MockImplementMockRecorder) ShouldSetupWorkers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldSetupWorkers", reflect.TypeOf((*MockImplement)(nil).ShouldSetupWorkers))
}

// ShouldUpdateUFS mocks base method.
func (m *MockImplement) ShouldUpdateUFS() *utils.UFSToUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldUpdateUFS")
	ret0, _ := ret[0].(*utils.UFSToUpdate)
	return ret0
}

// ShouldUpdateUFS indicates an expected call of ShouldUpdateUFS.
func (mr *MockImplementMockRecorder) ShouldUpdateUFS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldUpdateUFS", reflect.TypeOf((*MockImplement)(nil).ShouldUpdateUFS))
}

// Shutdown mocks base method.
func (m *MockImplement) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockImplementMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockImplement)(nil).Shutdown))
}

// SyncMetadata mocks base method.
func (m *MockImplement) SyncMetadata() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncMetadata")
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncMetadata indicates an expected call of SyncMetadata.
func (mr *MockImplementMockRecorder) SyncMetadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncMetadata", reflect.TypeOf((*MockImplement)(nil).SyncMetadata))
}

// SyncReplicas mocks base method.
func (m *MockImplement) SyncReplicas(ctx runtime.ReconcileRequestContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncReplicas", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncReplicas indicates an expected call of SyncReplicas.
func (mr *MockImplementMockRecorder) SyncReplicas(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncReplicas", reflect.TypeOf((*MockImplement)(nil).SyncReplicas), ctx)
}

// SyncScheduleInfoToCacheNodes mocks base method.
func (m *MockImplement) SyncScheduleInfoToCacheNodes() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncScheduleInfoToCacheNodes")
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncScheduleInfoToCacheNodes indicates an expected call of SyncScheduleInfoToCacheNodes.
func (mr *MockImplementMockRecorder) SyncScheduleInfoToCacheNodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncScheduleInfoToCacheNodes", reflect.TypeOf((*MockImplement)(nil).SyncScheduleInfoToCacheNodes))
}

// TotalFileNums mocks base method.
func (m *MockImplement) TotalFileNums() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalFileNums")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TotalFileNums indicates an expected call of TotalFileNums.
func (mr *MockImplementMockRecorder) TotalFileNums() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalFileNums", reflect.TypeOf((*MockImplement)(nil).TotalFileNums))
}

// TotalStorageBytes mocks base method.
func (m *MockImplement) TotalStorageBytes() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalStorageBytes")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TotalStorageBytes indicates an expected call of TotalStorageBytes.
func (mr *MockImplementMockRecorder) TotalStorageBytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalStorageBytes", reflect.TypeOf((*MockImplement)(nil).TotalStorageBytes))
}

// UpdateCacheOfDataset mocks base method.
func (m *MockImplement) UpdateCacheOfDataset() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCacheOfDataset")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCacheOfDataset indicates an expected call of UpdateCacheOfDataset.
func (mr *MockImplementMockRecorder) UpdateCacheOfDataset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCacheOfDataset", reflect.TypeOf((*MockImplement)(nil).UpdateCacheOfDataset))
}

// UpdateDatasetStatus mocks base method.
func (m *MockImplement) UpdateDatasetStatus(phase v1alpha1.DatasetPhase) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDatasetStatus", phase)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDatasetStatus indicates an expected call of UpdateDatasetStatus.
func (mr *MockImplementMockRecorder) UpdateDatasetStatus(phase interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDatasetStatus", reflect.TypeOf((*MockImplement)(nil).UpdateDatasetStatus), phase)
}

// UpdateOnUFSChange mocks base method.
func (m *MockImplement) UpdateOnUFSChange(ufsToUpdate *utils.UFSToUpdate) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOnUFSChange", ufsToUpdate)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOnUFSChange indicates an expected call of UpdateOnUFSChange.
func (mr *MockImplementMockRecorder) UpdateOnUFSChange(ufsToUpdate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOnUFSChange", reflect.TypeOf((*MockImplement)(nil).UpdateOnUFSChange), ufsToUpdate)
}

// UsedStorageBytes mocks base method.
func (m *MockImplement) UsedStorageBytes() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsedStorageBytes")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UsedStorageBytes indicates an expected call of UsedStorageBytes.
func (mr *MockImplementMockRecorder) UsedStorageBytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsedStorageBytes", reflect.TypeOf((*MockImplement)(nil).UsedStorageBytes))
}

// MockUnderFileSystemService is a mock of UnderFileSystemService interface.
type MockUnderFileSystemService struct {
	ctrl     *gomock.Controller
	recorder *MockUnderFileSystemServiceMockRecorder
}

// MockUnderFileSystemServiceMockRecorder is the mock recorder for MockUnderFileSystemService.
type MockUnderFileSystemServiceMockRecorder struct {
	mock *MockUnderFileSystemService
}

// NewMockUnderFileSystemService creates a new mock instance.
func NewMockUnderFileSystemService(ctrl *gomock.Controller) *MockUnderFileSystemService {
	mock := &MockUnderFileSystemService{ctrl: ctrl}
	mock.recorder = &MockUnderFileSystemServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnderFileSystemService) EXPECT() *MockUnderFileSystemServiceMockRecorder {
	return m.recorder
}

// FreeStorageBytes mocks base method.
func (m *MockUnderFileSystemService) FreeStorageBytes() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FreeStorageBytes")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FreeStorageBytes indicates an expected call of FreeStorageBytes.
func (mr *MockUnderFileSystemServiceMockRecorder) FreeStorageBytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FreeStorageBytes", reflect.TypeOf((*MockUnderFileSystemService)(nil).FreeStorageBytes))
}

// TotalFileNums mocks base method.
func (m *MockUnderFileSystemService) TotalFileNums() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalFileNums")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TotalFileNums indicates an expected call of TotalFileNums.
func (mr *MockUnderFileSystemServiceMockRecorder) TotalFileNums() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalFileNums", reflect.TypeOf((*MockUnderFileSystemService)(nil).TotalFileNums))
}

// TotalStorageBytes mocks base method.
func (m *MockUnderFileSystemService) TotalStorageBytes() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalStorageBytes")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TotalStorageBytes indicates an expected call of TotalStorageBytes.
func (mr *MockUnderFileSystemServiceMockRecorder) TotalStorageBytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalStorageBytes", reflect.TypeOf((*MockUnderFileSystemService)(nil).TotalStorageBytes))
}

// UsedStorageBytes mocks base method.
func (m *MockUnderFileSystemService) UsedStorageBytes() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsedStorageBytes")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UsedStorageBytes indicates an expected call of UsedStorageBytes.
func (mr *MockUnderFileSystemServiceMockRecorder) UsedStorageBytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsedStorageBytes", reflect.TypeOf((*MockUnderFileSystemService)(nil).UsedStorageBytes))
}
