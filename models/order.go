package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	OrderedAt  time.Time `json:"ordered_at"`
	CustomerID int       `json:"customer_id"`
	Items      []Item    `json:"items"`
}
