package main

import (
	"github.com/gin-gonic/gin"
	"example/web-service-gin/routes"
)

func main() {
	router := gin.Default()
	routes.Routes(router)
	router.Run(":3000")
}
