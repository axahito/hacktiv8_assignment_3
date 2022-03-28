package models

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	Name    string  `json:"customer_name"`
	Address string  `json:"address"`
	Orders  []Order `json:"orders"`
}
