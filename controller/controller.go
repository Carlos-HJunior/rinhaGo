package controller

import (
	"github.com/gin-gonic/gin"
	"rinhaGo/repository"
)

func ConfigApi(r *gin.Engine, repository repository.PersonRepository) {
	var controller = rootController{
		repository: repository,
	}

	r.POST("/pessoas", controller.createPerson)
}

type rootController struct {
	repository repository.PersonRepository
}
