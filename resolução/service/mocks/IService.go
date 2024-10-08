// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	models "desafio/models"

	mock "github.com/stretchr/testify/mock"
)

// IService is an autogenerated mock type for the IService type
type IService struct {
	mock.Mock
}

// ProcessGpsData provides a mock function with given fields: data
func (_m *IService) ProcessGpsData(data models.GpsRequest) error {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for ProcessGpsData")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.GpsRequest) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessGyroscopeData provides a mock function with given fields: data
func (_m *IService) ProcessGyroscopeData(data models.GyroscopeRequest) error {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for ProcessGyroscopeData")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.GyroscopeRequest) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessPhotoData provides a mock function with given fields: data
func (_m *IService) ProcessPhotoData(data models.PhotoRequest) error {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for ProcessPhotoData")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.PhotoRequest) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIService creates a new instance of IService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IService {
	mock := &IService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
