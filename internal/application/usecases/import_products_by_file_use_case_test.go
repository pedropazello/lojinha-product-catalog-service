package usecases_test

import (
	"errors"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/usecases"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/infra/parsers"
	"github.com/pedropazello/lojinha-product-catalog-service/mocks"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("ImportProductsByFileUseCase", func() {
	Context("when the file is valid", func() {
		It("should return nil", func() {
			fileStorage := &mocks.IFileStorage{}
			parser := &parsers.ProductsCSVParser{}
			repository := &mocks.IProductRepository{}

			readFileWithProductsUseCase := usecases.NewReadFileWithProductsUseCase(parser)
			importProductsUseCase := usecases.NewImportProductsUseCase(repository)

			usecase := usecases.NewImportProductsByFileUseCase(
				fileStorage,
				readFileWithProductsUseCase,
				importProductsUseCase,
			)

			fileName := "foo.txt"
			mockFile, err := os.ReadFile("../../../fixtures/fake_product_data.csv")
			if err != nil {
				Expect(err).To(BeNil())
			}

			fileStorage.On("GetFile", fileName).Return(mockFile, nil)
			repository.On("Save", mock.Anything).Return(nil)

			err = usecase.Execute(fileName)

			Expect(err).To(BeNil())
		})
	})

	Context("when file doesnt exists", func() {
		It("should return an error", func() {
			fileStorage := &mocks.IFileStorage{}
			parser := &parsers.ProductsCSVParser{}
			repository := &mocks.IProductRepository{}

			readFileWithProductsUseCase := usecases.NewReadFileWithProductsUseCase(parser)
			importProductsUseCase := usecases.NewImportProductsUseCase(repository)

			usecase := usecases.NewImportProductsByFileUseCase(
				fileStorage,
				readFileWithProductsUseCase,
				importProductsUseCase,
			)

			expectedErr := errors.New("foo")
			fileName := "foo.txt"

			fileStorage.On("GetFile", fileName).Return([]byte{}, expectedErr)

			err := usecase.Execute(fileName)

			Expect(err).To(BeEquivalentTo(expectedErr))
		})
	})
})
