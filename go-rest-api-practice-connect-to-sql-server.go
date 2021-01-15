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
"golang.org/x/crypto/bcrypt"
jwt "github.com/dgrijalva/jwt-go"
"time"
)
// Global Variables
type Product struct {
 gorm.Model
 Code string
 Price uint
}

type User struct {
 gorm.Model
 Email string    `json:"email" gorm:"unique"` 
 Password string `json:"password"`
}

// JwtWrapper wraps the signing key and the issuer
type JwtWrapper struct {
	SecretKey       string
	Issuer          string
}

// JwtClaim adds email as a claim to the token
type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

var jwtKey string = "my_secret_key"

	
	
//TO DO: Refactor follow DRY principle & Delegation Principle

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
    db.AutoMigrate(&User{})

    // Create
    //db.Create(&Product{Code: "P1", Price: 100}) // test
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
   myRouter.HandleFunc("/user", createNewUser).Methods("POST")
   myRouter.HandleFunc("/user/login", loginUser).Methods("POST")
   myRouter.HandleFunc("/users", returnAllUsers).Methods("GET")
   log.Fatal(http.ListenAndServe(":9000", myRouter))
}

// LOGIC
func createNewProduct(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: createNewProduct")
	isAuthorize := AuthenticateCurrentUser(w,r)
         if isAuthorize == nil {
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
}

func returnSingleProduct(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnSingleProduct")
	
	isAuthorize := AuthenticateCurrentUser(w,r)
         if isAuthorize == nil {
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
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllProducts")
	
	   	isAuthorize := AuthenticateCurrentUser(w,r)
         if isAuthorize == nil { 
		 db, err := gorm.Open(sqlite.Open("practice.db"), &gorm.Config{})
    
    if err != nil {
        panic("failed to connect database")
    }
    // Find all of our products.
    var products []Product
    db.Find(&products)
    json.NewEncoder(w).Encode(products)
		 }

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Endpoint Hit: deleteProduct")
  
   	isAuthorize := AuthenticateCurrentUser(w,r)
    
    if isAuthorize == nil {
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
} 

func updateProduct(w http.ResponseWriter, r *http.Request){
 fmt.Println("Endpoint Hit: updateProduct")
 
 	isAuthorize := AuthenticateCurrentUser(w,r)
    
	if isAuthorize == nil {
	 
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

}

func createNewUser(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: createNewUser")
	db, err := gorm.Open(sqlite.Open("practice.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    reqBody, _ := ioutil.ReadAll(r.Body)
    var user User 
	var hash_password string = ""
    json.Unmarshal(reqBody, &user)
    hash_password = hashPassword(user.Password)
    db.Create(&User{Email: user.Email, Password: hash_password})
	user.Password = hash_password
    json.NewEncoder(w).Encode(user)
}

func returnAllUsers(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllUsers")
	
	isAuthorize := AuthenticateCurrentUser(w,r)
	
	if isAuthorize == nil {
	  db, err := gorm.Open(sqlite.Open("practice.db"), &gorm.Config{})
        if err != nil {
          panic("failed to connect database")
         }
	  var users []User
      db.Find(&users)
      json.NewEncoder(w).Encode(users)
	 } 
}

func AuthenticateCurrentUser(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("token")
	
    if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return err
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	
	// Get the JWT string from the cookie
	token_string := cookie.Value
	
	claims := &JwtClaim{}
	
	token, err := jwt.ParseWithClaims(token_string, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return err
		}
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}
	
    fmt.Println("Authorize! " + claims.Email)
    //w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Email)))
	return nil
}

func loginUser(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: loginUser")
	  db, err := gorm.Open(sqlite.Open("practice.db"), &gorm.Config{})
    
    if err != nil {
        panic("failed to connect database")
    }
	
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User 
	var userPayload User
	json.Unmarshal(reqBody, &user)
	userPayload = user
    query := db.Where(&User{Email: user.Email}).First(&user)
	if query.Error == gorm.ErrRecordNotFound {
	    fmt.Println("Login Failed")
	} 
	 err = checkPassword(user.Password,userPayload.Password)
	 if err != nil {
	    fmt.Println("Login Failed")
	 } else {
	    fmt.Println("Login Success")
        initJWT(w,r,user)
	 }
}

func initJWT(w http.ResponseWriter, r *http.Request,user User){
	jwtWrapper :=  &JwtWrapper {
		SecretKey:       jwtKey,
		Issuer:          "AuthService",
	}
	
    // a token that expires in 5 minutes
  expirationTime := time.Now().Add(5 * time.Minute)
	
	claims := &JwtClaim{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    jwtWrapper.Issuer,
		},
	}
	
	// generates token 
   token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
   // Create JWT string
   tokenString, err := token.SignedString([]byte(jwtWrapper.SecretKey))
    if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		Path: "/",
	})
	
	test,errora := r.Cookie("token")
	fmt.Println(test)
    fmt.Println(errora)
	
	json.NewEncoder(w).Encode(tokenString)
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





