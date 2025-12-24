package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rqrniii/DevOps-Microservices/services/ai-service/internal/service"
)

type AIHandler struct {
	service *service.AIService
}

func NewAIHandler(s *service.AIService) *AIHandler {
	return &AIHandler{service: s}
}

func (h *AIHandler) Generate(c *gin.Context) {
	var req struct {
		Prompt string `json:"prompt"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.GenerateTask(c.Request.Context(), req.Prompt)
	if err != nil {
		log.Printf("ERROR: AI generation failed: %v", err) // Add this
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	lines := strings.Split(result, "\n")

	tasks := []string{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			tasks = append(tasks, line)
		}
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
