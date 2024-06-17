package controller

import (
	"context"
	"fmt"
	"time"

	"gihub.com/Uttkarsh-raj/Dist-Cache/connection"
	"github.com/gin-gonic/gin"
)

func AddNewNode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), time.Second*4)
		defer cancel()

		newNode := connection.NewNode(ctx.Request.RemoteAddr)
		fmt.Printf("%s\n", newNode.ID)
		println(newNode.Addr)
	}
}
