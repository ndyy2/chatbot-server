// routes/api.go
package routes

import (
	"ai-assistant/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterAPIRoutes(e *echo.Echo, aiCtrl *controllers.AIController) {
	api := e.Group("/api/v1")
	
	api.POST("/chat", aiCtrl.ChatHandler)
	//api.POST("/system/settings", aiCtrl.UpdateSystemSettings)
	//api.GET("/session/:id", aiCtrl.GetSession)
}