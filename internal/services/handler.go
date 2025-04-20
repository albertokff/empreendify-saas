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

	clientName  := c.Query("client_name")
	status 		:= c.Query("status")
	dateInitial := c.Query("date_initial")
	dateFinal   := c.Query("date_final")

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if clientName != "" {
		query = query.Where("client_name ILIKE ?", "%"+clientName+"%")
	}

	if dateInitial != "" {
		daInicial, err := time.Parse("2006-01-02", dateInitial)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido de Data Inicial"})
			return
		}

		query = query.Where("date >= ?", daInicial)
	}

	if dateFinal != "" {
		daFinal, err := time.Parse("2006-01-02", dateFinal)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido de Data Final"})
			return
		}

		query = query.Where("date <= ?", daFinal)
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

func GetServiceCounts(c *gin.Context) {
	var counts struct {
		Total		int64 `json:"total"`
		Pendente	int64 `json:"pendente"`
		Concluido   int64 `json:"concluido"`
		Andamento	int64 `json:"em andamento"`
	}

	if err := database.DB.Model(&models.Service{}).Count(&counts.Total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao contar serviços!"})
		return
	}

	if err := database.DB.Model(&models.Service{}).Where("status = ?", "pendente").Count(&counts.Pendente).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao contar serviços pendentes"})
		return
	}

	if err := database.DB.Model(&models.Service{}).Where("status = ?", "concluido").Count(&counts.Concluido).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao contar serviços concluídos"})
		return
	}

	if err := database.DB.Model(&models.Service{}).Where("status = ?", "em andamento").Count(&counts.Andamento).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao contar serviços em andamento"})
	}

	c.JSON(http.StatusOK, counts)
}