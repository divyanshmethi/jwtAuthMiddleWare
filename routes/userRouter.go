package routes

import (
	"github.com/divyansh/golang-jwt/controllers"
	"github.com/divyansh/golang-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	//These are protected routes and we need to verify the token before allowing access to these endpoints
	router.Use(middleware.Authenticate())
	router.GET("/users", controllers.GetUsers())
	router.GET("/users/:user_id", controllers.GetUser())
}
