package usecases_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/usecases"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/infra/parsers"
)

var _ = Describe("ReadFileWithProductsUseCase Execute", func() {
	It("convert all lines of file into products", func() {
		parser := parsers.ProductsCSVParser{}
		useCase := usecases.NewReadFileWithProductsUseCase(&parser)

		fileToParser, err := os.ReadFile("../../../fixtures/fake_product_data.csv")
		Expect(err).ToNot(HaveOccurred())

		products, err := useCase.Execute(fileToParser)

		Expect(products).To(HaveLen(2))
		Expect(err).To(BeNil())
	})
})
