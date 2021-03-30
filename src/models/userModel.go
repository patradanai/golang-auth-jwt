package models

import "gorm.io/gorm"

// Relationship 1 to mamy

type Role struct {
	gorm.Model
	Role  string
	Users []User
}

type User struct {
	gorm.Model
	Email    string
	Password string
	Fname    string
	Lname    string
	RoleID   uint
}
