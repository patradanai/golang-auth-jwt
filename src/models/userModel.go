package models

import "gorm.io/gorm"

// Relationship many to mamy

type Role struct {
	gorm.Model
	ID   uint `gorm:"primaryKey"`
	Role string
}

type User struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Email    string
	Username string
	Password string
	Fname    string
	Lname    string
	Roles    []Role `gorm:"many2many:user_roles;"`
}
