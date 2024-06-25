package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-looptask/src/infra/http/controllers"
	"github.com/grrlopes/go-looptask/src/middleware"
)

func AuthCtrl(app gin.IRouter) {
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "No valid endpoint provided!",
		})
	})

	app.POST("/login", controllers.Login())
}

func UserCtrl(app gin.IRouter) {
	app.POST("/createuser", controllers.CreateUser())
}

func LabelCtrl(app gin.IRouter) {
	app.POST("/createlabel", controllers.CreateLabelTray())
	app.POST("/fetchonelabel", middleware.AuthUserToken(), controllers.FetchOneTray())
}
