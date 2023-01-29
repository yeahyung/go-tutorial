package response

import "go-tutorial/gin-tutorial/model"

type PersonResponse struct {
	Person model.PersonDto `json:"person"`
}

func NewPersonResponse(person model.PersonDto) PersonResponse {
	var personResponse PersonResponse
	personResponse.Person = person
	return personResponse
}
