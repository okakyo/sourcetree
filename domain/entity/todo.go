package entity

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	UserId int
	Title  string
	Status int 
}

type Todos []Todo