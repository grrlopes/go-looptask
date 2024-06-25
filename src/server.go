package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-looptask/src/infra/http/routers"
)

func main() {
	if os.Getenv("MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	app := server.Group("/")

	routers.AuthCtrl(app)
	routers.UserCtrl(app)
	routers.LabelCtrl(app)

	server.Run()
}
