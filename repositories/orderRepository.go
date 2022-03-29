package repositories

import (
	"hacktiv8_assignment_3/models"

	"github.com/jinzhu/gorm"
)

type OrderRepository interface {
	IndexOrder() []models.Order
	StoreOrder(order models.Order) models.Order
}

type orderDB struct {
	connection *gorm.DB
}

func NewOrderDB(db *gorm.DB) OrderRepository {
	return &orderDB{
		connection: db,
	}
}

func (db orderDB) IndexOrder() []models.Order {

	return []models.Order{}
}

func (db orderDB) StoreOrder(order models.Order) models.Order {
	db.connection.Save(&order)
	db.connection.Preload("Items").Find(&order)

	return order
}
