package database

import (
	"fmt"
	"log"
	"net/http"

	"okakyo/practice-go/config/db"
	"okakyo/practice-go/domain/entity"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func NewSqlHandler() (*gorm.DB,error){
	
	db, err := gorm.Open(mysql.Open(db.DSN()),&gorm.Config{})
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

func ResetDB(c echo.Context) error {
	db, err := gorm.Open(mysql.Open(db.DSN()),&gorm.Config{})
	if err != nil {
		return c.String(http.StatusOK, "Reset Completed!")
	}
	dropTables(db)
	migrate(db)

	sqlDB,err:= db.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)

	if err := sqlDB.Ping(); err != nil {
		_ = fmt.Errorf("failed to ping: %w", err)
		return c.String(http.StatusInternalServerError, "Reset Failed")
	}
	return c.String(http.StatusOK, "Reset Completed!")
}

func migrate(db *gorm.DB) {
	log.Println("Start auto migration")
	db.AutoMigrate(
		&entity.Todo{},
	)
	log.Println("Finish auto migration.")
}

func dropTables(db *gorm.DB) {
	log.Println("Start drop tables")
	db.Migrator().DropTable(&entity.Todo{})
	log.Println("Finish drop tables")
}