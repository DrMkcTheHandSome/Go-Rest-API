![GitHub Logo](Golang.png)

### Getting Started with OAuth2 in Go

#### Google Project, OAuth2 keys

> First of all, let’s create our Google OAuth2 keys.

* Go to [Google Cloud Platform](https://console.developers.google.com/)
* Create new project or use an existing one
* Go to Credentials
* Click “Create credentials”
* Choose “OAuth client ID”
* Add authorized redirect URL, in our case it will be

```
localhost:9000/googlecallback
```
* Get client id and client secret
* Save it in a safe place

#### Initial handlers and OAuth2 config. 

###### Install the following

```
go get golang.org/x/oauth2
cloud.google.com/go/compute/metadata
```

###### OAuth Functions

```

type User struct{
 gorm.Model
 Email string    `json:"email" gorm:"unique"` 
 Password string `json:"password"`
 IsEmailVerified bool `json:"verified_email" gorm:"column:is_email_verified"` 
}

func initRoutesByGorillaMux(){
   myRouter := mux.NewRouter().StrictSlash(true)
   myRouter.HandleFunc("/", homePage)
   myRouter.HandleFunc("/user/loginViaGoogle", loginUserViaGoogle).Methods("GET")
   myRouter.HandleFunc("/googlecallback", handleGoogleCallback).Methods("GET")
   log.Fatal(http.ListenAndServe(":9000", myRouter))
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

func loginUserViaGoogle(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: loginUserViaGoogle")
 
    url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
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


```

##### References

* https://itnext.io/getting-started-with-oauth2-in-go-1c692420e03
* https://github.com/douglasmakey/oauth2-example 