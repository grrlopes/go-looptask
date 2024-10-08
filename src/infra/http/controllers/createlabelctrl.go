package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-looptask/src/application/usecase/createlabel"
	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/repository"
	"github.com/grrlopes/go-looptask/src/domain/validator"
	"github.com/grrlopes/go-looptask/src/helper"
	"github.com/grrlopes/go-looptask/src/infra/presenters"
	"github.com/grrlopes/go-looptask/src/infra/repositories/mongodb"
)

var (
	repositoryLabel    repository.IMongoTrayRepo = mongodb.NewTrayRepository()
	usecaseCreateLabel createlabel.InputBoundary = createlabel.NewCreateLabel(repositoryLabel)
)

func CreateLabel() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload entity.Tray
		c.ShouldBindJSON(&payload)

		checked, validErr := validator.Validate(&payload)
		if checked {
			fieldErr := presenters.ValidFieldResponse(validErr)
			c.JSON(http.StatusBadRequest, fieldErr)
			return
		}

		userInfo := helper.GetUserInfoJwt(helper.ExtractToken(c))
		payload.UserId = userInfo.ID

		_, err := usecaseCreateLabel.Execute(&payload)

		if err != nil {
			error := presenters.CreateLabelError(payload)
			c.JSON(http.StatusUnprocessableEntity, error)
			return
		}

		data := presenters.CreateLabelSuccess(payload)

		c.JSON(http.StatusOK, data)
	}
}
