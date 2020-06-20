package customer

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/nuttaphon-rd/finalexam/middleware"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func SetupRoutes(r *gin.Engine)  {
	r.Use(middleware.AuthMiddleware)

	var err error
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	s := NewService(db)
	h := NewHandle(s)

	customer := r.Group("/customers")

	customer.POST("/", h.CreateCustomer)
	customer.GET("/:id", h.GetCustomerById)
	customer.GET("/", h.GetAllCustomer)
	customer.PUT("/:id", h.UpdateCustomer)
	customer.DELETE("/:id", h.DeleteCustomer)

}
