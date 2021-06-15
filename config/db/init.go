package db

import (
	"fmt"
	"os"
)

func DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		"db","3306",
		os.Getenv("MYSQL_DATABASE"),
	)+"?charset=utf8mb4&parseTime=True&loc=Local"
}

