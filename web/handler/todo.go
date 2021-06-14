package handler

import (
	"log"
	"net/http"
	"okakyo/practice-go/usecase"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)


type TodoHandlers struct {
	todoUC *usecase.TodoUseCase
}

func NewTodoHandler(db *gorm.DB) *TodoHandlers {
	return &TodoHandlers{todoUC: usecase.NewUserCase(db)}
}

func (h *TodoHandlers) GetTodos(c echo.Context) error{
	todos, err := h.todoUC.ReadAllTodos()
	if err!= nil {
		
	}
	return c.JSON(http.StatusOK,todos)
}

func (h *TodoHandlers) GetTodo(c echo.Context) error {
	id := c.Param("id")
	todo, err := h.todoUC.ReadTodo(id)

}