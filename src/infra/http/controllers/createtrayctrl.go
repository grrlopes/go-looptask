package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-looptask/src/application/usecase/createlabeltray"
	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/repository"
	"github.com/grrlopes/go-looptask/src/domain/validator"
	"github.com/grrlopes/go-looptask/src/helper"
	"github.com/grrlopes/go-looptask/src/infra/presenters"
	"github.com/grrlopes/go-looptask/src/infra/repositories/mongodb"
)

var (
	repositoryLabelTray    repository.IMongoTrayRepo     = mongodb.NewTrayRepository()
	usecaseCreateLabelTray createlabeltray.InputBoundary = createlabeltray.NewListAllTrays(repositoryLabelTray)
)

func CreateTrayStack() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload entity.LabelTrayStack
		c.ShouldBindJSON(&payload)

		checked, validErr := validator.Validate(&payload)
		if checked {
			fieldErr := presenters.ValidFieldResponse(validErr)
			c.JSON(http.StatusBadRequest, fieldErr)
			return
		}


		userInfo := helper.GetUserInfoJwt(helper.ExtractToken(c))
		payload.Owner = userInfo.ID

		result, err := usecaseCreateLabelTray.Execute(&payload)

		if err != nil {
			error := presenters.CreateLabelTrayStackError(payload)
			c.JSON(http.StatusUnprocessableEntity, error)
			return
		}

		data := presenters.CreateLabelTrayStackSuccess(payload, result)

		c.JSON(http.StatusOK, data)
	}
}
