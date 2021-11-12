// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	entities "github.com/pedropazello/lojinha-product-catalog-service/internal/domain/entities"
	mock "github.com/stretchr/testify/mock"
)

// IProductsFileParser is an autogenerated mock type for the IProductsFileParser type
type IProductsFileParser struct {
	mock.Mock
}

// Parse provides a mock function with given fields: file
func (_m *IProductsFileParser) Parse(file []byte) ([]entities.Product, error) {
	ret := _m.Called(file)

	var r0 []entities.Product
	if rf, ok := ret.Get(0).(func([]byte) []entities.Product); ok {
		r0 = rf(file)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
