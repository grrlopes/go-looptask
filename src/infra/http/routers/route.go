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
	app.GET("/fetchoneuser", middleware.AuthUserToken(), controllers.FetchOneUser())
}

func LabelCtrl(app gin.IRouter) {
	app.POST("/createlabelstack", middleware.AuthUserToken(), controllers.CreateTrayStack())
	app.POST("/createlabeled", middleware.AuthUserToken(), controllers.CreateLabel())
	app.GET("/fetchonelabel", middleware.AuthUserToken(), controllers.FetchOneTray())
	app.GET("/listalltraystack", controllers.ListAllTrayStack())
	app.POST("/fetchtraystackbydate", middleware.AuthUserToken(), controllers.FetchTrayStackByDate())
}
