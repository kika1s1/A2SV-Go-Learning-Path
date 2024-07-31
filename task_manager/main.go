package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager/router"
)

func main() {
	r := gin.Default()
	router.SetRoutes(r)
	r.Run(":3000") 
}