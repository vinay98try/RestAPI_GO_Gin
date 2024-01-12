package main

import (
	"PracticeGDB/config"
	"PracticeGDB/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)

	router.Run(":8005")
}
