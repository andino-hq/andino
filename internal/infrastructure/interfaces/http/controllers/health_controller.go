package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

type HealthResponse struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

func (h *HealthController) Health(c *gin.Context) {
	response := HealthResponse{
		Status:    "ok",
		Service:   "andino",
		Version:   "1.0.0",
		Timestamp: time.Now(),
		Message:   "Hello World from Andino! 🏔️",
	}

	c.JSON(http.StatusOK, response)
}

func (h *HealthController) HelloWorld(c *gin.Context) {
	name := c.DefaultQuery("name", "World")

	response := gin.H{
		"message":   "Hello " + name + " from Andino! 🇨🇱",
		"service":   "andino",
		"timestamp": time.Now(),
		"info":      "Fiscal, tax, and accounting management system for Chile",
	}

	c.JSON(http.StatusOK, response)
}
