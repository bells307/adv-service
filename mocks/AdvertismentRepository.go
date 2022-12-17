// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bells307/adv-service/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// AdvertismentRepository is an autogenerated mock type for the AdvertismentRepository type
type AdvertismentRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, adv
func (_m *AdvertismentRepository) Create(ctx context.Context, adv domain.Advertisment) error {
	ret := _m.Called(ctx, adv)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Advertisment) error); ok {
		r0 = rf(ctx, adv)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, limit, offset
func (_m *AdvertismentRepository) Find(ctx context.Context, limit uint, offset uint) ([]domain.Advertisment, error) {
	ret := _m.Called(ctx, limit, offset)

	var r0 []domain.Advertisment
	if rf, ok := ret.Get(0).(func(context.Context, uint, uint) []domain.Advertisment); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Advertisment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, uint) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, id
func (_m *AdvertismentRepository) FindByID(ctx context.Context, id string) (domain.Advertisment, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Advertisment
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Advertisment); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Advertisment)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAdvertismentRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewAdvertismentRepository creates a new instance of AdvertismentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAdvertismentRepository(t mockConstructorTestingTNewAdvertismentRepository) *AdvertismentRepository {
	mock := &AdvertismentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}