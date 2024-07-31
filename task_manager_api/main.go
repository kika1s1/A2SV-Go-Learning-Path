package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager_api/routes"
)

func main() {
	router := gin.Default()
	routes.Routes(router)
	router.Run(":3000")
}