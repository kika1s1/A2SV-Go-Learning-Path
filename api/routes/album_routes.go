package routes

import (
	"github.com/gin-gonic/gin"
	"example/web-service-gin/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbum)
	router.POST("/albums/create_album", controllers.PostAlbums)
}
