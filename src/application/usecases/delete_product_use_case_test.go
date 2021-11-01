package usecases_test

import (
	"errors"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pedropazello/lojinha-product-catalog-service/mocks"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/usecases"
)

var _ = Describe("DeleteProductUseCase", func() {
	Context("when product is deleted", func() {
		productRepo := mocks.IProductRepository{}
		deleteProductUseCase := usecases.NewDeleteProductUseCase(&productRepo)
		productID := uuid.New().String()

		productRepo.On("Delete", productID).Return(nil)
		err := deleteProductUseCase.Execute(productID)

		It("should not return error", func() {
			Expect(err).To(BeNil())
		})
	})

	Context("when product is not deleted", func() {
		productRepo := mocks.IProductRepository{}
		deleteProductUseCase := usecases.NewDeleteProductUseCase(&productRepo)
		productID := uuid.New().String()
		expectedError := errors.New("error")

		productRepo.On("Delete", productID).Return(expectedError)
		err := deleteProductUseCase.Execute(productID)

		It("should return a error", func() {
			Expect(err).To(Equal(expectedError))
		})
	})
})
