package services

import(
	//"encoding/json"
	"net/http"
	"fmt"
	 repositories "repositories"
	//entities "entities"
	)



func CreateDatabaseSchema(w http.ResponseWriter, r *http.Request){
	fmt.Println("services CreateDatabaseSchema")
	repositories.SchemaMigration()
	}
	
func ReturnAllProducts(w http.ResponseWriter, r *http.Request) {
     fmt.Println("services ReturnAllProducts")
	
   // Get all records
	//var products []entities.Product 
	repositories.GetAllProducts()
	// fmt.Println(products)
    // json.NewEncoder(w).Encode(products)
}







