package controllers

import (
	"net/http"
	"regexp"

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
	host := "http://localhost:8080/"
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r, _ := regexp.Compile("^(http|https)://")
	if url.URL == "" || !r.MatchString(url.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	id := uuid.New()
	url.ID = id.String()
	shortUrl := useCase.GenerateShortLink(url.URL, url.ID)
	url.ShortURL = shortUrl
	database.DB.Create(&url)
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}
func HandleShortUrlRedirect(c *gin.Context) {
	var url models.URL
	shortUrl := c.Param("short_url")
	database.DB.Where("short_url = ?", shortUrl).First(&url)
	if url.ShortURL == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	url.Hits++
	database.DB.Save(&url)
	c.Redirect(302, url.URL)
}

func RoutesNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
}
