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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repositoryFetchOneTray repository.IMongoTrayRepo   = mongodb.NewTrayRepository()
	usecaseFetchOneTray    fetchtraybyid.InputBoundary = fetchtraybyid.NewFetchOneTray(repositoryLabelTray)
)

func FetchOneTray() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload entity.TrayId
    var param entity.LabelId
		err := c.ShouldBind(&param)

    payload.Id, _ = primitive.ObjectIDFromHex(param.Id)

		checked, validErr := validator.Validate(&payload)
		if checked {
			fieldErr := presenters.ValidFieldResponse(validErr)
			c.JSON(http.StatusBadRequest, fieldErr)
			return
		}

		result, err := usecaseFetchOneTray.Execute(&payload)

		if err != nil {
			error := presenters.FetchOneLabelTrayError(err)
			c.JSON(http.StatusNotFound, error)
			return
		}

		data := presenters.FetchOneLabelTraySuccess(result)

		c.JSON(http.StatusOK, data)
	}
}
