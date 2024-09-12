package server

import "github.com/gin-gonic/gin"

type App struct {
	Engine *gin.Engine
}

func NewApp() *App {
	engine := gin.Default()
	return &App{engine}
}

func (app *App) Run(port string) {
	err := app.Engine.Run(port)
	if err != nil {
		panic(err)
	}
}
