package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rafaelmgr12/Mark-URL/pkg/controllers"
)

func HandleRequest() {
	router := gin.Default()
	router.GET("/url", controllers.ShowURLs)
	router.POST("/url", controllers.CreateUrl)
	router.GET("/:short_url", controllers.HandleShortUrlRedirect)
	router.NoRoute(controllers.RoutesNotFound)
	router.Use(cors.Default())
	router.Run(":8080")

}
