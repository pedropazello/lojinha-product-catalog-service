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

var _ = Describe("CreateProductUseCase", func() {
	Describe("Execute", func() {
		Context("when create product", func() {
			productRepo := &mocks.IProductRepository{}
			productUseCase := usecases.NewCreateProductUseCase(productRepo)

			productInputData := input.ProductData{
				Name:        "lp Beth Carvalho No Pagode",
				Description: "Semi novo",
			}

			expectedId := uuid.New().String()
			productDataWithId := entities.Product{
				ID:          expectedId,
				Name:        "lp Beth Carvalho No Pagode",
				Description: "Semi novo",
			}

			productRepo.On("Create", mock.Anything).Return(&productDataWithId, nil)

			fetchedProduct, _ := productUseCase.Execute(&productInputData)

			It("should return a Product with ID", func() {
				Expect(fetchedProduct.ID).To(Equal(expectedId))
			})
		})

		Context("when returns error", func() {
			productRepo := &mocks.IProductRepository{}
			productUseCase := usecases.NewCreateProductUseCase(productRepo)

			productInputData := input.ProductData{
				Name:        "lp Beth Carvalho No Pagode",
				Description: "Semi novo",
			}

			expectedError := errors.New("could not create product")

			productRepo.On("Create", mock.Anything).Return(&entities.Product{}, expectedError)

			_, err := productUseCase.Execute(&productInputData)

			It("should return a error", func() {
				Expect(err).To(Equal(expectedError))
			})
		})
	})
})
