// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	usecase "github.com/bells307/adv-service/internal/usecase"
	mock "github.com/stretchr/testify/mock"
)

// FindAllAdvertismentSummaryUseCase is an autogenerated mock type for the FindAllAdvertismentSummaryUseCase type
type FindAllAdvertismentSummaryUseCase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *FindAllAdvertismentSummaryUseCase) Execute(_a0 context.Context, _a1 usecase.FindAllAdvertismentSummaryInput) (usecase.FindAllAdvertismentSummaryOutput, error) {
	ret := _m.Called(_a0, _a1)

	var r0 usecase.FindAllAdvertismentSummaryOutput
	if rf, ok := ret.Get(0).(func(context.Context, usecase.FindAllAdvertismentSummaryInput) usecase.FindAllAdvertismentSummaryOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(usecase.FindAllAdvertismentSummaryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, usecase.FindAllAdvertismentSummaryInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFindAllAdvertismentSummaryUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewFindAllAdvertismentSummaryUseCase creates a new instance of FindAllAdvertismentSummaryUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFindAllAdvertismentSummaryUseCase(t mockConstructorTestingTNewFindAllAdvertismentSummaryUseCase) *FindAllAdvertismentSummaryUseCase {
	mock := &FindAllAdvertismentSummaryUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
