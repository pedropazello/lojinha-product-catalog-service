// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	output "github.com/pedropazello/lojinha-product-catalog-service/internal/application/ports/output"
	mock "github.com/stretchr/testify/mock"
)

// IFindProductUseCase is an autogenerated mock type for the IFindProductUseCase type
type IFindProductUseCase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: productId
func (_m *IFindProductUseCase) Execute(productId string) (output.ProductData, error) {
	ret := _m.Called(productId)

	var r0 output.ProductData
	if rf, ok := ret.Get(0).(func(string) output.ProductData); ok {
		r0 = rf(productId)
	} else {
		r0 = ret.Get(0).(output.ProductData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(productId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
