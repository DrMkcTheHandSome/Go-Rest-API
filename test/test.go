package main

import (
	"fmt"
	services "services"
)

func main(){
	fmt.Println("TEST START")
	services.ReturnAllProducts()
	services.CreateDatabaseSchema()
}
