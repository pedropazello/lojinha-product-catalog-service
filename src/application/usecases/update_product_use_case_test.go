package usecases_test

import (
	"errors"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pedropazello/lojinha-product-catalog-service/mocks"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/input"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/usecases"
	"github.com/pedropazello/lojinha-product-catalog-service/src/domain/entities"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("UpdateProductUseCase", func() {
	Context("When updating a product", func() {
		productRepo := mocks.IProductRepository{}
		updateProductUseCase := usecases.NewUpdateProductUseCase(&productRepo)
		updatedProductUUID := uuid.New()

		inputProduct := input.ProductData{
			ID:          updatedProductUUID,
			Name:        "Updated Product",
			Description: "This is an updated product",
		}

		originalProductData := entities.Product{
			ID:          updatedProductUUID,
			Name:        "Original Product",
			Description: "This is an original product",
		}

		productRepo.On("GetById", mock.Anything).Return(&originalProductData, nil)
		productRepo.On("Save", mock.Anything).Return(nil)

		outputProduct, err := updateProductUseCase.Execute(&inputProduct)

		It("should update the product", func() {
			Expect(err).To(BeNil())
			Expect(outputProduct.ID).To(Equal(updatedProductUUID))
			Expect(outputProduct.Name).To(Equal("Updated Product"))
			Expect(outputProduct.Description).To(Equal("This is an updated product"))
		})
	})

	Context("when the product is not updated", func() {
		productRepo := mocks.IProductRepository{}
		updateProductUseCase := usecases.NewUpdateProductUseCase(&productRepo)
		updatedProductUUID := uuid.New()
		expectedError := errors.New("error updating product")

		inputProduct := input.ProductData{
			ID:   updatedProductUUID,
			Name: "Updated Product",
		}

		originalProductData := entities.Product{
			ID:          updatedProductUUID,
			Name:        "Original Product",
			Description: "This is an original product",
		}

		productRepo.On("GetById", mock.Anything).Return(&originalProductData, expectedError)
		productRepo.On("Save", mock.Anything).Return(nil)

		_, err := updateProductUseCase.Execute(&inputProduct)

		It("should return the error", func() {
			Expect(err).To(Equal(expectedError))
		})
	})
})
