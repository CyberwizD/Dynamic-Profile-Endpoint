package routes

import (
	"github.com/CyberwizD/Dynamic-Profile-Endpoint/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/me", handlers.GetProfile)
}
