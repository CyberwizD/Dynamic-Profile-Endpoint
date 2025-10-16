package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/CyberwizD/Dynamic-Profile-Endpoint/api/models"
	"github.com/CyberwizD/Dynamic-Profile-Endpoint/api/services"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	fact, err := services.GetCatFact()
	if err != nil {
		log.Printf("could not get cat fact: %v", err)
		fact = "Could not retrieve a cat fact."
	}

	response := models.ProfileResponse{
		Status: "success",
		User: models.User{
			Email: "wisdomce19@gmail.com",
			Name:  "Wisdom Udoye",
			Stack: "Go/Gin",
		},
		Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
		Fact:      fact,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}
