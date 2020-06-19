package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
)

func SetupModels() *gorm.DB {
	postgresConn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		"localhost", "5432", "", "", "golang-demo")

	db, err := gorm.Open("postgres", postgresConn)
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&User{})
	//m := User{Name: "Nathan", Email: "nathan@example.com"}
	//db.Create(&m)
	return db
}
