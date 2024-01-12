package routes

import (
	"PracticeGDB/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/get", controller.GetUser)
	router.GET("/get/:id", controller.GetOnlyUser)
	router.POST("/", controller.CreateUser)
	router.PUT("/:id", controller.UpdateUser)
	router.DELETE("/:id", controller.DeleteUser)
}
