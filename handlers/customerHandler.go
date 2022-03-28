package handlers

import (
	"fmt"
	"hacktiv8_assignment_3/config"
	"hacktiv8_assignment_3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexCustomer(ctx *gin.Context) {
	var customers []models.Customer
	var result gin.H
	db := config.GetDB()

	db.Preload("Orders").Find(&customers)
	if len(customers) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": customers,
			"count":  len(customers),
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func ShowCustomer(ctx *gin.Context) {
	var customer models.Customer
	var result gin.H
	db := config.GetDB()
	id := ctx.Param("customer")

	err := db.Where("id = ?", id).Preload("Orders").First(&customer).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	} else {
		result = gin.H{
			"result": customer,
		}
	}

	ctx.JSON(http.StatusOK, result)

}

func CreateCustomer(ctx *gin.Context) {
	var customer models.Customer
	var result gin.H
	db := config.GetDB()
	ctx.Bind(&customer)

	err := db.Create(&customer).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	result = gin.H{
		"new Customer": customer,
	}

	ctx.JSON(http.StatusOK, result)
}

func UpdateCustomer(ctx *gin.Context) {
	var customer models.Customer
	var newCustomer models.Customer
	var result gin.H
	db := config.GetDB()
	id := ctx.Param("customer")
	fmt.Println("customer id : ", id)
	ctx.Bind(&newCustomer)

	err := db.First(&customer, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
	}

	err = db.Model(&customer).Updates(newCustomer).Error
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

func DeleteCustomer(ctx *gin.Context) {
	var customer models.Customer
	var result gin.H
	db := config.GetDB()
	id := ctx.Param("customer")

	err := db.First(&customer, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	err = db.Delete(&customer).Error
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
