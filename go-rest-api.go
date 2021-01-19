package main;

import(
	"net/http"
	_ "docs"
    services "services"
    controllers "controllers"
)

// @title Users Product Go Rest API
// @version 1.0
// @description Go Rest API with SQL SERVER DB
// @contact.name Marc Kenneth Lomio
// @contact.email marckenneth.lomio@gmail.com
// @host localhost:9000
// @BasePath /
func main() { 
	controllers.HandleRequests()
}

// homePage godoc
// @Summary show html that navigates to google auth login
// @Description 
// @Produce  json
// @Success 200
// @Router / [get]
func CreateDatabaseSchema(w http.ResponseWriter, r *http.Request){
    services.CreateDatabaseSchema(w,r)
   }


func ReturnAllProducts(w http.ResponseWriter, r *http.Request){
    services.ReturnAllProducts(w,r)
   }


