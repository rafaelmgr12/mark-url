package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rafaelmgr12/Mark-URL/database"
	"github.com/rafaelmgr12/Mark-URL/models"
	"github.com/rafaelmgr12/Mark-URL/useCase"
)

func ShowURLs(c *gin.Context) {
	var urls []models.URL
	database.DB.Find(&urls)
	c.JSON(200, urls)
}

func CreateUrl(c *gin.Context) {
	var url models.URL
	host := "http://mark-url.com/"
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := uuid.New()
	url.ID = id.String()
	shortUrl := useCase.GenerateShortLink(url.URL, url.ID)
	url.ShortURL = host + shortUrl
	database.DB.Create(&url)
	c.JSON(http.StatusOK, url)
}

func RoutesNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
}
