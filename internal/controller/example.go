package controller

import (
	"casbin/internal/midware"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {

	c.JSON(200, "hello world")
}

func RegisterBooks(router *gin.RouterGroup) {
	router.GET("/books", midware.Auth.RequiresPermissions([]string{"book:read"}), GetBooks)
}
