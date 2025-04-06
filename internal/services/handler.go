package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/albertokff/empreendify-saas/internal/database"
)

func CreateService(c *gin.Context) {
	var input struct {
		ClientName  string  `json:"client_name"`
		Description string  `json:"description"`
		Price		float64 `json:"price"`
		Status		string  `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return	
	}

	service := Service{
		ClientName:  input.ClientName,
		Description: input.Description,
		Price:		 input.Price,
		Status: 	 input.Status,
		Date: 		 time.Now(),
	}

	if err := database.DB.Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar serviço"})
		return
	}

	c.JSON(http.StatusCreated, service)
}