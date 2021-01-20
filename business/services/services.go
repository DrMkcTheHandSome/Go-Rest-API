package services

import(
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
	repositories "repositories"
	entities "entities"
	"github.com/gorilla/mux"
	)



func CreateDatabaseSchema(w http.ResponseWriter, r *http.Request){
	fmt.Println("services CreateDatabaseSchema")
	repositories.SchemaMigration()
	w.WriteHeader(http.StatusCreated)
	}
	
func ReturnAllProducts(w http.ResponseWriter, r *http.Request) {
     fmt.Println("services ReturnAllProducts")
	
	var products []entities.Product 
	products = repositories.GetAllProducts()
	 json.NewEncoder(w).Encode(products)
	 w.WriteHeader(http.StatusOK)
}

func CreateNewProduct(w http.ResponseWriter, r *http.Request){
    fmt.Println("services CreateNewProduct")

    reqBody, _ := ioutil.ReadAll(r.Body)
	var product entities.Product 
	json.Unmarshal(reqBody, &product)
	product = repositories.CreateNewProduct(product)
	json.NewEncoder(w).Encode(product)
	w.WriteHeader(http.StatusCreated)
}

func ReturnSingleProduct(w http.ResponseWriter, r *http.Request) {
    fmt.Println("services returnSingleProduct")
	
    vars := mux.Vars(r)
    key := vars["id"]

	var product entities.Product
		
	product = repositories.GetSingleProduct(key)

	json.NewEncoder(w).Encode(product)  
	w.WriteHeader(http.StatusOK)
}





/* HELPERS */

// func HashPassword(password string) string {
//     bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	if err != nil {
// 		return "failed generate bcrypt password"
// 	}
    
// 	var hash_password string = ""
// 	hash_password = string(bytes)

// 	return hash_password
// }


