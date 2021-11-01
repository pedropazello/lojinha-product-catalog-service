// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	input "github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/input"

	mock "github.com/stretchr/testify/mock"

	output "github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/output"
)

// IUpdateProductUseCase is an autogenerated mock type for the IUpdateProductUseCase type
type IUpdateProductUseCase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: inputProduct
func (_m *IUpdateProductUseCase) Execute(inputProduct *input.ProductData) (output.ProductData, error) {
	ret := _m.Called(inputProduct)

	var r0 output.ProductData
	if rf, ok := ret.Get(0).(func(*input.ProductData) output.ProductData); ok {
		r0 = rf(inputProduct)
	} else {
		r0 = ret.Get(0).(output.ProductData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*input.ProductData) error); ok {
		r1 = rf(inputProduct)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}