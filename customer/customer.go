package customer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nuttaphon-rd/finalexam/errors"
	"net/http"
)

type CustomerHandler interface {
	CreateCustomer(*gin.Context)
	GetCustomerById(*gin.Context)
	GetAllCustomer(*gin.Context)
	UpdateCustomer(*gin.Context)
	DeleteCustomer(*gin.Context)
}

type CustomerHandle struct {
	Service CustomerServicer
}

func NewHandle(s CustomerServicer) CustomerHandle {
	return CustomerHandle{s}
}

func (ch *CustomerHandle) CreateCustomer(c *gin.Context) {
	cus := Customer{}
	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error bind": err.Error()})
		return
	}

	if err := ch.Service.CreateCustomerDB(&cus); err != nil {
		c.JSON(err.Code, err.Message)
		return
	}

	c.JSON(http.StatusCreated, cus)
}

func (ch *CustomerHandle) GetCustomerById(c *gin.Context) {
	id := c.Param("id")
	var customer *Customer
	var err *errors.Error
	if customer, err = ch.Service.GetCustomerByIdDB(id); err != nil {
		c.JSON(err.Code, err.Message)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (ch *CustomerHandle) GetAllCustomer(c *gin.Context) {
	var customers []*Customer
	var err *errors.Error
	if customers, err = ch.Service.GetAllCustomerDB(); err != nil {
		c.JSON(err.Code, err.Message)
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (ch *CustomerHandle) UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	cus := Customer{}
	if err := c.ShouldBindJSON(&cus); err != nil {
		fmt.Println("error binding," + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error bind": err.Error()})
		return
	}

	var err *errors.Error
	var customer *Customer
	if customer, err = ch.Service.UpdateCustomerDB(cus, id); err != nil {
		c.JSON(err.Code, err.Message)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (ch *CustomerHandle) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	if err := ch.Service.DeleteCustomerDB(id); err != nil {
		c.JSON(err.Code, err.Message)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}
