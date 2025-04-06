package main

import (
	"github.com/gin-gonic/gin"
	"github.com/albertokff/empreendify-saas/internal/database"
	"github.com/albertokff/empreendify-saas/internal/services"
)

func main() {
	database.Connect()

	err := database.DB.AutoMigrate(&services.Service{})

	if err != nil {
		panic("Erro ao rodar a migração: " + err.Error())
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/services", services.CreateService)

	r.Run()
}