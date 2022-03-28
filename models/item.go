package models

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	ID          int    `json:"item_id"`
	Item_Code   int    `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"order_id"`
}
