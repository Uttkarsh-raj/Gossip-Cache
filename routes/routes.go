package routes

import (
	"gihub.com/Uttkarsh-raj/Dist-Cache/controller"
	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) {
	router.GET("/", controller.AddNewNode())
}
