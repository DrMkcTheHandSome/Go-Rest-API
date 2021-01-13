package main;

import(
"encoding/json"
"fmt"
"log"
"net/http"
"github.com/gorilla/mux"
"gorm.io/gorm"
"gorm.io/driver/sqlite"
)
// Global Variables
type Product struct {
 gorm.Model
 Code string
 Price uint
}

var Products []Product

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
   myRouter.HandleFunc("/article", createNewProduct).Methods("POST")
   log.Fatal(http.ListenAndServe(":9000", myRouter))
}

// LOGIC
func createNewProduct(w http.ResponseWriter, r *http.Request) {
    // reqBody, _ := ioutil.ReadAll(r.Body)
    //var product Product 
    fmt.Println(r)
    // json.Unmarshal(reqBody, &product)
    // fmt.Println(&product)
    // fmt.Println(&product.Code)
    json.NewEncoder(w).Encode(r)
}

