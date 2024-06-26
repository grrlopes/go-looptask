package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-looptask/src/application/usecase/fetchtraybyid"
	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/repository"
	"github.com/grrlopes/go-looptask/src/domain/validator"
	"github.com/grrlopes/go-looptask/src/infra/presenters"
	"github.com/grrlopes/go-looptask/src/infra/repositories/mongodb"
)

var (
	repositoryFetchOneTray repository.IMongoTrayRepo   = mongodb.NewTrayRepository()
	usecaseFetchOneTray    fetchtraybyid.InputBoundary = fetchtraybyid.NewFetchOneTray(repositoryLabelTray)
)

func FetchOneTray() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload entity.Tray
		err := c.ShouldBindJSON(&payload)

		checked, validErr := validator.Validate(&payload)
		if checked {
			fieldErr := presenters.ValidFieldResponse(validErr)
			c.JSON(http.StatusBadRequest, fieldErr)
			return
		}

		result, err := usecaseFetchOneTray.Execute(&payload)

		if err != nil {
			error := presenters.FetchOneLabelTrayError(payload, err)
			c.JSON(http.StatusNotFound, error)
			return
		}

		data := presenters.FetchOneLabelTraySuccess(result)

		c.JSON(http.StatusOK, data)
	}
}
