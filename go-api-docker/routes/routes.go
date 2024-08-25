package routes

import (
	"go-api-docker/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	r.Use(cors.New(config))
	api := r.Group("/api")
	{
		api.GET("/schedule", handler.Get)
		api.POST("/schedule", handler.Create)
		api.PATCH("/schedule/:id", handler.Update)
		api.DELETE("/schedule/:id", handler.Delete)
	}
	health := r.Group("/")
	{
		health.GET("/health", handler.HealthCheck)
	}
	return r
}
