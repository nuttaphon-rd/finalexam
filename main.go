package main

import (
	"fmt"
	"log"
	"github.com/nuttaphon-rd/finalexam/customer"
)

func main() {
	fmt.Println("Customer API starting")
	r := customer.SetupRouter()
	log.Fatal(r.Run(":2019"))
}
