package routes

import (
	"github.com/divyansh/golang-jwt/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	//These are not protected routes because user does not have a token at the time he tries to access these
	router.POST("/users/signup", controllers.Signup())
	router.POST("/users/login", controllers.Login())
}
