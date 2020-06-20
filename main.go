package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nuttaphon-rd/finalexam/customer"
	"log"
)

func main() {
	fmt.Println("Customer API starting")
	r := gin.Default()
	customer.SetupRoutes(r)
	log.Fatal(r.Run(":2019"))
}
