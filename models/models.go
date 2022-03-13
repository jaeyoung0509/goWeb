package models

import "gorm.io/gorm"

//Book Model
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}

type User struct {
	gorm.Model
	Id       uint   `json"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"password"`
}
