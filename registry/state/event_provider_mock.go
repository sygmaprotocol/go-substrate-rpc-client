// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package state

import (
	types "github.com/sygmaprotocol/go-substrate-rpc-client/v4/types"
	mock "github.com/stretchr/testify/mock"
)

// EventProviderMock is an autogenerated mock type for the EventProvider type
type EventProviderMock struct {
	mock.Mock
}

// GetStorageEvents provides a mock function with given fields: meta, blockHash
func (_m *EventProviderMock) GetStorageEvents(meta *types.Metadata, blockHash types.Hash) (*types.StorageDataRaw, error) {
	ret := _m.Called(meta, blockHash)

	var r0 *types.StorageDataRaw
	if rf, ok := ret.Get(0).(func(*types.Metadata, types.Hash) *types.StorageDataRaw); ok {
		r0 = rf(meta, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.StorageDataRaw)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.Metadata, types.Hash) error); ok {
		r1 = rf(meta, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewEventProviderMockT interface {
	mock.TestingT
	Cleanup(func())
}

// NewEventProviderMock creates a new instance of EventProviderMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventProviderMock(t NewEventProviderMockT) *EventProviderMock {
	mock := &EventProviderMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
