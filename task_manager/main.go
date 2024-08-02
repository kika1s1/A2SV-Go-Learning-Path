package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager/router"
	"github.com/kika1s1/task_manager/config"
)

func main() {
	config.ConnectDB()
	defer config.DisconnectDB()
	r := gin.Default()
	router.SetRoutes(r)
	r.Run(":3000") 
}