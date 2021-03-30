package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Initial Role
var roles = []Role{{Role: "Admin"}, {Role: "Customer"}, {Role: "Support"}}

func OpenDb() {
	var err error
	// Please define your username and password for MySQL.
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Connection Failed to Open")
	} else {
		fmt.Println("Connection Established")
	}

	DB.AutoMigrate(&Role{}, &User{})
	// db.Create(&roles)

}
