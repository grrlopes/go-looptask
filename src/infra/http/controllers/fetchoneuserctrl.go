package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-looptask/src/application/usecase/fetchuserbyid"
	"github.com/grrlopes/go-looptask/src/domain/repository"
	"github.com/grrlopes/go-looptask/src/helper"
	"github.com/grrlopes/go-looptask/src/infra/presenters"
	"github.com/grrlopes/go-looptask/src/infra/repositories/mongodb"
)

var (
	repositoryFetchOneUser repository.IMongoUserRepo   = mongodb.NewUserRepository()
	usecaseFetchOneUser    fetchuserbyid.InputBoundary = fetchuserbyid.NewFetchOneUser(repositoryFetchOneUser)
)

func FetchOneUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo := helper.GetUserInfoJwt(helper.ExtractToken(c))

		result, err := usecaseFetchOneUser.Execute(&userInfo)

		if err != nil {
			error := presenters.FetchOneUserError(userInfo)
			c.JSON(http.StatusNotFound, error)
			return
		}

		data := presenters.FetchOneUserSuccess(result)

		c.JSON(http.StatusOK, data)
	}
}
