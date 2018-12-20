package structs

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID int
	First_Name string
	Last_Name string
	Email string
	Password string
}
