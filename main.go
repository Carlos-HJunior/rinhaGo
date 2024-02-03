package main

import (
	"github.com/gin-gonic/gin"
	"rinhaGo/controller"
	"rinhaGo/repository"
)

func main() {
	r := gin.Default()

	personRepository := repository.NewPersonRepository()
	controller.ConfigApi(r, personRepository)

	err := r.Run()
	if err != nil {
		return
	}
}
