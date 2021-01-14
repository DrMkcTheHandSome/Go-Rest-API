![GitHub Logo](Golang.png)

## JWT

###### What is JWT? 

> JWT or JSON web token is a digitally signed string used to securely transmit information between parties. It’s an RFC7519 standard.

> A JWT consists of three parts:
```
header.payload.signature
```
> Below is a sample JWT.
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

* **The header** is a Base64-encoded string and it contains the token type (JWT in this case) and the signing algorithm (HMAC SHA256 in this case, or HS256 for short).

```
{
  "alg": "HS256",
  "typ": "JWT"
}
```

* **The payload** is a Base64-encoded string that contains claims. Claims are a collection of data related to the user and the token itself. Example claims are: exp (expiration time), iat (issued at), name (user name), and sub (subject).
```
{
  "sub": "1234567890",
  "name": "John Doe",
  "iat": 1516239022
}
```
* **The signature** is a signed string. For HMAC signing algorithms, we use the Base64-encoded header, the Base64-encoded payload, and a signing secret to create it.

```
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)
```


###### HOW JWT Works?

> Let's say You have an API consumed by a client (web, mobile app, CLI, etc.) and you want to protect it and > only authorize registered users to access it.

> The solution would be using JSON web tokens (JWT for short) to log in and authorize users as explained in the simple image below

![JWT Image](jwt.png)

1. The client will log the user in by sending the credentials to the API server.
2. The API server will validate the user credentials, sign a JWT, and return it in the HTTP response.
3. The client will use the received JWT to access API resources.
4. The API server will validate the JWT and authorize the user to access the resource.

#### References:
 https://learn.vonage.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr/
 https://medium.com/better-programming/hands-on-with-jwt-in-golang-8c986d1bb4c0
 https://www.sohamkamani.com/golang/jwt-authentication/
 
 #### Installation
 
 ```
 go get github.com/dgrijalva/jwt-go
 ```
 
 #### Example Code for generating JWT Token
```
// JwtWrapper wraps the signing key and the issuer
type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// JwtClaim adds email as a claim to the token
type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

func initJWT(w http.ResponseWriter, r *http.Request,user User){
		jwtWrapper :=  &JwtWrapper {
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	
  // a token that expires in 24 hours
  	expirationTime := time.Now().Add(1440 * time.Minute)
	
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
	})
	
	json.NewEncoder(w).Encode(tokenString)
}
``` 


## Bcrypt
#### Installation

```
go get "golang.org/x/crypto/bcrypt"

```

#### Example usage

````
func hashPassword(password string) string {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "failed generate bcrypt password"
	}
    
	var hash_password string = ""
	hash_password = string(bytes)

	return hash_password
}
  // Validate password when login
  func checkPassword(userPasswordfromDB,providedPassword string) error {
	  err := bcrypt.CompareHashAndPassword([]byte(userPasswordfromDB), []byte(providedPassword))
	  if err != nil {
		  return err
	  }

	  return nil
 }
````
