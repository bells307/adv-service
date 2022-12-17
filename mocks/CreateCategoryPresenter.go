// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/bells307/adv-service/internal/domain"
	mock "github.com/stretchr/testify/mock"

	usecase "github.com/bells307/adv-service/internal/usecase"
)

// CreateCategoryPresenter is an autogenerated mock type for the CreateCategoryPresenter type
type CreateCategoryPresenter struct {
	mock.Mock
}

// Output provides a mock function with given fields: _a0
func (_m *CreateCategoryPresenter) Output(_a0 domain.Category) usecase.CreateCategoryOutput {
	ret := _m.Called(_a0)

	var r0 usecase.CreateCategoryOutput
	if rf, ok := ret.Get(0).(func(domain.Category) usecase.CreateCategoryOutput); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(usecase.CreateCategoryOutput)
	}

	return r0
}

type mockConstructorTestingTNewCreateCategoryPresenter interface {
	mock.TestingT
	Cleanup(func())
}

// NewCreateCategoryPresenter creates a new instance of CreateCategoryPresenter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCreateCategoryPresenter(t mockConstructorTestingTNewCreateCategoryPresenter) *CreateCategoryPresenter {
	mock := &CreateCategoryPresenter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
