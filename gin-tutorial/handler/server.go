package handler

import (
	"github.com/gin-gonic/gin"
	"go-tutorial/gin-tutorial/model"
	"go-tutorial/gin-tutorial/repository"
	"go-tutorial/gin-tutorial/request"
	"go-tutorial/gin-tutorial/response"
	"log"
	"net/http"
)

type ServerHandler struct {
	personRepo repository.PersonRepo
}

func NewServerHandler(personRepo repository.PersonRepo) *ServerHandler {
	return &ServerHandler{
		personRepo: personRepo,
	}
}

// @Summary getPerson
// @Tags Person
// @Param person-id path int true "person-id"
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Success 200 {object} response.Person
// @Router /api/v1/getPerson/{person-id} [get]
func (m *ServerHandler) GetPersonWithPathHandler(c *gin.Context) {
	var personRequestUri request.PersonRequestUri
	var httpStatus int
	var responseBody interface{}
	if err := c.ShouldBindUri(&personRequestUri); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(response.NewServerError(404, err.Error(), err)))
		return
	}

	log.Printf("[request] get person info with %+v", personRequestUri)
	person, err := m.personRepo.FindPerson(personRequestUri.PersonId)
	if err != nil {
		log.Printf("get person list with path: %s", err.Error())
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(response.NewServerError(500, err.Error(), err)))
		return
	} else {
		httpStatus = http.StatusOK
		responseBody = response.NewPersonResponse(*model.NewPersonDto(person))
	}
	c.JSON(httpStatus, responseBody)
	log.Printf("[response] get person with (%d)%+v", httpStatus, responseBody)
}
