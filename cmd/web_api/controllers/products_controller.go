package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/input"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/usecases/interfaces"
)

type productsController struct {
	findProductUseCase   interfaces.IFindProductUseCase
	createProductUseCase interfaces.ICreateProductUseCase
}

func NewProductsController(findProductUseCase interfaces.IFindProductUseCase,
	createProductUseCase interfaces.ICreateProductUseCase) *productsController {
	return &productsController{
		findProductUseCase:   findProductUseCase,
		createProductUseCase: createProductUseCase,
	}
}

func (controller *productsController) Create(c *gin.Context) {
	productInput := input.ProductData{}
	if err := c.ShouldBindJSON(&productInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		output, err := controller.createProductUseCase.Execute(&productInput)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusCreated, output)
	}
}

func (controller *productsController) Show(c *gin.Context) {
	id := c.Param("id")

	output, err := controller.findProductUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, output)
	}
}
