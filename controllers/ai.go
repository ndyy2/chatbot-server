// controllers/ai.go
package controllers

import (
	"ai-assistant/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AIController struct {
	aiService *services.AIService
}

func NewAIController(aiService *services.AIService) *AIController {
	return &AIController{aiService: aiService}
}

func (c *AIController) ChatHandler(ctx echo.Context) error {
	type Request struct {
		SessionID string `json:"session_id"`
		Message   string `json:"message"`
	}

	var req Request
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	response, err := c.aiService.ProcessMessage(req.SessionID, req.Message)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"response": response})
}