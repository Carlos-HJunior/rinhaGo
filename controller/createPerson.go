package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"rinhaGo/entities"
)

func (rc rootController) createPerson(c *gin.Context) {
	var newPerson entities.Person
	err := c.BindJSON(&newPerson)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("invalid request"))
		return
	}

	err = validateCreatePerson(newPerson)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id, err := rc.repository.CreatePerson(newPerson)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.Header("location", fmt.Sprintf("/pessoas/%d", id))
	c.Status(http.StatusCreated)
}

func validateCreatePerson(person entities.Person) error {
	if person.Nickname == "" || len(person.Nickname) > 32 {
		return errors.New("invalid nickname")
	}

	if person.Name == "" || len(person.Name) > 100 {
		return errors.New("invalid name")
	}

	re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	var match = re.FindStringSubmatch(person.Birth)
	if len(match) != 1 {
		return errors.New("invalid birth")
	}

	if person.Stack != nil {
		for _, s := range person.Stack {
			if len(s) > 32 {
				return errors.New("invalid stack")
			}
		}
	}

	return nil
}
