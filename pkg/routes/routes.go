package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	urlControllers "github.com/rafaelmgr12/Mark-URL/pkg/controllers/url"
	userControllers "github.com/rafaelmgr12/Mark-URL/pkg/controllers/user"
)

func HandleRequest() {
	router := gin.Default()
	router.GET("/url", urlControllers.ShowURLs)
	router.POST("/url", urlControllers.CreateUrl)
	router.GET("/:short_url", urlControllers.HandleShortUrlRedirect)
	router.NoRoute(urlControllers.RoutesNotFound)

	router.GET("/user", userControllers.ShowUser)
	router.POST("/user", userControllers.Signup)

	router.Use(cors.Default())
	router.Run(":8080")

}
