package urlControllers

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rafaelmgr12/Mark-URL/pkg/database"
	"github.com/rafaelmgr12/Mark-URL/pkg/models"
	useCase "github.com/rafaelmgr12/Mark-URL/pkg/useCase/url"
)

type CreateUrlForm struct {
	URL    string `json:"url" binding:"required"`
	UserID string `json:"user_id" gorm:"default:null"`
	Public bool   `json:"public" gorm:"default:false"`
}

func ShowURLs(c *gin.Context) {
	var urls []models.URL
	database.DB.Where("public = ?", true).Find(&urls)
	c.JSON(200, urls)
}

func CreateUrl(c *gin.Context) {
	var input CreateUrlForm

	host := "http://localhost:8080/"
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r, _ := regexp.Compile("^(http|https)://")
	if input.URL == "" || !r.MatchString(input.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	var url models.URL

	id := uuid.New()
	url.ID = id.String()

	url.URL = input.URL
	url.UserID = input.UserID
	if input.UserID == "" || input.UserID == "null" {
		input.UserID = uuid.New().String()
	}
	shortUrl := useCase.GenerateShortLink(input.URL, input.UserID)
	url.Public = input.Public
	url.ShortURL = shortUrl
	url.URL = input.URL

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
