package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-looptask/src/domain/entity"
	validate "github.com/grrlopes/go-looptask/src/domain/validator"
	"github.com/grrlopes/go-looptask/src/helper"
	"github.com/grrlopes/go-looptask/src/infra/presenters"
)

func ValidateJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenize entity.ValidateJwt
		authHeader := strings.Split(c.GetHeader("Authorization"), " ")

		if len(authHeader) != 2 || authHeader[0] != "Bearer" {
			error := presenters.HeaderFailed()
			c.JSON(http.StatusBadRequest, error)
			return
		}

		tokenize.Token = authHeader[1]

		checked, validErr := validate.Validate(&tokenize)
		if checked {
			fieldErr := presenters.JwtValidField(validErr)
			c.JSON(http.StatusBadRequest, fieldErr)
			return
		}

		err := helper.VerifyJwt(tokenize.Token)
		if err != nil {
			error := presenters.JwtError(tokenize)
			c.JSON(http.StatusUnauthorized, error)
			return
		}

		data := presenters.JwtSuccess(tokenize.Token)

		c.JSON(http.StatusOK, data)
	}
}
