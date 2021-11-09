package prompt_test

import (
	"errors"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/usecases"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/delivery/prompt"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/domain/entities"
	"github.com/pedropazello/lojinha-product-catalog-service/mocks"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Handler", func() {
	Describe("Handle", func() {
		Context("when the command is 'create-product'", func() {
			Context("when the product is valid", func() {
				repository := &mocks.IProductRepository{}

				createProductUseCase := usecases.NewCreateProductUseCase(repository)
				findProductUseCase := usecases.NewFindProductUseCase(repository)
				updateProductUseCase := usecases.NewUpdateProductUseCase(repository)
				deleteProductUseCase := usecases.NewDeleteProductUseCase(repository)
				cliMock := &mocks.ICLI{}
				serializer := &mocks.ISerializer{}

				handler := prompt.NewHandler(
					createProductUseCase,
					findProductUseCase,
					updateProductUseCase,
					deleteProductUseCase,
					cliMock,
					serializer,
				)

				productCreated := entities.Product{
					ID:          uuid.New().String(),
					Name:        "name",
					Description: "description",
				}

				outputProduct := "{\"ID\":\"00000000-0000-0000-0000-000000000000\",\"Name\":\"name\",\"Description\":\"description\"}"

				cliMock.On("GetStdInput").Return([]string{"create-product", "name", "description"})
				cliMock.On("PutStdOutput", mock.Anything).Return(outputProduct)
				serializer.On("Serialize", mock.Anything).Return([]byte(outputProduct), nil)
				repository.On("Create", mock.Anything).Return(&productCreated, nil)

				handler.Start()
				output := handler.StdOutput

				It("should create a new product", func() {
					Expect(output).To(Equal(outputProduct))
				})
			})

			Context("when the product is invalid", func() {
				repository := &mocks.IProductRepository{}

				createProductUseCase := usecases.NewCreateProductUseCase(repository)
				findProductUseCase := usecases.NewFindProductUseCase(repository)
				updateProductUseCase := usecases.NewUpdateProductUseCase(repository)
				deleteProductUseCase := usecases.NewDeleteProductUseCase(repository)
				cliMock := &mocks.ICLI{}
				serializer := &mocks.ISerializer{}

				handler := prompt.NewHandler(
					createProductUseCase,
					findProductUseCase,
					updateProductUseCase,
					deleteProductUseCase,
					cliMock,
					serializer,
				)

				expectedError := errors.New("could not create product")

				cliMock.On("GetStdInput").Return([]string{"create-product", "", ""})
				cliMock.On("PutStdOutput", mock.Anything).Return("could not create product")
				repository.On("Create", mock.Anything).Return(&entities.Product{}, expectedError)

				handler.Start()

				It("should not create a new product", func() {
					Expect(handler.StdOutput).To(Equal("could not create product"))
				})
			})
		})

		Context("when the command is 'find-product'", func() {
			Context("when the product is found", func() {
				repository := &mocks.IProductRepository{}

				createProductUseCase := usecases.NewCreateProductUseCase(repository)
				findProductUseCase := usecases.NewFindProductUseCase(repository)
				updateProductUseCase := usecases.NewUpdateProductUseCase(repository)
				deleteProductUseCase := usecases.NewDeleteProductUseCase(repository)
				cliMock := &mocks.ICLI{}
				serializer := &mocks.ISerializer{}

				handler := prompt.NewHandler(
					createProductUseCase,
					findProductUseCase,
					updateProductUseCase,
					deleteProductUseCase,
					cliMock,
					serializer,
				)

				productId := uuid.New().String()

				product := entities.Product{
					ID:          productId,
					Name:        "name",
					Description: "description",
				}

				outputProduct := "{\"ID\":\"00000000-0000-0000-0000-000000000000\",\"Name\":\"name\",\"Description\":\"description\"}"

				cliMock.On("GetStdInput").Return([]string{"find-product", "00000000-0000-0000-0000-000000000000"})
				cliMock.On("PutStdOutput", mock.Anything).Return(outputProduct)
				serializer.On("Serialize", mock.Anything).Return([]byte(outputProduct), nil)
				repository.On("GetById", mock.Anything).Return(&product, nil)

				handler.Start()
				output := handler.StdOutput

				It("should find a product", func() {
					Expect(output).To(Equal(outputProduct))
				})
			})

			Context("when the product is not found", func() {
				repository := &mocks.IProductRepository{}

				createProductUseCase := usecases.NewCreateProductUseCase(repository)
				findProductUseCase := usecases.NewFindProductUseCase(repository)
				updateProductUseCase := usecases.NewUpdateProductUseCase(repository)
				deleteProductUseCase := usecases.NewDeleteProductUseCase(repository)
				cliMock := &mocks.ICLI{}
				serializer := &mocks.ISerializer{}

				handler := prompt.NewHandler(
					createProductUseCase,
					findProductUseCase,
					updateProductUseCase,
					deleteProductUseCase,
					cliMock,
					serializer,
				)

				cliMock.On("GetStdInput").Return([]string{"find-product", "00000000-0000-0000-0000-000000000000"})
				cliMock.On("PutStdOutput", mock.Anything).Return("could not find product")
				repository.On("GetById", mock.Anything).Return(&entities.Product{}, errors.New("could not find product"))

				handler.Start()

				It("should not find a product", func() {
					Expect(handler.StdOutput).To(Equal("could not find product"))
				})
			})
		})

		Context("when the command is 'update-product'", func() {
			Context("when the product is updated", func() {
				repository := &mocks.IProductRepository{}

				createProductUseCase := usecases.NewCreateProductUseCase(repository)
				findProductUseCase := usecases.NewFindProductUseCase(repository)
				updateProductUseCase := usecases.NewUpdateProductUseCase(repository)
				deleteProductUseCase := usecases.NewDeleteProductUseCase(repository)
				cliMock := &mocks.ICLI{}
				serializer := &mocks.ISerializer{}

				handler := prompt.NewHandler(
					createProductUseCase,
					findProductUseCase,
					updateProductUseCase,
					deleteProductUseCase,
					cliMock,
					serializer,
				)

				originalProduct := entities.Product{
					ID:          uuid.New().String(),
					Name:        "name",
					Description: "description",
				}

				outputProduct := "{\"ID\":\"00000000-0000-0000-0000-000000000000\",\"Name\":\"name\",\"Description\":\"description\"}"

				cliMock.On("GetStdInput").Return([]string{"update-product", "00000000-0000-0000-0000-000000000000", "name", "description"})
				cliMock.On("PutStdOutput", mock.Anything).Return(outputProduct)
				serializer.On("Serialize", mock.Anything).Return([]byte(outputProduct), nil)
				repository.On("GetById", mock.Anything).Return(&originalProduct, nil)
				repository.On("Save", mock.Anything).Return(nil)

				handler.Start()
				output := handler.StdOutput

				It("should update a product", func() {
					Expect(output).To(Equal(outputProduct))
				})
			})

			Context("when the product is not updated", func() {
				repository := &mocks.IProductRepository{}

				createProductUseCase := usecases.NewCreateProductUseCase(repository)
				findProductUseCase := usecases.NewFindProductUseCase(repository)
				updateProductUseCase := usecases.NewUpdateProductUseCase(repository)
				deleteProductUseCase := usecases.NewDeleteProductUseCase(repository)
				cliMock := &mocks.ICLI{}
				serializer := &mocks.ISerializer{}

				handler := prompt.NewHandler(
					createProductUseCase,
					findProductUseCase,
					updateProductUseCase,
					deleteProductUseCase,
					cliMock,
					serializer,
				)

				cliMock.On("GetStdInput").Return([]string{"update-product", "00000000-0000-0000-0000-000000000000", "name", "description"})
				cliMock.On("PutStdOutput", mock.Anything).Return("could not update product")
				repository.On("GetById", mock.Anything).Return(&entities.Product{}, errors.New("could not update product"))

				handler.Start()

				It("should not update a product", func() {
					Expect(handler.StdOutput).To(Equal("could not update product"))
				})
			})
		})

		Context("when the command is 'delete-product'", func() {
			Context("when the product is deleted", func() {
				repository := &mocks.IProductRepository{}

				createProductUseCase := usecases.NewCreateProductUseCase(repository)
				findProductUseCase := usecases.NewFindProductUseCase(repository)
				updateProductUseCase := usecases.NewUpdateProductUseCase(repository)
				deleteProductUseCase := usecases.NewDeleteProductUseCase(repository)
				cliMock := &mocks.ICLI{}
				serializer := &mocks.ISerializer{}

				handler := prompt.NewHandler(
					createProductUseCase,
					findProductUseCase,
					updateProductUseCase,
					deleteProductUseCase,
					cliMock,
					serializer,
				)

				cliMock.On("GetStdInput").Return([]string{"delete-product", "00000000-0000-0000-0000-000000000000"})
				cliMock.On("PutStdOutput", mock.Anything).Return("product deleted")
				repository.On("Delete", mock.Anything).Return(nil)

				handler.Start()

				It("should delete a product", func() {
					Expect(handler.StdOutput).To(Equal("product deleted"))
				})
			})

			Context("when the product is not deleted", func() {
				repository := &mocks.IProductRepository{}

				createProductUseCase := usecases.NewCreateProductUseCase(repository)
				findProductUseCase := usecases.NewFindProductUseCase(repository)
				updateProductUseCase := usecases.NewUpdateProductUseCase(repository)
				deleteProductUseCase := usecases.NewDeleteProductUseCase(repository)
				cliMock := &mocks.ICLI{}
				serializer := &mocks.ISerializer{}

				handler := prompt.NewHandler(
					createProductUseCase,
					findProductUseCase,
					updateProductUseCase,
					deleteProductUseCase,
					cliMock,
					serializer,
				)

				cliMock.On("GetStdInput").Return([]string{"delete-product", "00000000-0000-0000-0000-000000000000"})
				cliMock.On("PutStdOutput", mock.Anything).Return("could not delete product")
				repository.On("Delete", mock.Anything).Return(errors.New("could not delete product"))

				handler.Start()

				It("should not delete a product", func() {
					Expect(handler.StdOutput).To(Equal("could not delete product"))
				})
			})
		})
	})
})
