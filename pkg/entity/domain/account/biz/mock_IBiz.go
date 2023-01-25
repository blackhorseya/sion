// Code generated by mockery v2.16.0. DO NOT EDIT.

package biz

import (
	contextx "github.com/blackhorseya/irent/pkg/contextx"
	mock "github.com/stretchr/testify/mock"

	model "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
)

// MockIBiz is an autogenerated mock type for the IBiz type
type MockIBiz struct {
	mock.Mock
}

// GetByAccessToken provides a mock function with given fields: ctx, token
func (_m *MockIBiz) GetByAccessToken(ctx contextx.Contextx, token string) (*model.Profile, error) {
	ret := _m.Called(ctx, token)

	var r0 *model.Profile
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string) *model.Profile); ok {
		r0 = rf(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Liveness provides a mock function with given fields: ctx
func (_m *MockIBiz) Liveness(ctx contextx.Contextx) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: ctx, id, password
func (_m *MockIBiz) Login(ctx contextx.Contextx, id string, password string) (*model.Profile, error) {
	ret := _m.Called(ctx, id, password)

	var r0 *model.Profile
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string, string) *model.Profile); ok {
		r0 = rf(ctx, id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string, string) error); ok {
		r1 = rf(ctx, id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Readiness provides a mock function with given fields: ctx
func (_m *MockIBiz) Readiness(ctx contextx.Contextx) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockIBiz interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIBiz creates a new instance of MockIBiz. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIBiz(t mockConstructorTestingTNewMockIBiz) *MockIBiz {
	mock := &MockIBiz{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
