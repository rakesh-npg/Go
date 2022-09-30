package controller

import "github.com/gin-gonic/gin"

func UserController(ctx *gin.Context) {
	ctx.String(200, "Hello World")
}
