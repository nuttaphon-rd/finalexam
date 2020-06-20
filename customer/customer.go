package customer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nuttaphon-rd/finalexam/errors"
	"github.com/nuttaphon-rd/finalexam/middleware"
	"net/http"
)

func CreateCustomer(c *gin.Context) {
	cus := Customer{}
	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error bind": err.Error()})
		return
	}

	if err := CreateCustomerDB(&cus); err != nil {
		c.JSON(err.Code, err.Message)
		return
	}

	c.JSON(http.StatusCreated, cus)
}

func GetCustomerById(c *gin.Context) {
	id := c.Param("id")
	var customer *Customer
	var err *errors.Error
	if customer, err = GetCustomerByIdDB(id); err != nil {
		c.JSON(err.Code, err.Message)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func GetAllCustomer(c *gin.Context) {
	var customers []*Customer
	var err *errors.Error
	if customers, err = GetAllCustomerDB(); err != nil {
		c.JSON(err.Code, err.Message)
		return
	}
	c.JSON(http.StatusOK, customers)
}

func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	cus := Customer{}
	if err := c.ShouldBindJSON(&cus); err != nil {
		fmt.Println("error binding," + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error bind": err.Error()})
		return
	}

	var err *errors.Error
	var customer *Customer
	if customer, err = UpdateCustomerDB(cus, id); err != nil {
		c.JSON(err.Code, err.Message)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	if err := DeleteCustomerDB(id); err != nil {
		c.JSON(err.Code, err.Message)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.AuthMiddleware)

	customer := r.Group("/customers")

	customer.POST("/", CreateCustomer)
	customer.GET("/:id", GetCustomerById)
	customer.GET("/", GetAllCustomer)
	customer.PUT("/:id", UpdateCustomer)
	customer.DELETE("/:id", DeleteCustomer)

	return r
}
