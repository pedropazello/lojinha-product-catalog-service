package usecases_test

import (
	"errors"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pedropazello/lojinha-product-catalog-service/mocks"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/usecases"
	"github.com/pedropazello/lojinha-product-catalog-service/src/domain/entities"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("FindProductUseCase", func() {
	Describe("Execute", func() {
		Context("when find product by ID", func() {
			productRepo := mocks.IProductRepository{}
			findProductUseCase := usecases.NewFindProductUseCase(&productRepo)
			expectedUUID := uuid.New()

			productDataWithId := entities.Product{
				ID:          expectedUUID,
				Name:        "lp Beth Carvalho No Pagode",
				Description: "Semi novo",
			}

			productRepo.On("GetById", mock.Anything).Return(&productDataWithId, nil)

			product, err := findProductUseCase.Execute(expectedUUID)

			It("should return product", func() {
				Expect(err).To(BeNil())
				Expect(product.ID).To(Equal(expectedUUID))
			})
		})

		Context("when product does not exists", func() {
			productRepo := mocks.IProductRepository{}
			findProductUseCase := usecases.NewFindProductUseCase(&productRepo)
			expectedUUID := uuid.Nil

			expectedError := errors.New("product not found")

			productDataWithoutId := entities.Product{
				ID: expectedUUID,
			}

			productRepo.On("GetById", mock.Anything).Return(&productDataWithoutId, expectedError)

			product, err := findProductUseCase.Execute(expectedUUID)

			It("should return product", func() {
				Expect(err).To(Equal(expectedError))
				Expect(product.ID).To(Equal(expectedUUID))
			})
		})
	})
})
