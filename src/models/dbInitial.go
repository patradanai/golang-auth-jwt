package models

import (
	"gorm.io/gorm"
)

type Handler struct{ Db *gorm.DB }

// Initial Role
var roles = []Role{{Role: "Admin"}, {Role: "Customer"}, {Role: "Support"}}

func (pointerDb *Handler) OpenDb() {
	(*pointerDb).Db.AutoMigrate(&Role{}, &User{})
	// DB.Create(&roles)

}
