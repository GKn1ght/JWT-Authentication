package routes

import (
	controller "jwt-project/controllers"
	"jwt-project/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	//! use middleware to insure that both of these routes are protected routes
	// should not be able to use the Users Routes without a token
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
