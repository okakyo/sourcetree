package usecase

import (
	"fmt"
	"okakyo/practice-go/database"
	"okakyo/practice-go/domain/entity"
	"okakyo/practice-go/domain/repository"
	"gorm.io/gorm"
)

type TodoUseCase struct {
    todoRepository repository.Todo
}

func NewUserCase(db *gorm.DB) *TodoUseCase {
    return &TodoUseCase{todoRepository: database.NewTodoRepository(db)}
}

func (u *TodoUseCase) ReadAllTodos() ( *entity.Todos ,error) { 
    todos, err := u.todoRepository.FindAll()
    if err!=nil {
	    return nil, fmt.Errorf("Find Todos from Repo: %w",err)
    }
    return todos, nil
}

func (u *TodoUseCase) ReadTodo(id int )( *entity.Todo ,error){
    todo, err := u.todoRepository.FindOne(id)
    if err!= nil {
	    return nil, fmt.Errorf("Find todo from Repo: %w",err)
    }
    return todo, nil
}

func (u *TodoUseCase) CreateTodo(todo *entity.Todo)( *entity.Todo ,error) {
    todo, err := u.todoRepository.Store(todo)
    if err!= nil {
	    return nil, fmt.Errorf("Create todo from Repo: %w",err)
    }
    return todo, nil 
}

func (u *TodoUseCase) UpdateTodo(todo *entity.Todo) ( *entity.Todo ,error){
	todo, err:= u.todoRepository.Update(todo)
	if err!=nil {
		return nil, fmt.Errorf("Update Todo from Repo: %w",err)
	}
	return todo, nil
}
func (u *TodoUseCase) DeleteTodo( id int )error{
	err:= u.todoRepository.Delete(id)
	if err!= nil {
		return fmt.Errorf("Delete Todo from Repo: %w", err)
	}
	return nil
}


