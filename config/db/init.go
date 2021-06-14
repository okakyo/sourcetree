package db

import (
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() *SqlHandler{
	user:= os.Getenv("MYSQL_USERNAME")
	dbname:= os.Getenv("MYSQL_DATABASE")
	password:= os.Getenv("MYSQL_PASSWORD")
	dsn := user+":"+password+"@tcp(db:3306)/"+dbname+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		panic("DB への接続に失敗しました。")
	}
	sqlHander:= new(SqlHandler)
	sqlHander.db = db 
	return sqlHander
	

}