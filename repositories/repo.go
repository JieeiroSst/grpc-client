package repositories

import (
	"github.com/JIeeiroSst/go-app/entities"
)

type ItemRepository interface {
	FindAll() ([]entities.Item,error)
	FindByID(id string) (entities.Item,error)
	UpdateItem(item entities.Item,id string) error 
	DeleteItem(id string) error
	CreateItem(item entities.Item) error
}