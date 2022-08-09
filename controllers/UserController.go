package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelmgr12/Mark-URL/database"
	"github.com/rafaelmgr12/Mark-URL/models"
)

func showUser(c *gin.Context) {
	var user models.User
	database.DB.Find(&user)
	c.JSON(200, user)
}
