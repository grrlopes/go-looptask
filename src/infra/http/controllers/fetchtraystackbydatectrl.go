package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-looptask/src/application/usecase/fetchtraystackbydate"
	"github.com/grrlopes/go-looptask/src/domain/repository"
	"github.com/grrlopes/go-looptask/src/infra/presenters"
	"github.com/grrlopes/go-looptask/src/infra/repositories/mongodb"
)

var (
	repositoryFetchTrayStack repository.IMongoTrayRepo          = mongodb.NewTrayRepository()
	usecaseFetchTrayStack    fetchtraystackbydate.InputBoundary = fetchtraystackbydate.NewFetchtrayStackByDate(repositoryFetchTrayStack)
)

func FetchTrayStackByDate() gin.HandlerFunc {
	return func(c *gin.Context) {

		result, err := usecaseFetchTrayStack.Execute()

		if err != nil {
			error := presenters.FetchTrayStackByDateError(err)
			c.JSON(http.StatusNotFound, error)
			return
		}

		data := presenters.FetchTrayStackByDateSuccess(result)

		c.JSON(http.StatusOK, data)
	}
}
