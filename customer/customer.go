package customer

import (
	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {

}

func GetCustomerById(c *gin.Context) {

}

func GetAllCustomer(c *gin.Context) {

}

func UpdateCustomer(c *gin.Context) {

}

func DeleteCustomer(c *gin.Context) {

}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	customer := r.Group("/customers")

	customer.POST("/", CreateCustomer)
	customer.GET("/:id", GetCustomerById)
	customer.GET("/", GetAllCustomer)
	customer.PUT("/:id", UpdateCustomer)
	customer.DELETE("/:id", DeleteCustomer)

	return r
}
