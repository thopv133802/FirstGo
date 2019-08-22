package services

import (
	"GoFirst/src/domain/models"
	"GoFirst/src/domain/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommonService interface {
	GetProducts(context *gin.Context)
	GetPromos(context *gin.Context)
	AddUser(context *gin.Context)
	SignIn(context *gin.Context)
	SignOut(context *gin.Context)
	GetOrders(context *gin.Context)
	Charge(context *gin.Context)
}

type CommonServiceImpl struct {
	Repository repositories.CommonRepository
}

func (service CommonServiceImpl) GetProducts(context *gin.Context) {
	products, err := service.Repository.GetAllProducts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, products)
}


func (service CommonServiceImpl) GetPromos(context *gin.Context) {
	promos, err := service.Repository.GetPromos()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, promos)
}

func (service CommonServiceImpl) AddUser(context *gin.Context) {
	var customer models.Customer
	err := context.ShouldBindJSON(&customer)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer,err = service.Repository.AddUser(customer)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, customer)
}

func (service CommonServiceImpl) SignIn(context *gin.Context) {
	var customer models.Customer
	err := context.ShouldBindJSON(&customer)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err = service.Repository.SignInUser(customer.Email, customer.Password)
	if err != nil {
		if err == repositories.ERROR_INVALID_PASSWORD {
			context.JSON(http.StatusForbidden, gin.H {"error": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, customer)
}

func (service CommonServiceImpl) SignOut(context *gin.Context) {
	userID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = service.Repository.SignOutUserById(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (service CommonServiceImpl) GetOrders(context *gin.Context) {
	userID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orders, err := service.Repository.GetCustomerOrdersByID(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, orders)
}

func (service CommonServiceImpl) Charge(context *gin.Context) {
	if service.Repository == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
}






