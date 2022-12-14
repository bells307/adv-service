// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/bells307/adv-service/internal/domain"
	mock "github.com/stretchr/testify/mock"

	usecase "github.com/bells307/adv-service/internal/usecase"
)

// CreateAdvertismentPresenter is an autogenerated mock type for the CreateAdvertismentPresenter type
type CreateAdvertismentPresenter struct {
	mock.Mock
}

// Output provides a mock function with given fields: _a0
func (_m *CreateAdvertismentPresenter) Output(_a0 domain.Advertisment) usecase.CreateAdvertismentOutput {
	ret := _m.Called(_a0)

	var r0 usecase.CreateAdvertismentOutput
	if rf, ok := ret.Get(0).(func(domain.Advertisment) usecase.CreateAdvertismentOutput); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(usecase.CreateAdvertismentOutput)
	}

	return r0
}

type mockConstructorTestingTNewCreateAdvertismentPresenter interface {
	mock.TestingT
	Cleanup(func())
}

// NewCreateAdvertismentPresenter creates a new instance of CreateAdvertismentPresenter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCreateAdvertismentPresenter(t mockConstructorTestingTNewCreateAdvertismentPresenter) *CreateAdvertismentPresenter {
	mock := &CreateAdvertismentPresenter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
