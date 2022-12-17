// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	usecase "github.com/bells307/adv-service/internal/usecase"
	mock "github.com/stretchr/testify/mock"
)

// FindAdvertismentUseCase is an autogenerated mock type for the FindAdvertismentUseCase type
type FindAdvertismentUseCase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *FindAdvertismentUseCase) Execute(_a0 context.Context, _a1 usecase.FindAdvertismentInput) (usecase.FindAdvertismentOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 usecase.FindAdvertismentOutput
	if rf, ok := ret.Get(0).(func(context.Context, usecase.FindAdvertismentInput) usecase.FindAdvertismentOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(usecase.FindAdvertismentOutput)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, usecase.FindAdvertismentInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFindAdvertismentUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewFindAdvertismentUseCase creates a new instance of FindAdvertismentUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFindAdvertismentUseCase(t mockConstructorTestingTNewFindAdvertismentUseCase) *FindAdvertismentUseCase {
	mock := &FindAdvertismentUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}