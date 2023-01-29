package repository

import (
	"go-tutorial/gin-tutorial/model"
	"gorm.io/gorm"
)

type PersonRepo interface {
	FindPerson(id int32) (*model.Person, error)
}

type GormPersonRepo struct {
	db *gorm.DB
}

func NewGormPersonRepo(db *gorm.DB) *GormPersonRepo {
	return &GormPersonRepo{db: db}
}

func (g *GormPersonRepo) FindPerson(id int32) (*model.Person, error) {
	var person model.Person
	if err := g.db.Where("id = ?", id).First(&person).Error; err != nil {
		return nil, err
	}
	return &person, nil
}
