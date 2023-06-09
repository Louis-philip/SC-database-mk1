// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dbmodel "github.com/milvus-io/milvus/internal/metastore/db/dbmodel"
	mock "github.com/stretchr/testify/mock"
)

// IIndexDb is an autogenerated mock type for the IIndexDb type
type IIndexDb struct {
	mock.Mock
}

// Get provides a mock function with given fields: tenantID, collectionID
func (_m *IIndexDb) Get(tenantID string, collectionID int64) ([]*dbmodel.Index, error) {
	ret := _m.Called(tenantID, collectionID)

	var r0 []*dbmodel.Index
	if rf, ok := ret.Get(0).(func(string, int64) []*dbmodel.Index); ok {
		r0 = rf(tenantID, collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbmodel.Index)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(tenantID, collectionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: in
func (_m *IIndexDb) Insert(in []*dbmodel.Index) error {
	ret := _m.Called(in)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*dbmodel.Index) error); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: tenantID
func (_m *IIndexDb) List(tenantID string) ([]*dbmodel.IndexResult, error) {
	ret := _m.Called(tenantID)

	var r0 []*dbmodel.IndexResult
	if rf, ok := ret.Get(0).(func(string) []*dbmodel.IndexResult); ok {
		r0 = rf(tenantID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbmodel.IndexResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tenantID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarkDeletedByCollectionID provides a mock function with given fields: tenantID, collID
func (_m *IIndexDb) MarkDeletedByCollectionID(tenantID string, collID int64) error {
	ret := _m.Called(tenantID, collID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64) error); ok {
		r0 = rf(tenantID, collID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkDeletedByIndexID provides a mock function with given fields: tenantID, idxID
func (_m *IIndexDb) MarkDeletedByIndexID(tenantID string, idxID int64) error {
	ret := _m.Called(tenantID, idxID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64) error); ok {
		r0 = rf(tenantID, idxID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: in
func (_m *IIndexDb) Update(in *dbmodel.Index) error {
	ret := _m.Called(in)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dbmodel.Index) error); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIIndexDb interface {
	mock.TestingT
	Cleanup(func())
}

// NewIIndexDb creates a new instance of IIndexDb. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIIndexDb(t mockConstructorTestingTNewIIndexDb) *IIndexDb {
	mock := &IIndexDb{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
