package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	ID         int       `json:"order_id"`
	OrderedAt  time.Time `json:"ordered_at"`
	CustomerID int       `json:"customer_id"`
	Customer   Customer
	Items      []Item
}
