// Code generated by mockery v2.14.0. DO NOT EDIT.

package wireguard

import (
	wireguardv1 "github.com/clly/wireguard-cni/gen/wgcni/wireguard/v1"
	mock "github.com/stretchr/testify/mock"
)

// MockPeers is an autogenerated mock type for the Peers type
type MockPeers struct {
	mock.Mock
}

// ListPeers provides a mock function with given fields:
func (_m *MockPeers) ListPeers() ([]*wireguardv1.Peer, error) {
	ret := _m.Called()

	var r0 []*wireguardv1.Peer
	if rf, ok := ret.Get(0).(func() []*wireguardv1.Peer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*wireguardv1.Peer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockPeers interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPeers creates a new instance of MockPeers. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPeers(t mockConstructorTestingTNewMockPeers) *MockPeers {
	mock := &MockPeers{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
