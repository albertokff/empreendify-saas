package main

import (
	"time"

	"github.com/albertokff/empreendify-saas/internal/controllers"
	"github.com/albertokff/empreendify-saas/internal/database"
	"github.com/albertokff/empreendify-saas/internal/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect() // Abrindo a conexão com o banco

	err := database.DB.AutoMigrate(&models.Service{}) // Criando as migrações

	if err != nil {
		panic("Erro ao rodar a migração: " + err.Error())
	}

	r := gin.Default() // Aqui estou iniciando o roteador, e já aplicando 2 middlewares: logs e recovery.

	r.Use(cors.New(cors.Config{
		AllowOrigins:	[]string{"http://localhost:5173"},
		AllowMethods:	[]string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:	[]string{"Origin", "Content-Type"},
		AllowCredentials: true,
		MaxAge:			  12 * time.Hour,
	}))

	controllers.RegisterServiceRoutes(r)

	r.Run()
}