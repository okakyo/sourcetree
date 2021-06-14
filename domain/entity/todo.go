package entity

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	userId int
	title  string
	status int 
}

type Todos []Todo