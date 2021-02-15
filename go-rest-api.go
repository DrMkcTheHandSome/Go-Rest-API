package main;

import(
	"net/http"
	_ "docs"
    services "services"
    controllers "controllers"
    helpers "helpers"
)

type ProductDTO struct {
    Code string `json:"code"`
    Price uint  `json:"price"`
    }

type UserDTO struct{
        Email string    `json:"email"` 
        Password string `json:"password"`
        IsEmailVerified bool `json:"verified_email"`
        }



// @title Go Rest API
// @version 1.0
// @description Go Rest API with SQL SERVER DB
// @contact.name Marc Kenneth Lomio
// @contact.email marckenneth.lomio@gmail.com
// @host localhost:9000
// @BasePath /
func main() { 
    helpers.InitializeOauth2Configuration()
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
// @Param product body ProductDTO true "Create product dto"
// @Success 200 {object} ProductDTO
// @Router /product [post]
func CreateNewProduct(w http.ResponseWriter, r *http.Request){
    services.CreateNewProduct(w,r)
   }

// @Summary Get specific Product
// @Description 
// @Tags products
// @Produce  json
// @Param id path string true "Get product by id"
// @Success 200 {object} ProductDTO
// @Router /product/{id} [get]
func GetProduct(w http.ResponseWriter, r *http.Request){
    services.ReturnSingleProduct(w,r)
   }


// @Summary Update Product
// @Description 
// @Tags products
// @Produce  json
// @Param id path string true "Get product by id"
// @Param product body ProductDTO true "Update product dto"
// @Success 200 {object} ProductDTO
// @Router /product/{id} [put]
func UpdateProduct(w http.ResponseWriter, r *http.Request){
    services.UpdateProduct(w,r)
   }

// @Summary Delete specific Product
// @Description 
// @Tags products
// @Produce  json
// @Param id path string true "Delete product by id"
// @Success 200
// @Router /product/{id} [delete]
func DeleteProduct(w http.ResponseWriter, r *http.Request){
    services.DeleteProduct(w,r)
   }

// @Summary Create User
// @Description 
// @Tags users
// @Produce  json
// @Param user body UserDTO true "Create user dto"
// @Success 200 {object} UserDTO
// @Router /user [post]
func CreateNewUser(w http.ResponseWriter, r *http.Request){
    services.CreateNewUser(w,r)
   }

// @Summary Get user by authcode
// @Description 
// @Tags users
// @Produce  json
// @Param authcode path string true "Get user by authcode"
// @Success 200 {object} UserDTO
// @Router /user/{authcode} [get]
func GetUserByAuthCode(w http.ResponseWriter, r *http.Request){
    services.GetUserByAuthCode(w,r)
   }

// @Summary Verify user
// @Description 
// @Tags users
// @Produce  json
// @Param id path string true "Get user by id"
// @Success 200 {object} UserDTO
// @Router /user/verification/{id} [get]
func VerifyUserEmail(w http.ResponseWriter, r *http.Request){
    services.VerifyUserEmail(w,r)
   }  

// @Summary Get all users
// @Description 
// @Tags users
// @Produce  json
// @Success 200
// @Router /users [get]
func ReturnAllUsers(w http.ResponseWriter, r *http.Request){
    services.ReturnAllUsers(w,r)
   }

// @Summary html Home Page that have link for google log-in Oauth2
// @Description 
// @Produce  json
// @Success 200
// @Router / [get]
func HomePage(w http.ResponseWriter, r *http.Request){
    services.HomePage(w,r)
}

// @Summary Login User
// @Description 
// @Tags users
// @Produce json
// @Param user body UserDTO true "Login user dto"
// @Success 200 {object} UserDTO
// @Router /user/login [post]
func LoginUserWithPassword(w http.ResponseWriter, r *http.Request){
    services.LoginUserWithPassword(w,r)
}

// @Summary Login user via google account
// @Description 
// @Tags users
// @Produce  json
// @Success 200
// @Router /user/loginViaGoogle [get]
func LoginUserViaGoogle(w http.ResponseWriter, r *http.Request){
  services.LoginUserViaGoogle(w,r)
}

// @Summary Google send user data 
// @Description 
// @Tags 3rdParty
// @Produce  json
// @Success 200
// @Router /googlecallback [get]
func HandleGoogleCallback(w http.ResponseWriter, r *http.Request){
    services.HandleGoogleCallback(w,r)
}
