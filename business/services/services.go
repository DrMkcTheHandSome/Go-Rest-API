package services

import(
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
	repositories "repositories"
	entities "entities"
	connections "connections"
	"github.com/gorilla/mux"
	helpers "helpers"
	globalvariables "globalvariables"
	"golang.org/x/oauth2"
	models "models"
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
	   w.WriteHeader(http.StatusOK)
	}

   func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("services DeleteProduct")

	 vars := mux.Vars(r)
	 key := vars["id"]
  
	 repositories.DeleteProduct(key)	
	 ReturnAllProducts(w,r)
	 w.WriteHeader(http.StatusOK)
  } 


  func CreateNewUser(w http.ResponseWriter, r *http.Request){
    fmt.Println("services createNewUser")

    reqBody, _ := ioutil.ReadAll(r.Body)
    var user entities.User 
	var hash_password string = ""
    json.Unmarshal(reqBody, &user)
	hash_password = helpers.HashPassword(user.Password)
	user = repositories.CreateNewUser(user,hash_password,false)
	user.Password = hash_password
	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusCreated)
}

func ReturnAllUsers(w http.ResponseWriter, r *http.Request){
    fmt.Println("services returnAllUsers")
	
	  var users []entities.User
	
	  users = repositories.GetAllUsers() 
	
	  json.NewEncoder(w).Encode(users)
	  w.WriteHeader(http.StatusOK)
}

func LoginUserWithPassword(w http.ResponseWriter, r *http.Request){
    fmt.Println("services loginUserWithPassword")
    	
	reqBody, _ := ioutil.ReadAll(r.Body)
	
	var user entities.User 
	var userPayload entities.User
	
	json.Unmarshal(reqBody, &user)
	
	userPayload = user
	user = repositories.GetUserByEmail(user.Email)
	
	err := helpers.CheckPassword(user.Password,userPayload.Password)

	 if err != nil {
	    fmt.Println("Login Failed")
	 } else {
		fmt.Println("Login Success")
		w.WriteHeader(http.StatusOK)
     }
}

func LoginUserViaGoogle(w http.ResponseWriter, r *http.Request){
    fmt.Println("services LoginUserViaGoogle")
 
    url := globalvariables.GoogleOauthConfig.AuthCodeURL(globalvariables.OauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGoogleCallback(w http.ResponseWriter, r *http.Request){
    fmt.Println("services handleGoogleCallback")
    content, err := GetUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
    }
    
    var googleAuthResponse models.GoogleAuthResponse 
    if err = json.Unmarshal(content, &googleAuthResponse); err != nil {
        fmt.Println(err)
    } else {
        CreateAuthGoogleUser(googleAuthResponse)
        fmt.Fprintf(w, "Content: %s\n", googleAuthResponse)
    }
}

func GetUserInfo(state string, code string) ([]byte, error) {
	fmt.Println("services GetUserInfo")
    if state != globalvariables.OauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := globalvariables.GoogleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get(connections.GoogleApisOauth2 + token.AccessToken)
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

func CreateAuthGoogleUser(user models.GoogleAuthResponse){
	fmt.Println("services CreateAuthGoogleUser")
    var userFromDB entities.User 
	
	userFromDB = repositories.GetUserByEmail(user.Email)
  
    if userFromDB.Email == "" {
		// If not Exist Create User in DB
		var newUser = entities.User{
			Email: user.Email,
			Password: "",
			IsEmailVerified: user.IsEmailVerified,
		}

		userFromDB = repositories.CreateNewUser(newUser,newUser.Password,newUser.IsEmailVerified) 
    } 
    
    if userFromDB.IsEmailVerified == true {
       // means the user was created and his/her email was verified
       fmt.Println("Google Login Success")
    }
}