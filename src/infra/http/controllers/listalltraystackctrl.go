package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	listalltraystack "github.com/grrlopes/go-looptask/src/application/usecase/listalltrayStack"
	"github.com/grrlopes/go-looptask/src/domain/repository"
	"github.com/grrlopes/go-looptask/src/infra/presenters"
	"github.com/grrlopes/go-looptask/src/infra/repositories/mongodb"
)

var (
	repositoryListAllTrayStack repository.IMongoTrayRepo      = mongodb.NewTrayRepository()
	usecaseListAllTrayStack    listalltraystack.InputBoundary = listalltraystack.NewListAllTrayStack(repositoryListAllTrayStack)
)

func ListAllTrayStack() gin.HandlerFunc {
	return func(c *gin.Context) {

		result, err := usecaseListAllTrayStack.Execute()

		if err != nil {
			error := presenters.ListAllLabelTrayStackError(err)
			c.JSON(http.StatusNotFound, error)
			return
		}

		data := presenters.ListAllLabelTrayStackSuccess(result)

		c.JSON(http.StatusOK, data)
	}
}
