package handler

import (
	"log"
	"net/http"
	"product-service/api/v1/product"

	"github.com/gin-gonic/gin"
)

// dependencies
type ProductHandler struct {
	logger         *log.Logger
	productService product.Service
}

func NewProductHandler(logger *log.Logger, productService product.Service) *ProductHandler {
	return &ProductHandler{logger, productService}
}

func (h *ProductHandler) Greetings(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"message": "helo from product service"})

	products, err := h.productService.GetAllProduct()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})

}
