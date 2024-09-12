package controller

import (
	"casbin/internal/midware"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
)

var jwtKey = jwt.NewHS256([]byte("jwtkey"))

func Login(c *gin.Context) {
	username := c.Query("name")
	token, err := midware.GetToken(username)
	if err != nil {
		c.JSON(500, "some error")
		return
	}

	c.JSON(200, token)
}

func RegisterUser(router *gin.RouterGroup) {
	router.GET("/login", Login)
}
