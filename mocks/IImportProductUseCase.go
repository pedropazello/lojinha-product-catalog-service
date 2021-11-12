// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	input "github.com/pedropazello/lojinha-product-catalog-service/internal/application/ports/input"

	mock "github.com/stretchr/testify/mock"
)

// IImportProductUseCase is an autogenerated mock type for the IImportProductUseCase type
type IImportProductUseCase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *IImportProductUseCase) Execute(_a0 []input.ProductData) (int, []error) {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func([]input.ProductData) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 []error
	if rf, ok := ret.Get(1).(func([]input.ProductData) []error); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]error)
		}
	}

	return r0, r1
}