package handler

import (
	"froztypoint/backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) List(c *gin.Context) {
	products, err := h.service.List() // panggil service
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {

}

func (h *ProductHandler) GetProductByID(c *gin.Context) {

}
