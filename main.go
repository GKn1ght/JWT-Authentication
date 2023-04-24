package main

import (
	"jwt-project/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// get the PORT # from env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// enables you to avoid having to pass in http.ResponseWriter & http.Request into each func
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access Granted for api-1"})
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access Granted for api-2"})
	})

	router.Run(":" + port)
}
