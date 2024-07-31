package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"example/web-service-gin/data"
	"example/web-service-gin/models"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Albums)
}

func GetAlbum(c *gin.Context) {
	id := c.Param("id")
	for _, album := range data.Albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PostAlbums(c *gin.Context) {
	var newAlbum models.Album
	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}
	data.Albums = append(data.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
