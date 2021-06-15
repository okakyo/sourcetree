package db

import (
	"fmt"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



func NewSqlHandler() (*gorm.DB,error){
	user:= os.Getenv("MYSQL_USERNAME")
	dbname:= os.Getenv("MYSQL_DATABASE")
	password:= os.Getenv("MYSQL_PASSWORD")
	dsn := user+":"+password+"@tcp(db:3306)/"+dbname+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to open MySQL: %w",err)
	}
	sqlDB,err := db.DB()
	if err!= nil {
		return nil, fmt.Errorf("Failed to Ping: %W",err)
	}
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)

	if err := sqlDB.Ping();err!=nil {
		return nil, fmt.Errorf("Failed to Ping: %W",err)
	}
	return db,nil
}