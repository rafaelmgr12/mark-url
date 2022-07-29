package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelmgr12/Mark-URL/controllers"
)

func HandleRequest() {
	router := gin.Default()
	router.GET("/url", controllers.ShowURLs)
	router.POST("/url", controllers.CreateUrl)
	router.NoRoute(controllers.RoutesNotFound)
	router.Run(":8080")

}
