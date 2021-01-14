package main;

import(
"encoding/json"
"fmt"
"log"
"net/http"
"github.com/gorilla/mux"
 "io/ioutil"
"gorm.io/gorm"
"gorm.io/driver/sqlite"
)
// Global Variables
type Product struct {
 gorm.Model
 Code string
 Price uint
}

//TO DO: Refactor follow DRY principle

// Trigger Functions at start  
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}
 
func handleRequests() {
   initializeRoutes()
   fmt.Println("Hello Go!") 
}
 
func main() { // like ngOnInit in Angular
    createDatabaseSchema()
    handleRequests()
}

func createDatabaseSchema(){
 db, err := gorm.Open(sqlite.Open("practice.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
 
    // Migrate the schema
    db.AutoMigrate(&Product{})
 
    // Create
    db.Create(&Product{Code: "P1", Price: 100})
}

func initializeRoutes(){
   initRoutesByGorillaMux()
}

func initRoutesByGorillaMux(){
   myRouter := mux.NewRouter().StrictSlash(true)
   myRouter.HandleFunc("/", homePage)
   myRouter.HandleFunc("/product", createNewProduct).Methods("POST")
   myRouter.HandleFunc("/product/{id}", updateProduct).Methods("PUT")
   myRouter.HandleFunc("/products", returnAllProducts).Methods("GET")
   myRouter.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")
   myRouter.HandleFunc("/product/{code}",returnSingleProduct).Methods("GET")
   log.Fatal(http.ListenAndServe(":9000", myRouter))
}

// LOGIC
func createNewProduct(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: createNewProduct")
  db, err := gorm.Open(sqlite.Open("practice.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    reqBody, _ := ioutil.ReadAll(r.Body)
    var product Product 
    json.Unmarshal(reqBody, &product)
    db.Create(&Product{Code: product.Code, Price: product.Price})
    json.NewEncoder(w).Encode(product)
}

func returnSingleProduct(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnSingleProduct")
db, err := gorm.Open(sqlite.Open("practice.db"), &gorm.Config{})
    
    if err != nil {
        panic("failed to connect database")
    }
    
    vars := mux.Vars(r)
    key := vars["code"]
    var product Product
    
    db.First(&product,"code = ?",key)
 
    json.NewEncoder(w).Encode(product)    
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllProducts")
db, err := gorm.Open(sqlite.Open("practice.db"), &gorm.Config{})
    
    if err != nil {
        panic("failed to connect database")
    }
    // Find all of our products.
    var products []Product
    db.Find(&products)
    json.NewEncoder(w).Encode(products)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Endpoint Hit: deleteProduct")
  db, err := gorm.Open(sqlite.Open("practice.db"), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

   vars := mux.Vars(r)
    key := vars["id"]
    
   var product Product
   db.Delete(&product,key)
   returnAllProducts(w,r)
} 

func updateProduct(w http.ResponseWriter, r *http.Request){
 fmt.Println("Endpoint Hit: updateProduct")
  db, err := gorm.Open(sqlite.Open("practice.db"), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

    vars := mux.Vars(r)
    key := vars["id"]
    reqBody, _ := ioutil.ReadAll(r.Body)
    var product Product 
   //Update multiple columns
    json.Unmarshal(reqBody, &product)
    db.Model(&product).Where("id = ?", key).Updates(Product{Code: product.Code, Price: product.Price})
    json.NewEncoder(w).Encode(product)
}

