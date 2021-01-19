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
"os"
"golang.org/x/oauth2"
"golang.org/x/oauth2/google"
"math/rand"
httpSwagger "github.com/swaggo/http-swagger"
)

// TO DO: Refactor

// Product represents the model for an Product
type Product struct {
//  gorm.Model
 Code string `gorm:"column:code"`
 Price uint  `gorm:"column:price"`
}

// User represents the model for an User
type User struct{
//  gorm.Model
 Email string    `json:"email" gorm:"unique"` 
 Password string `json:"password"`
 IsEmailVerified bool `json:"verified_email" gorm:"column:is_email_verified"` 
}

type GoogleAuthResponse struct {
    Id string `json:"id"` 
    Email string `json:"email"` 
    IsEmailVerified bool `json:"verified_email"` 
}

const (
 lettersWithNumbers = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)
var (
    googleOauthConfig *oauth2.Config
    connectionString = "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
    // NOTE: randomize it
	oauthStateString = "pseudo-random"
)



// @title Users Product Go Rest API
// @version 1.0
// @description Go Rest API with SQL SERVER DB
// @contact.name Marc Kenneth Lomio
// @contact.email marckenneth.lomio@gmail.com
// @host localhost:9000
// @BasePath /
func main() { 
    initializeOauth2Configuration()
	handleRequests()
}


func initializeOauth2Configuration(){
     // Setup Google's example test keys
     oauthStateString = RandStringBytes(14)
     os.Setenv("CLIENT_ID", "876220489172-i1msr7n6o01anrcanjg3gqj00h08hain.apps.googleusercontent.com")
     os.Setenv("SECRET_KEY", "H6sWMHe-OiBqC1Nd70prnWvB")
    googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:9000/googlecallback",
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("SECRET_KEY"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
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
   myRouter.HandleFunc("/", homePage).Methods("GET")
   myRouter.HandleFunc("/migration", createDatabaseSchema).Methods("POST")
   myRouter.HandleFunc("/product", createNewProduct).Methods("POST")
   myRouter.HandleFunc("/product/{id}", updateProduct).Methods("PUT")
   myRouter.HandleFunc("/products", returnAllProducts).Methods("GET")
   myRouter.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")
   myRouter.HandleFunc("/product/{id}",returnSingleProduct).Methods("GET")
   myRouter.HandleFunc("/user", createNewUser).Methods("POST")
   myRouter.HandleFunc("/user/loginViaGoogle", loginUserViaGoogle).Methods("GET")
   myRouter.HandleFunc("/user/login", loginUserWithPassword).Methods("POST")
   myRouter.HandleFunc("/users", returnAllUsers).Methods("GET")
   myRouter.HandleFunc("/googlecallback", handleGoogleCallback).Methods("GET")
   myRouter.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
   log.Fatal(http.ListenAndServe(":9000", myRouter))
}

// LOGIC

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = lettersWithNumbers[rand.Intn(len(lettersWithNumbers))]
    }
    return string(b)
}

// @Summary Migrate tables to the SQL Server
// @Description Create database schema
// @Tags migrations
// @Accept  json
// @Produce  json
// @Param user body User true "Create user"
// @Success 200
// @Router /migration [create]
func createDatabaseSchema(w http.ResponseWriter, r *http.Request){
     db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
        if err != nil {
            panic("failed to connect database")
        }
     
        // Migrate the schema
        db.Migrator().CreateTable(&Product{})
        db.Migrator().CreateTable(&User{})	
    }

// homePage godoc
// @Summary show html that navigates to google auth login
// @Description 
// @Produce  json
// @Success 200
// @Router / [get]
func homePage(w http.ResponseWriter, r *http.Request){
    var htmlIndex = `<html>
<body>
   <h1>Welcome to the homepage!</h1>
	<a href="/user/loginViaGoogle">Google Log In</a>
</body>
</html>`
	fmt.Fprintf(w, htmlIndex)
    fmt.Println("Endpoint Hit: homePage")
}

// @Summary Create user product 
// @Description Create the product corresponding by user request
// @Tags products
// @Accept  json
// @Produce  json
// @Param user body User true "Create user"
// @Success 200
// @Router /product [create]
func createNewProduct(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: createNewProduct")
	
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    reqBody, _ := ioutil.ReadAll(r.Body)
    var product Product 
    json.Unmarshal(reqBody, &product)
	db.Exec("INSERT INTO products (created_at,code,price) VALUES (?,?,?)",time.Now(), product.Code,product.Price)
    json.NewEncoder(w).Encode(product)	 
}

// @Summary Update product identified by the given productId
// @Description Update the product corresponding to the input productId
// @Tags products
// @Accept  json
// @Produce  json
// @Param productId path int true "ID of the product to be updated"
// @Success 200
// @Router /product/{id} [update]
func updateProduct(w http.ResponseWriter, r *http.Request){
 fmt.Println("Endpoint Hit: updateProduct")
 
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
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


// returnAllProducts godoc
// @Summary Get details of all products
// @Description Get details of all products
// @Produce  json
// @Success 200 {array} Product
// @Router /products [get]
func returnAllProducts(w http.ResponseWriter, r *http.Request) {
     fmt.Println("Endpoint Hit: returnAllProducts")
	
    db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
     }
	
    // Get all records
	var products []Product
    db.Exec("select * from products").Scan(&products)
	
    json.NewEncoder(w).Encode(products)
}

// @Summary Delete product identified by the given productId
// @Description Delete the product corresponding to the input productId
// @Tags products
// @Accept  json
// @Produce  json
// @Param productId path int true "ID of the product to be deleted"
// @Success 204 "No Content"
// @Router /product/{id} [delete]
func deleteProduct(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Endpoint Hit: deleteProduct")
  
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

   vars := mux.Vars(r)
    key := vars["id"]
    
   db.Exec("DELETE FROM products WHERE id = ?", key)
   returnAllProducts(w,r)
} 

// @Summary retrieve product identified by the given productId
// @Description retrieve the product corresponding to the input productId
// @Tags products
// @Accept  json
// @Produce  json
// @Param productId path int true "ID of the product to be retrieve"
// @Success 200
// @Router /product/{id} [get]
func returnSingleProduct(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnSingleProduct")
	
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
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

// @Summary Login user using google account 
// @Description 
// @Tags login
// @Accept  json
// @Produce  json
// @Param user body User true "Create user"
// @Success 200
// @Router /user/loginViaGoogle [get]
func loginUserViaGoogle(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: loginUserViaGoogle")
 
    url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// @Summary Login user with password 
// @Description 
// @Tags login
// @Accept  json
// @Produce  json
// @Param user body User true "Create user"
// @Success 200
// @Router /user/login [get]
func loginUserWithPassword(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: loginUserWithPassword")
    
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
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

func getUserInfo(state string, code string) ([]byte, error) {
    if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}


// @Summary Get details of the user
// @Description Get details of the user
// @Produce  json
// @Success 200 Google User Info
// @Router /googlecallback [get]
func handleGoogleCallback(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: handleGoogleCallback")
    content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
    }
    
    var googleAuthResponse GoogleAuthResponse 
    if err = json.Unmarshal(content, &googleAuthResponse); err != nil {
        fmt.Println(err)
    } else {
        createAuthGoogleUser(googleAuthResponse)
        fmt.Fprintf(w, "Content: %s\n", googleAuthResponse)
    }
}

func createAuthGoogleUser(user GoogleAuthResponse){
    fmt.Println("Endpoint Hit: createAuthGoogleUser")
    var userFromDB User 
    
    db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.Exec("select * from users where email = ?",user.Email).Scan(&userFromDB)
    if userFromDB.Email == "" {
        // If not Exist Create User in DB
        db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
        if err != nil {
            panic("failed to connect database")
        }
        db.Exec("INSERT INTO users (created_at,email,password,is_email_verified) VALUES (?,?,?,?)",time.Now(), user.Email,"",user.IsEmailVerified)
    } 
    
    if userFromDB.IsEmailVerified == true {
       // means the user was created and his/her email was verified
       fmt.Println("Google Login Success")
    }
}


// @Summary Create new user 
// @Description Create user with email & password
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body User true "Create user"
// @Success 200
// @Router /user [create]
func createNewUser(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: createNewUser")
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var user User 
	var hash_password string = ""
    json.Unmarshal(reqBody, &user)
    hash_password = hashPassword(user.Password)
    db.Exec("INSERT INTO users (created_at,email,password,is_email_verified) VALUES (?,?,?,?)",time.Now(), user.Email,hash_password,false)
    db.Create(&User{Email: user.Email, Password: hash_password})
	user.Password = hash_password
    json.NewEncoder(w).Encode(user)
}

// returnAllUsers godoc
// @Summary Get details of all users
// @Description Get details of all users
// @Produce  json
// @Success 200 {array} User
// @Router /users [get]
func returnAllUsers(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllUsers")
	
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
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





