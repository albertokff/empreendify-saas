package main

import (
	"github.com/gin-gonic/gin"
	"github.com/albertokff/empreendify-saas/internal/database"
	"github.com/albertokff/empreendify-saas/internal/models"
	"github.com/albertokff/empreendify-saas/internal/services"
)

func main() {
	database.Connect() // Abrindo a conexão com o banco

	err := database.DB.AutoMigrate(&models.Service{}) // Criando as migrações

	if err != nil {
		panic("Erro ao rodar a migração: " + err.Error())
	}

	r := gin.Default() // Aqui estou iniciando o roteador, e já aplicando 2 middlewares: logs e recovery.

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/services", services.CreateService)

	r.GET("/services", services.GetServices)

	r.GET("/services/:id", services.GetServiceById)

	r.POST("/services/:id", services.UpdateService)

	r.DELETE("/services/:id", services.DeleteService)

	r.Run()
}