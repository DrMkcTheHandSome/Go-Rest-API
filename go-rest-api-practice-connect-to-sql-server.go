package main;

import(
"encoding/json"
"fmt"
"log"
"net/http"
"gorm.io/gorm"
"io/ioutil"
"github.com/gorilla/mux"
"gorm.io/driver/sqlserver"
"golang.org/x/crypto/bcrypt"
"time"
)

// TO DO: Refactor

// Global Variables
type Product struct {
 gorm.Model
 Code string `gorm:"column:code"`
 Price uint  `gorm:"column:price"`
}

type User struct{
 gorm.Model
 Email string    `json:"email" gorm:"unique"` 
 Password string `json:"password"`
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func main() { 
	handleRequests()
}

func handleRequests() {
   initializeRoutes()
   fmt.Println("Hello Go!") 
}

func initializeRoutes(){
   initRoutesByGorillaMux()
}

func initRoutesByGorillaMux(){
   myRouter := mux.NewRouter().StrictSlash(true)
   myRouter.HandleFunc("/", homePage)
   myRouter.HandleFunc("/migration", createDatabaseSchema).Methods("POST")
   myRouter.HandleFunc("/product", createNewProduct).Methods("POST")
   myRouter.HandleFunc("/product/{id}", updateProduct).Methods("PUT")
   myRouter.HandleFunc("/products", returnAllProducts).Methods("GET")
   myRouter.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")
   myRouter.HandleFunc("/product/{id}",returnSingleProduct).Methods("GET")
   myRouter.HandleFunc("/user", createNewUser).Methods("POST")
   myRouter.HandleFunc("/user/login", loginUser).Methods("POST")
   myRouter.HandleFunc("/users", returnAllUsers).Methods("GET")
   log.Fatal(http.ListenAndServe(":9000", myRouter))
}

 

func createDatabaseSchema(w http.ResponseWriter, r *http.Request){
 
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
 db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		fmt.Println("failed to connect database") 
        panic("failed to connect database")
    }
 
    // Migrate the schema
	db.Migrator().CreateTable(&Product{})
	db.Migrator().CreateTable(&User{})	
}

// LOGIC
func createNewProduct(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: createNewProduct")
	
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		fmt.Println("failed to connect database") 
        panic("failed to connect database")
    }
    reqBody, _ := ioutil.ReadAll(r.Body)
    var product Product 
    json.Unmarshal(reqBody, &product)
	db.Exec("INSERT INTO products (created_at,code,price) VALUES (?,?,?)",time.Now(), product.Code,product.Price)
    json.NewEncoder(w).Encode(product)	 
}

func updateProduct(w http.ResponseWriter, r *http.Request){
 fmt.Println("Endpoint Hit: updateProduct")
 
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		fmt.Println("failed to connect database") 
        panic("failed to connect database")
    }

    vars := mux.Vars(r)
    key := vars["id"]
    reqBody, _ := ioutil.ReadAll(r.Body)
    var product Product 
   //Update multiple columns
    json.Unmarshal(reqBody, &product)
	db.Exec("UPDATE products SET code=?, price = ? WHERE id = ?", product.Code, product.Price, key)
    json.NewEncoder(w).Encode(product)

}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
     fmt.Println("Endpoint Hit: returnAllProducts")
	
   connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
    db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		 fmt.Println("failed to connect database") 
        panic("failed to connect database")
     }
	
    // Get all records
	var products []Product
    db.Exec("select * from products").Scan(&products)
	
    json.NewEncoder(w).Encode(products)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Endpoint Hit: deleteProduct")
  
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		fmt.Println("failed to connect database") 
        panic("failed to connect database")
    }

   vars := mux.Vars(r)
    key := vars["id"]
    
   db.Exec("DELETE FROM products WHERE id = ?", key)
   returnAllProducts(w,r)
} 

func returnSingleProduct(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnSingleProduct")
	
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		fmt.Println("failed to connect database") 
        panic("failed to connect database")
    }
    
	
    vars := mux.Vars(r)
    key := vars["id"]

	var product Product
    db.Exec("select * from products where id = ?",key).Scan(&product)
		
    //Multiple Query Example
    //db.Raw("select code,price from products; drop table product;").First(&product)

    json.NewEncoder(w).Encode(product)  
}

func loginUser(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: loginUser")
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		fmt.Println("failed to connect database") 
        panic("failed to connect database")
    }
	
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User 
	var userPayload User
	json.Unmarshal(reqBody, &user)
	userPayload = user
	db.Exec("select * from users where email = ?",user.Email).Scan(&user)
	 err = checkPassword(user.Password,userPayload.Password)
	 if err != nil {
	    fmt.Println("Login Failed")
	 } else {
	    fmt.Println("Login Success")
	 }
}

func createNewUser(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: createNewUser")
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		fmt.Println("failed to connect database") 
        panic("failed to connect database")
    }
    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var user User 
	var hash_password string = ""
    json.Unmarshal(reqBody, &user)
    hash_password = hashPassword(user.Password)
    db.Exec("INSERT INTO users (created_at,email,password) VALUES (?,?,?)",time.Now(), user.Email,hash_password)
    db.Create(&User{Email: user.Email, Password: hash_password})
	user.Password = hash_password
    json.NewEncoder(w).Encode(user)
}

func returnAllUsers(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllUsers")
	
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		fmt.Println("failed to connect database") 
        panic("failed to connect database")
    }
	  var users []User
    db.Exec("select * from users").Scan(&users)
      json.NewEncoder(w).Encode(users)
}

func hashPassword(password string) string {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "failed generate bcrypt password"
	}
    
	var hash_password string = ""
	hash_password = string(bytes)

	return hash_password
}


  func checkPassword(userPasswordfromDB,providedPassword string) error {
	  err := bcrypt.CompareHashAndPassword([]byte(userPasswordfromDB), []byte(providedPassword))
	  if err != nil {
		  return err
	  }

	  return nil
 }





