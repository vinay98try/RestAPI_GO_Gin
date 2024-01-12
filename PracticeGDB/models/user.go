package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// ID       int    `json:"id" gorm: "primary_key"` 	
	// No need to define Id here because it already present in the GORM Model...!
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
