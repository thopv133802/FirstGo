package models

import "github.com/jinzhu/gorm"

type Product struct{
	gorm.Model
	Image string `json:"img"`
	ImagAlt string `json:"imgalt" gorm:"column:imgalt"`
	Price float64 `json:"price"`
	Promotion float64 `json:"promotion"`
	ProductName string `gorm:"column:productname" json:"productname"`
	Description string `json:"desc"`
}
func (Product) TableName() string {
	return "products"
}