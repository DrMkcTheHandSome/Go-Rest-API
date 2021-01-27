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
	jwt "github.com/dgrijalva/jwt-go"
	constants "constants"
	"time"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
	)

	func HomePage(w http.ResponseWriter, r *http.Request){
		fmt.Println("services homePage")
		isAuthorize := AuthenticateCurrentUser(w,r,globalvariables.JwtKey)
		if isAuthorize == nil { 
			var htmlIndex = `<html>
			<body>
			   <h1>Welcome to the homepage!</h1>
				<a href="/user/loginViaGoogle">Google Log In</a>
			</body>
			</html>`
				fmt.Fprintf(w, htmlIndex)
	   } else {
		   w.WriteHeader(http.StatusUnauthorized)
	   }
	}

func CreateDatabaseSchema(w http.ResponseWriter, r *http.Request){
	fmt.Println("services CreateDatabaseSchema")
	repositories.SchemaMigration()
	w.WriteHeader(http.StatusCreated)
}
	
func ReturnAllProducts(w http.ResponseWriter, r *http.Request) {
     fmt.Println("services ReturnAllProducts")
	
	 isAuthorize := AuthenticateCurrentUser(w,r,globalvariables.JwtKey)
	 if isAuthorize == nil { 
		var products []entities.Product 
		products = repositories.GetAllProducts()
		 json.NewEncoder(w).Encode(products)
		 w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func CreateNewProduct(w http.ResponseWriter, r *http.Request){
    fmt.Println("services CreateNewProduct")

	isAuthorize := AuthenticateCurrentUser(w,r,globalvariables.JwtKey)

	if isAuthorize == nil { 
	
		reqBody, _ := ioutil.ReadAll(r.Body)
		var product entities.Product 
		json.Unmarshal(reqBody, &product)
		product = repositories.CreateNewProduct(product)
		json.NewEncoder(w).Encode(product)
		w.WriteHeader(http.StatusCreated)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func ReturnSingleProduct(w http.ResponseWriter, r *http.Request) {
    fmt.Println("services returnSingleProduct")
	isAuthorize := AuthenticateCurrentUser(w,r,globalvariables.JwtKey)

	if isAuthorize == nil { 
		vars := mux.Vars(r)
		key := vars["id"]
	
		var product entities.Product
			
		product = repositories.GetSingleProduct(key)
	
		json.NewEncoder(w).Encode(product)  
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	fmt.Println("services updateProduct")
	isAuthorize := AuthenticateCurrentUser(w,r,globalvariables.JwtKey)

	   if isAuthorize == nil { 
		vars := mux.Vars(r)
		key := vars["id"]
		reqBody, _ := ioutil.ReadAll(r.Body)
		var product entities.Product 
 
		json.Unmarshal(reqBody, &product)
		repositories.UpdateProduct(key,product)
		product = repositories.GetSingleProduct(key)
		json.NewEncoder(w).Encode(product)
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

   func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("services DeleteProduct")
	isAuthorize := AuthenticateCurrentUser(w,r,globalvariables.JwtKey)

	 if isAuthorize == nil { 
		vars := mux.Vars(r)
		key := vars["id"]
	 
		repositories.DeleteProduct(key)	
		ReturnAllProducts(w,r)
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
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
	user = repositories.GetUserByEmail(user.Email)
	SendEmailVerification(user.Email, fmt.Sprint(user.ID))
	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusCreated)
}

func ReturnAllUsers(w http.ResponseWriter, r *http.Request){
    fmt.Println("services returnAllUsers")
	
	isAuthorize := AuthenticateCurrentUser(w,r,globalvariables.JwtKey)
	
	if isAuthorize == nil { 
		var users []entities.User
	
		users = repositories.GetAllUsers() 
	  
		json.NewEncoder(w).Encode(users)
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
	
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
		w.WriteHeader(http.StatusBadRequest)
	 } else {
		fmt.Println("Login Success")
		globalvariables.JwtKey = "my_secret_key"
		InitJWT(w,r,user,globalvariables.JwtKey)
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
        CreateAuthGoogleUser(w,r,googleAuthResponse)
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

func CreateAuthGoogleUser(w http.ResponseWriter, r *http.Request,user models.GoogleAuthResponse){
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
	   globalvariables.JwtKey = globalvariables.GoogleOauthConfig.ClientSecret
	   InitJWT(w,r,userFromDB,globalvariables.JwtKey)
    }
}

func InitJWT(w http.ResponseWriter, r *http.Request,user entities.User,secretkey string){
	jwtWrapper :=  &models.JwtWrapper {
		SecretKey:   secretkey,
		Issuer:      constants.AuthService,
	}
	
    // a token that expires in 2 hours
  expirationTime := time.Now().Add(120 * time.Minute)
	
	claims := &models.JwtClaim{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    jwtWrapper.Issuer,
		},
	}
	
	// generates token 
   token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
   // Create JWT string
   fmt.Println("Secret Key" + jwtWrapper.SecretKey)
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
	
	json.NewEncoder(w).Encode(tokenString)
}


func AuthenticateCurrentUser(w http.ResponseWriter, r *http.Request, jwtKey string) error {
	
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
	
	claims := &models.JwtClaim{}
	
	token, err := jwt.ParseWithClaims(token_string, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Println(err)
			return err
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return err
	}
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println(err)
		return err
	}
	
    fmt.Println("Authorize! " + claims.Email)
	return nil
}

func SendEmailVerification(email string, id string) {
	from := mail.NewEmail("Marc Kenneth Lomio", "mlomio@blastasia.com")
	subject := "Sending with Twilio SendGrid is Fun"
	to := mail.NewEmail("Test User", email)
	plainTextContent := ""
	htmlContent :=  `<html>
	<body>
	<h1> Welcome to GoRestAPI email using send grid! </h1> 
	<h2> Hi! ` + email + ` </h2>
	<p>Kindly verify your account` + `<a href='http://localhost:9000/user/verification/` + id + `'> <i>here</i> </a></p>
	</body>
	</html>
	`

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv(constants.SendGridAPI))
	response, err := client.Send(message)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func VerifyUserEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	repositories.UpdateUserEmailVerification(key)
	fmt.Println("services VerifiedUserEmail")
		var htmlIndex = `<html>
		<body>
		   <h1>Your email was verified!</h1>
		</body>
		</html>`
	fmt.Fprintf(w, htmlIndex)
}