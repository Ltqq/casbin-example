package main

import (
	"casbin/internal/controller"
	"casbin/internal/midware"
	"casbin/internal/server"
)

func main() {
	app := server.NewApp()
	err := midware.InitCasbinAuth()
	if err != nil {
		panic(err)
	}
	controller.RegisterBooks(app.Engine.Group("/v1"))
	controller.RegisterUser(app.Engine.Group("/v1"))

	app.Run(":8080")
}
