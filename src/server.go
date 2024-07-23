package main

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/grrlopes/go-looptask/src/infra/http/routers"
)

func main() {
	if os.Getenv("MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	app := server.Group("/")

	routers.AuthCtrl(app)
	routers.UserCtrl(app)
	routers.LabelCtrl(app)

	server.Run()
}
