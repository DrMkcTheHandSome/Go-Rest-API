package main;

import(
	"net/http"
	_ "docs"
	httpSwagger "github.com/swaggo/http-swagger"
	services "services"
)

// @title Users Product Go Rest API
// @version 1.0
// @description Go Rest API with SQL SERVER DB
// @contact.name Marc Kenneth Lomio
// @contact.email marckenneth.lomio@gmail.com
// @host localhost:9000
// @BasePath /
func main() { 
	handleRequests()
}

// homePage godoc
// @Summary show html that navigates to google auth login
// @Description 
// @Produce  json
// @Success 200
// @Router / [get]
func createDatabaseSchema(w http.ResponseWriter, r *http.Request){
    services.createDatabaseSchema()
   }

// returnAllProducts godoc
// @Summary show html that navigates to google auth login
// @Description 
// @Produce  json
// @Success 200
// @Router / [get]
func returnAllProducts(w http.ResponseWriter, r *http.Request){
    services.returnAllProducts()
   }