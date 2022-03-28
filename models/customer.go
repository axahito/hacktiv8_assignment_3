package models

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	ID      int    `json:"customer_id"`
	Name    string `json:"customer_name"`
	Address string `json:"address"`
	Orders  []Order
}
