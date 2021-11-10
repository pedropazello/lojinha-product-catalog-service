package usecases_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/ports/input"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/usecases"
	"github.com/pedropazello/lojinha-product-catalog-service/mocks"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("ImportProductsUseCase Execute", func() {
	Context("When any product are invalid", func() {
		It("Should return error message for this product", func() {
			repository := &mocks.IProductRepository{}
			usecase := usecases.NewImportProductsUseCase(repository)

			product1 := input.ProductData{
				Name:        "Product 1",
				Description: "Description 1",
			}

			product2 := input.ProductData{
				Name:        "Product 2",
				Description: "Description 2",
			}

			product3 := input.ProductData{
				Name:        "Product 3",
				Description: "Description 3",
			}

			productListInput := []input.ProductData{}
			productListInput = append(productListInput, product1)
			productListInput = append(productListInput, product2)
			productListInput = append(productListInput, product3)

			errorMsg := errors.New("fail")

			repository.On("Save", mock.Anything).Return(errorMsg)

			productsImported, errors := usecase.Execute(productListInput)

			Expect(productsImported).To(BeEquivalentTo(0))
			Expect(len(errors)).To(BeEquivalentTo(3))
			Expect(errors[0]).To(BeEquivalentTo(errorMsg))
		})
	})

	Context("When all products are valid", func() {
		It("import all products", func() {
			repository := &mocks.IProductRepository{}
			usecase := usecases.NewImportProductsUseCase(repository)

			product1 := input.ProductData{
				Name:        "Product 1",
				Description: "Description 1",
			}

			product2 := input.ProductData{
				Name:        "Product 2",
				Description: "Description 2",
			}

			product3 := input.ProductData{
				Name:        "Product 3",
				Description: "Description 3",
			}

			productListInput := []input.ProductData{}
			productListInput = append(productListInput, product1)
			productListInput = append(productListInput, product2)
			productListInput = append(productListInput, product3)

			repository.On("Save", mock.Anything).Return(nil)

			productsImported, errors := usecase.Execute(productListInput)

			Expect(productsImported).To(BeEquivalentTo(3))
			Expect(len(errors)).To(BeEquivalentTo(0))
		})
	})
})
