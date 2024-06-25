package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-looptask/src/helper"
	"github.com/grrlopes/go-looptask/src/infra/presenters"
)

func AuthUserToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := helper.ExtractToken(c)
		err := helper.VerifyJwt(tokenString)
		if err != nil {
      authFailed := presenters.AuthError()
			c.JSON(http.StatusUnauthorized, authFailed)
			c.Abort()
			return
		}
		c.Next()
	}
}
