package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	DB, err = gorm.Open(sqlite.Open("storage-service.db"), &gorm.Config{})
	if err != nil {
		panic("cannot conncet to database")
	}
	fmt.Println("connected to database")
}
