package database

import (
	"okakyo/practice-go/domain/entity"
	"okakyo/practice-go/domain/repository"

	"gorm.io/gorm"
)


var _ repository.Todo = &TodoRepository{}

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func(r *TodoRepository) FindAll()(*entity.Todos,error){
	var todos entity.Todos
	if err:= r.db.Find(&todos).Error;err!= nil {
		return nil, err
	}
	return &todos,nil
}
func(r *TodoRepository) FindOne(id int)(*entity.Todo, error){
	var todo entity.Todo
	if err:= r.db.First(&todo).Error;err!= nil {
		return nil, err
	}
	return &todo,nil
}
func(r *TodoRepository) Store(todo *entity.Todo)(*entity.Todo ,error){
	if err:= r.db.Preload("todos").Create(&todo).First(&todo).Error;err!= nil {
		return nil, err
	}
	return todo,nil
}

func(r *TodoRepository) Update(todo *entity.Todo)(*entity.Todo,error) {
	if err := r.db.Model(&entity.Todo{}).Updates(&todo).First(&todo).Error; err != nil {
		return nil, err
	}
	return todo,nil
}

func(r *TodoRepository) Delete(id int) error{
	var todo entity.Todo
	if err:= r.db.Where("id =?",id).Delete(&todo).Error;err!= nil {
		return  err
	}
	return nil
}