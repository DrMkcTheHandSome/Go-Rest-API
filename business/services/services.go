package services

import(
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
	repositories "repositories"
	entities "entities"
	"github.com/gorilla/mux"
	helpers "helpers"
	)

	func HomePage(w http.ResponseWriter, r *http.Request){
		fmt.Println("services homePage")
		var htmlIndex = `<html>
	<body>
	   <h1>Welcome to the homepage!</h1>
		<a href="/user/loginViaGoogle">Google Log In</a>
	</body>
	</html>`
		fmt.Fprintf(w, htmlIndex)
	}

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

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	fmt.Println("services updateProduct")
	   
	   vars := mux.Vars(r)
	   key := vars["id"]
	   reqBody, _ := ioutil.ReadAll(r.Body)
	   var product entities.Product 

	   json.Unmarshal(reqBody, &product)
	   repositories.UpdateProduct(key,product)
	   product = repositories.GetSingleProduct(key)
	   json.NewEncoder(w).Encode(product)
	   w.WriteHeader(http.StatusCreated)
   }

   func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("services DeleteProduct")

	 vars := mux.Vars(r)
	 key := vars["id"]
  
	 repositories.DeleteProduct(key)	
	 ReturnAllProducts(w,r)
  } 


  func CreateNewUser(w http.ResponseWriter, r *http.Request){
    fmt.Println("services createNewUser")

    reqBody, _ := ioutil.ReadAll(r.Body)
    var user entities.User 
	var hash_password string = ""
    json.Unmarshal(reqBody, &user)
	hash_password = helpers.HashPassword(user.Password)
	user = repositories.CreateNewUser(user,hash_password)
	user.Password = hash_password
    json.NewEncoder(w).Encode(user)
}

func ReturnAllUsers(w http.ResponseWriter, r *http.Request){
    fmt.Println("services returnAllUsers")
	
	  var users []entities.User
	
	  users = repositories.GetAllUsers() 
	
	  json.NewEncoder(w).Encode(users)
}