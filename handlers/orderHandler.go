package handlers

import (
	"fmt"
	"hacktiv8_assignment_3/config"
	"hacktiv8_assignment_3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexOrder(ctx *gin.Context) {
	var orders []models.Order
	var result gin.H
	db := config.GetDB()

	db.Preload("Orders").Find(&orders)
	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func ShowOrder(ctx *gin.Context) {
	var order models.Order
	var result gin.H
	db := config.GetDB()
	id := ctx.Param("order")

	err := db.Where("id = ?", id).Preload("Customer").First(&order).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	} else {
		result = gin.H{
			"result": order,
		}
	}

	ctx.JSON(http.StatusOK, result)

}

func CreateOrder(ctx *gin.Context) {
	var order models.Order
	var customer models.Customer
	var result gin.H
	db := config.GetDB()
	ctx.Bind(&order)
	ctx.Bind(&customer)

	err := db.Model(&order).Association("Customer").Append(&customer)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	result = gin.H{
		"new order": order,
	}

	ctx.JSON(http.StatusOK, result)
}

func UpdateOrder(ctx *gin.Context) {
	var order models.Order
	var newOrder models.Order
	var result gin.H
	db := config.GetDB()
	id := ctx.Param("order")
	fmt.Println("order id : ", id)
	ctx.Bind(&newOrder)

	err := db.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
	}

	err = db.Model(&order).Updates(newOrder).Error
	if err != nil {
		result = gin.H{
			"result": "Update Failed",
		}
	} else {
		result = gin.H{
			"result": "Update Success",
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func DeleteOrder(ctx *gin.Context) {
	var order models.Order
	var result gin.H
	db := config.GetDB()
	id := ctx.Param("order")

	err := db.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	err = db.Delete(&order).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "delete success",
		}
	}

	ctx.JSON(http.StatusOK, result)
}
