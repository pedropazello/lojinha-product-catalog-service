package prompt

import (
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/input"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/output"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/usecases/interfaces"
	"github.com/pedropazello/lojinha-product-catalog-service/src/delivery/prompt/clis"
	"github.com/pedropazello/lojinha-product-catalog-service/src/delivery/serializers"
)

func NewHandler(c interfaces.ICreateProductUseCase, f interfaces.IFindProductUseCase,
	u interfaces.IUpdateProductUseCase, d interfaces.IDeleteProductUseCase,
	cli clis.ICLI, serializer serializers.ISerializer) *Handler {
	return &Handler{
		createProductUseCase: c,
		findProductUseCase:   f,
		updateProductUseCase: u,
		deleteProductUseCase: d,
		cli:                  cli,
		serializer:           serializer,
	}
}

type Handler struct {
	createProductUseCase interfaces.ICreateProductUseCase
	findProductUseCase   interfaces.IFindProductUseCase
	updateProductUseCase interfaces.IUpdateProductUseCase
	deleteProductUseCase interfaces.IDeleteProductUseCase
	StdInput             []string
	StdOutput            string
	cli                  clis.ICLI
	serializer           serializers.ISerializer
}

func (h *Handler) Start() {
	h.StdInput = h.cli.GetStdInput()
	mainCommand := h.StdInput[0]

	switch mainCommand {
	case "create-product":
		h.create()
	case "find-product":
		h.find()
	case "update-product":
		h.update()
	case "delete-product":
		h.delete()
	default:
		h.setStdOutput("Invalid command")
	}
}

func (h *Handler) create() {
	params := input.ProductData{
		Name:        h.StdInput[1],
		Description: h.StdInput[2],
	}

	createdProduct, err := h.createProductUseCase.Execute(&params)

	if err != nil {
		h.setStdOutput(err.Error())
	} else {
		h.putsProduct(createdProduct)
	}
}

func (h *Handler) find() {
	id := h.StdInput[1]

	product, err := h.findProductUseCase.Execute(id)

	if err != nil {
		h.setStdOutput(err.Error())
	} else {
		h.putsProduct(&product)
	}
}

func (h *Handler) update() {
	id := h.StdInput[1]

	params := input.ProductData{
		ID:          id,
		Name:        h.StdInput[2],
		Description: h.StdInput[3],
	}

	updatedProduct, err := h.updateProductUseCase.Execute(&params)

	if err != nil {
		h.setStdOutput(err.Error())
	} else {
		h.putsProduct(&updatedProduct)
	}
}

func (h *Handler) delete() {
	id := h.StdInput[1]

	err := h.deleteProductUseCase.Execute(id)

	if err != nil {
		h.setStdOutput(err.Error())
	} else {
		h.setStdOutput("product deleted")
	}
}

func (h *Handler) putsProduct(product *output.ProductData) {
	serializedProduct, err := h.serializer.Serialize(product)

	if err != nil {
		h.setStdOutput(err.Error())
	}

	h.setStdOutput(string(serializedProduct))
}

func (h *Handler) setStdOutput(msg string) {
	h.StdOutput = msg
	h.cli.PutStdOutput(msg)
}
