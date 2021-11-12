package parsers_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/infra/parsers"
)

var _ = Describe("Products CSV Parser Parse", func() {
	It("convert all lines of file into products", func() {
		parser := parsers.ProductsCSVParser{}

		fileToParser, err := os.ReadFile("../../../fixtures/fake_product_data.csv")
		Expect(err).ToNot(HaveOccurred())

		products, err := parser.Parse(fileToParser)

		Expect(products).To(HaveLen(2))
		Expect(err).To(BeNil())
	})
})