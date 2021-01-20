package main;

import(
	"net/http"
	_ "docs"
    services "services"
    controllers "controllers"
)

type ProductDTO struct {
    Code string `json:"code"`
    Price uint  `json:"price"`
    }

// @title Go Rest API
// @version 1.0
// @description Go Rest API with SQL SERVER DB
// @contact.name Marc Kenneth Lomio
// @contact.email marckenneth.lomio@gmail.com
// @host localhost:9000
// @BasePath /
func main() { 
	controllers.HandleRequests()
}

// @Summary Migrate tables to the SQL Server
// @Description 
// @Tags migration
// @Produce  json
// @Success 200
// @Router /migration [post]
func CreateDatabaseSchema(w http.ResponseWriter, r *http.Request){
    services.CreateDatabaseSchema(w,r)
   }

// @Summary Get all products
// @Description 
// @Tags products
// @Produce  json
// @Success 200
// @Router /products [get]
func ReturnAllProducts(w http.ResponseWriter, r *http.Request){
    services.ReturnAllProducts(w,r)
   }

// @Summary Create Product
// @Description 
// @Tags products
// @Produce  json
// @Param product body ProductDTO true "Create product"
// @Success 200 {object} ProductDTO
// @Router /product [post]
func CreateNewProduct(w http.ResponseWriter, r *http.Request){
    services.CreateNewProduct(w,r)
   }


