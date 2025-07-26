package routes

import (
	"github.com/gin-gonic/gin"
	controller "jwt/controllers"
	"jwt/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users",controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}