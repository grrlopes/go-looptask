package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-looptask/src/helper"
)

func AuthUserToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := helper.ExtractToken(c)
		err := helper.VerifyJwt(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
