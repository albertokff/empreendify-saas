package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/albertokff/empreendify-saas/internal/services"
)

func RegisterServiceRoutes(r *gin.Engine) {
	servicesGroup := r.Group("/services")

	servicesGroup.GET("/", services.GetServices)
	servicesGroup.POST("/", services.CreateService)
	servicesGroup.GET("/service-counts", services.GetServiceCounts)
	servicesGroup.GET("/:id", services.GetServiceById)
	servicesGroup.PUT("/:id", services.UpdateService)
	servicesGroup.DELETE("/:id", services.DeleteService)
}