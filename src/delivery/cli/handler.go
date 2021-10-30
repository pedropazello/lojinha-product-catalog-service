package cli

import (
	"github.com/google/uuid"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/input"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/usecases/interfaces"
	"github.com/pedropazello/lojinha-product-catalog-service/src/delivery/cli/clis"
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
		h.StdOutput = "Invalid command"
	}
}

func (h *Handler) create() {
	params := input.ProductData{
		Name:        h.StdInput[1],
		Description: h.StdInput[2],
	}

	createdProduct, err := h.createProductUseCase.Execute(&params)

	if err != nil {
		h.StdOutput = err.Error()
	} else {
		serializedProduct, err := h.serializer.Serialize(createdProduct)

		if err != nil {
			h.StdOutput = err.Error()
		}

		h.StdOutput = string(serializedProduct)
	}
}

func (h *Handler) find() {
	id, err := uuid.Parse(h.StdInput[1])

	if err != nil {
		h.StdOutput = err.Error()
	}

	product, err := h.findProductUseCase.Execute(id)

	if err != nil {
		h.StdOutput = err.Error()
	} else {
		serializedProduct, err := h.serializer.Serialize(product)

		if err != nil {
			h.StdOutput = err.Error()
		}

		h.StdOutput = string(serializedProduct)
	}
}

func (h *Handler) update() {
	id, err := uuid.Parse(h.StdInput[1])
	if err != nil {
		h.StdOutput = err.Error()
	}

	params := input.ProductData{
		ID:          id,
		Name:        h.StdInput[2],
		Description: h.StdInput[3],
	}

	updatedProduct, err := h.updateProductUseCase.Execute(&params)

	if err != nil {
		h.StdOutput = err.Error()
	} else {
		serializedProduct, err := h.serializer.Serialize(updatedProduct)

		if err != nil {
			h.StdOutput = err.Error()
		}

		h.StdOutput = string(serializedProduct)
	}
}

func (h *Handler) delete() {
	id, err := uuid.Parse(h.StdInput[1])

	if err != nil {
		h.StdOutput = err.Error()
	}

	err = h.deleteProductUseCase.Execute(id)

	if err != nil {
		h.StdOutput = err.Error()
	} else {
		h.StdOutput = "product deleted"
	}
}
