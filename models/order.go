package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	OrderedAt    time.Time `json:"ordered_at"`
	CustomerName string    `json:"customer_name"`
	Items        []Item    `json:"items" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
