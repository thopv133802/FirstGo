package models

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	Name string `json:"name"`
	FirstName string `gorm:"column:firstname" json:"firstname"`
	LastName string `gorm:"column:lastname" json:"lastname"`
	Email string `gorm:"column:email" json:"email"`
	LoggedIn bool `gorm:"column:loggedin" json:"loggedin"`
	Password string `json:"password"`
	Orders []Order `json:"orders"`
}
func (Customer) TableName() string {
	return "customers"
}