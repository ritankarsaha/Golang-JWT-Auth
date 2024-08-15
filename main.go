package main

import (
	routes "github.com/ritankarsaha/Golang-JWT-Auth/routes"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

    routes.AuthRoutes(router)
	routes.UserRoutes(router)

	

	if err := router.Run(":" + port); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}