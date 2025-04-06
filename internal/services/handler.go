package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/albertokff/empreendify-saas/internal/database"
	"github.com/albertokff/empreendify-saas/internal/models"
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

	service := models.Service{
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

func GetServices(c *gin.Context) {
	query := database.DB

	clientName := c.Query("client_name")
	status := c.Query("status")

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if clientName != "" {
		query = query.Where("client_name ILIKE ?", "%"+clientName+"%")
	}

	var services []models.Service

	if err := query.Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar serviços!"})
		return
	}

	c.JSON(http.StatusOK, services)
}

func GetServiceById(c *gin.Context) {
	id := c.Param("id")

	var service models.Service

	if err := database.DB.Find(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Serviço não encontrado!"})
		return
	}

	c.JSON(http.StatusOK, service)
}

func UpdateService(c *gin.Context) {
	id := c.Param("id")

	var service models.Service

	if err := database.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Serviço não encontrado!"})
		return
	}

	var input struct {
		ClientName  *string  `json:"client_name"`
		Description *string  `json:"description"`	
		Price 		*float64 `json:"price"`
		Status 		*string  `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON Inválido!"})
		return
	}

	if input.ClientName != nil {
		service.ClientName = *input.ClientName
	}

	if input.Description != nil {
		service.Description = *input.Description
	}

	if input.Price != nil {
		service.Price = *input.Price
	}

	if input.Status != nil {
		service.Status = *input.Status
	}

	if err := database.DB.Save(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o Serviço!"})
		return
	}

	c.JSON(http.StatusOK, service)
}

func DeleteService(c *gin.Context) {
	id := c.Param("id")

	var service models.Service

	if err := database.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Serviço não encontrado"})
		return
	}

	if err := database.DB.Delete(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar o serviço"})
		return 
	}

	c.Status(http.StatusNoContent)
}