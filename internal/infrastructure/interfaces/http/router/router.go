package router

import (
	"github.com/andino-hq/andino/internal/infrastructure/config"
	"github.com/andino-hq/andino/internal/infrastructure/interfaces/http/controllers"
	"github.com/andino-hq/andino/internal/infrastructure/interfaces/http/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Setup(cfg *config.Config, logger *zap.Logger) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.Logger(logger))

	healthController := controllers.NewHealthController()

	// Basic routes
	r.GET("/", healthController.HelloWorld)
	r.GET("/health", healthController.Health)
	r.GET("/hello", healthController.HelloWorld)

	// API v1 routes
	api := r.Group("/api/v1")
	{
		api.GET("/health", healthController.Health)
		api.GET("/hello", healthController.HelloWorld)

		// Future DTE routes
		dte := api.Group("/dte")
		{
			dte.GET("/", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Endpoints DTE - Próximamente 📄",
					"info":    "Documentos Tributarios Electrónicos",
				})
			})
		}

		// Future Payment routes
		payments := api.Group("/payments")
		{
			payments.GET("/", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Endpoints de Pagos - Próximamente 💳",
					"info":    "Integración con sistemas de pago chilenos",
				})
			})
		}

		// Future SII routes
		sii := api.Group("/sii")
		{
			sii.GET("/", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Endpoints SII - Próximamente 🏛️",
					"info":    "Integración con Servicio de Impuestos Internos",
				})
			})
		}
	}
	return r
}
