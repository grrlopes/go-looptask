package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-looptask/src/application/usecase/login"
	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/repository"
	validate "github.com/grrlopes/go-looptask/src/domain/validator"
	"github.com/grrlopes/go-looptask/src/infra/presenters"
	"github.com/grrlopes/go-looptask/src/infra/repositories/mongodb"
)

var (
	repositoryLogin repository.IMongoUserRepo = mongodb.NewUserRepository()
	usecaseLogin    login.InputBoundary       = login.NewLogin(repositoryLogin)
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload entity.Users
		err := c.ShouldBindJSON(&payload)

		checked, validErr := validate.Validate(&payload)
		if checked {
			fieldErr := presenters.LoginValidField(validErr)
			c.JSON(http.StatusBadRequest, fieldErr)
			return
		}

		result, err := usecaseLogin.Execute(&payload)

		if err != nil {
			error := presenters.LoginError(payload)
			c.JSON(http.StatusInternalServerError, error)
			return
		}

		data := presenters.LoginSuccess(result)

		c.JSON(http.StatusOK, data)
	}
}
