package requests

type TodoGet struct{
	Id int `param:"id"`
}

type TodoPost struct {
	UserId int `json:"userId"`
	Title string `json:"title"`
	Status int `json:"status"`
}

type TodoPut struct {
	Id int `param:"id"`
	UserId int `json:"userId"`
	Title string `json:"title"`
	Status int `json:"status"`
}

type TodoDelete struct {
	Id int `param:"id"`
}