package cli_test

import (
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pedropazello/lojinha-product-catalog-service/mocks"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/usecases"
	"github.com/pedropazello/lojinha-product-catalog-service/src/delivery/cli"
	"github.com/pedropazello/lojinha-product-catalog-service/src/domain/entities"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Handler", func() {
	Describe("Handle", func() {
		Context("when the command is 'create-product'", func() {
			repository := &mocks.IProductRepository{}

			createProductUseCase := usecases.NewCreateProductUseCase(repository)
			findProductUseCase := usecases.NewFindProductUseCase(repository)
			updateProductUseCase := usecases.NewUpdateProductUseCase(repository)
			deleteProductUseCase := usecases.NewDeleteProductUseCase(repository)
			cliMock := &mocks.ICLI{}
			serializer := &mocks.ISerializer{}

			handler := cli.NewHandler(
				createProductUseCase,
				findProductUseCase,
				updateProductUseCase,
				deleteProductUseCase,
				cliMock,
				serializer,
			)

			productCreated := entities.Product{
				ID:          uuid.UUID{},
				Name:        "name",
				Description: "description",
			}

			outputProduct := "{\"ID\":\"00000000-0000-0000-0000-000000000000\",\"Name\":\"name\",\"Description\":\"description\"}"

			cliMock.On("GetStdInput").Return([]string{"create-product", "name", "description"})
			serializer.On("Serialize", mock.Anything).Return([]byte(outputProduct), nil)
			repository.On("Create", mock.Anything).Return(&productCreated, nil)

			handler.Start()
			output := handler.StdOutput

			It("should create a new product", func() {
				Expect(output).To(Equal(outputProduct))
			})
		})

		Context("when the command is 'find-product'", func() {
			repository := &mocks.IProductRepository{}

			createProductUseCase := usecases.NewCreateProductUseCase(repository)
			findProductUseCase := usecases.NewFindProductUseCase(repository)
			updateProductUseCase := usecases.NewUpdateProductUseCase(repository)
			deleteProductUseCase := usecases.NewDeleteProductUseCase(repository)
			cliMock := &mocks.ICLI{}
			serializer := &mocks.ISerializer{}

			handler := cli.NewHandler(
				createProductUseCase,
				findProductUseCase,
				updateProductUseCase,
				deleteProductUseCase,
				cliMock,
				serializer,
			)

			productId := uuid.UUID{}

			product := entities.Product{
				ID:          productId,
				Name:        "name",
				Description: "description",
			}

			outputProduct := "{\"ID\":\"00000000-0000-0000-0000-000000000000\",\"Name\":\"name\",\"Description\":\"description\"}"

			cliMock.On("GetStdInput").Return([]string{"find-product", "00000000-0000-0000-0000-000000000000"})
			serializer.On("Serialize", mock.Anything).Return([]byte(outputProduct), nil)
			repository.On("GetById", mock.Anything).Return(&product, nil)

			handler.Start()
			output := handler.StdOutput

			It("should find a product", func() {
				Expect(output).To(Equal(outputProduct))
			})
		})

		Context("when the command is 'update-product'", func() {
			repository := &mocks.IProductRepository{}

			createProductUseCase := usecases.NewCreateProductUseCase(repository)
			findProductUseCase := usecases.NewFindProductUseCase(repository)
			updateProductUseCase := usecases.NewUpdateProductUseCase(repository)
			deleteProductUseCase := usecases.NewDeleteProductUseCase(repository)
			cliMock := &mocks.ICLI{}
			serializer := &mocks.ISerializer{}

			handler := cli.NewHandler(
				createProductUseCase,
				findProductUseCase,
				updateProductUseCase,
				deleteProductUseCase,
				cliMock,
				serializer,
			)

			originalProduct := entities.Product{
				ID:          uuid.UUID{},
				Name:        "name",
				Description: "description",
			}

			outputProduct := "{\"ID\":\"00000000-0000-0000-0000-000000000000\",\"Name\":\"name\",\"Description\":\"description\"}"

			cliMock.On("GetStdInput").Return([]string{"update-product", "00000000-0000-0000-0000-000000000000", "name", "description"})
			serializer.On("Serialize", mock.Anything).Return([]byte(outputProduct), nil)
			repository.On("GetById", mock.Anything).Return(&originalProduct, nil)
			repository.On("Save", mock.Anything).Return(nil)

			handler.Start()
			output := handler.StdOutput

			It("should update a product", func() {
				Expect(output).To(Equal(outputProduct))
			})
		})

		Context("when the command is 'delete-product'", func() {
			repository := &mocks.IProductRepository{}

			createProductUseCase := usecases.NewCreateProductUseCase(repository)
			findProductUseCase := usecases.NewFindProductUseCase(repository)
			updateProductUseCase := usecases.NewUpdateProductUseCase(repository)
			deleteProductUseCase := usecases.NewDeleteProductUseCase(repository)
			cliMock := &mocks.ICLI{}
			serializer := &mocks.ISerializer{}

			handler := cli.NewHandler(
				createProductUseCase,
				findProductUseCase,
				updateProductUseCase,
				deleteProductUseCase,
				cliMock,
				serializer,
			)

			cliMock.On("GetStdInput").Return([]string{"delete-product", "00000000-0000-0000-0000-000000000000"})
			repository.On("Delete", mock.Anything).Return(nil)

			handler.Start()

			It("should delete a product", func() {
				Expect(handler.StdOutput).To(Equal("product deleted"))
			})
		})
	})
})
