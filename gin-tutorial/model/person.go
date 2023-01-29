package model

// Person gorm model
type Person struct {
	Id   int32
	Name string
}

func (Person) TableName() string {
	return "person"
}

// Person dto
type PersonDto struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

func NewPersonDto(person *Person) *PersonDto {
	return &PersonDto{
		Id:   person.Id,
		Name: person.Name,
	}
}
