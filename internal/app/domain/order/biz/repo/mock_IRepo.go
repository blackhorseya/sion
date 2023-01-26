// Code generated by mockery v2.16.0. DO NOT EDIT.

package repo

import (
	contextx "github.com/blackhorseya/irent/pkg/contextx"
	mock "github.com/stretchr/testify/mock"

	model "github.com/blackhorseya/irent/pkg/entity/domain/account/model"

	ordermodel "github.com/blackhorseya/irent/pkg/entity/domain/order/model"

	rentalmodel "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
)

// MockIRepo is an autogenerated mock type for the IRepo type
type MockIRepo struct {
	mock.Mock
}

// BookCar provides a mock function with given fields: ctx, from, target
func (_m *MockIRepo) BookCar(ctx contextx.Contextx, from *model.Profile, target *rentalmodel.Car) (*ordermodel.Lease, error) {
	ret := _m.Called(ctx, from, target)

	var r0 *ordermodel.Lease
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *model.Profile, *rentalmodel.Car) *ordermodel.Lease); ok {
		r0 = rf(ctx, from, target)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ordermodel.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, *model.Profile, *rentalmodel.Car) error); ok {
		r1 = rf(ctx, from, target)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelBooking provides a mock function with given fields: ctx, from, target
func (_m *MockIRepo) CancelBooking(ctx contextx.Contextx, from *model.Profile, target *ordermodel.Lease) error {
	ret := _m.Called(ctx, from, target)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *model.Profile, *ordermodel.Lease) error); ok {
		r0 = rf(ctx, from, target)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchArrears provides a mock function with given fields: ctx, from, target
func (_m *MockIRepo) FetchArrears(ctx contextx.Contextx, from *model.Profile, target *model.Profile) ([]*ordermodel.ArrearsRecord, error) {
	ret := _m.Called(ctx, from, target)

	var r0 []*ordermodel.ArrearsRecord
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *model.Profile, *model.Profile) []*ordermodel.ArrearsRecord); ok {
		r0 = rf(ctx, from, target)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ordermodel.ArrearsRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, *model.Profile, *model.Profile) error); ok {
		r1 = rf(ctx, from, target)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryBookings provides a mock function with given fields: ctx, from
func (_m *MockIRepo) QueryBookings(ctx contextx.Contextx, from *model.Profile) ([]*ordermodel.Lease, error) {
	ret := _m.Called(ctx, from)

	var r0 []*ordermodel.Lease
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *model.Profile) []*ordermodel.Lease); ok {
		r0 = rf(ctx, from)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ordermodel.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, *model.Profile) error); ok {
		r1 = rf(ctx, from)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockIRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIRepo creates a new instance of MockIRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIRepo(t mockConstructorTestingTNewMockIRepo) *MockIRepo {
	mock := &MockIRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
