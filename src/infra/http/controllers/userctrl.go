package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-looptask/src/application/usecase/usersave"
	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/repository"
	"github.com/grrlopes/go-looptask/src/domain/validator"
	"github.com/grrlopes/go-looptask/src/infra/presenters"
	"github.com/grrlopes/go-looptask/src/infra/repositories/mongodb"
)

var (
	repositoryUser    repository.IMongoUserRepo = mongodb.NewUserRepository()
	usecaseCreateUser usersave.InputBoundary    = usersave.NewUserSave(repositoryUser)
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload entity.Users
		err := c.ShouldBindJSON(&payload)

		checked, validErr := validator.Validate(&payload)
		if checked {
			fieldErr := presenters.ValidFieldResponse(validErr)
			c.JSON(http.StatusBadRequest, fieldErr)
			return
		}

		result, err := usecaseCreateUser.Execute(&payload)

		if err != nil {
			error := presenters.ErrorResponse(result)
			c.JSON(http.StatusInternalServerError, error)
			return
		}

		data := presenters.SuccessResponse(result)

		c.JSON(http.StatusOK, data)
	}
}
