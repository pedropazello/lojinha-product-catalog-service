// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// IDeleteProductUseCase is an autogenerated mock type for the IDeleteProductUseCase type
type IDeleteProductUseCase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: id
func (_m *IDeleteProductUseCase) Execute(id uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
