// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package parser

import (
	registry "github.com/sygmaprotocol/go-substrate-rpc-client/v4/registry"
	types "github.com/sygmaprotocol/go-substrate-rpc-client/v4/types"
	mock "github.com/stretchr/testify/mock"
)

// EventParserMock is an autogenerated mock type for the EventParser type
type EventParserMock struct {
	mock.Mock
}

// ParseEvents provides a mock function with given fields: eventRegistry, sd
func (_m *EventParserMock) ParseEvents(eventRegistry registry.EventRegistry, sd *types.StorageDataRaw) ([]*Event, error) {
	ret := _m.Called(eventRegistry, sd)

	var r0 []*Event
	if rf, ok := ret.Get(0).(func(registry.EventRegistry, *types.StorageDataRaw) []*Event); ok {
		r0 = rf(eventRegistry, sd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(registry.EventRegistry, *types.StorageDataRaw) error); ok {
		r1 = rf(eventRegistry, sd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewEventParserMockT interface {
	mock.TestingT
	Cleanup(func())
}

// NewEventParserMock creates a new instance of EventParserMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventParserMock(t NewEventParserMockT) *EventParserMock {
	mock := &EventParserMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
