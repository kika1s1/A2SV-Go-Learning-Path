package main

import (
	"log"
	"net/http"
	"github.com/kika1s1/task_manager/config"
	"github.com/kika1s1/task_manager/router"
)

func main() {
	config.ConnectDB()
	defer config.DisconnectDB()
		// Setup routes
	r := router.SetupRouter()
	// Start the server
	log.Fatal(http.ListenAndServe(":3000", r))
}