package repository

import (
	"okakyo/practice-go/domain/entity"
)

type Todo interface {
	FindAll()(*entity.Todos,error)
	FindOne(id int)(*entity.Todo,error)
	Store(todo *entity.Todo)(*entity.Todo,error)
	Update(cultivation *entity.Todo) (*entity.Todo, error)
	Delete(id int)error
}