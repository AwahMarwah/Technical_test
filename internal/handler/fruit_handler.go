package handlers

import (
	"net/http"

	"github.com/AwahMarwah/Technical_test/service"
	"github.com/gin-gonic/gin"
)

type FruitHandler struct {
	service *services.FruitService
}

func NewFruitHandler(service *services.FruitService) *FruitHandler {
	return &FruitHandler{service: service}
}

func (h *FruitHandler) GetAllFruitNames(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"fruitNames": h.service.GetAllFruitNames()})
}

func (h *FruitHandler) SeparateFruitsByType(c *gin.Context) {
	c.JSON(http.StatusOK, h.service.SeparateFruitsByType())
}

func (h *FruitHandler) GetTotalStockByType(c *gin.Context) {
	c.JSON(http.StatusOK, h.service.GetTotalStockByType())
}

func (h *FruitHandler) Comments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"comments": h.service.Comments()})
}
