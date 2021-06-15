package handler

import (
	"net/http"
	"okakyo/practice-go/domain/entity"
	"okakyo/practice-go/usecase"
	"okakyo/practice-go/web/dto/requests"

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
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK,todos)
}

func (h *TodoHandlers) GetTodo(c echo.Context) error {
	var args requests.TodoGet
	if err:= c.Bind(&args);err!=nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	todo, err := h.todoUC.ReadTodo(args.Id)
	if err!= nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK,todo)

}

func (h *TodoHandlers) PostTodo(c echo.Context) error {
	var req requests.TodoPost
	if err:= c.Bind(&req);err!= nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	createTodo := &entity.Todo{
		UserId: req.UserId ,
		Title: req.Title,
		Status: req.Status,
	}
	todo,err:= h.todoUC.CreateTodo(createTodo)
	if err!= nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK,todo)

}

func (h *TodoHandlers) PutTodo(c echo.Context) error {
	var req requests.TodoPut
	if err:= c.Bind(&req);err!= nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	updateTodo := &entity.Todo{
		UserId: req.UserId,
		Title: req.Title,
		Status: req.Status,
	}
	todo,err:= h.todoUC.UpdateTodo(updateTodo)
	if err!= nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK,todo)
}

func (h *TodoHandlers) DeleteTodo(c echo.Context) error {
	var args requests.TodoDelete
	if err:=c.Bind(&args);err!=nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err:= h.todoUC.DeleteTodo(args.Id)
	if err!= nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK,nil)
}
