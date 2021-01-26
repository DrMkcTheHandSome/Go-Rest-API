package helpers

import(
	"golang.org/x/crypto/bcrypt"
	connections "connections"
	constants "constants"
	globalvariables "globalvariables"
"golang.org/x/oauth2"
"golang.org/x/oauth2/google"
"os"
"math/rand"
)

func HashPassword(password string) string {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "failed generate bcrypt password"
	}
    
	var hash_password string = ""
	hash_password = string(bytes)

	return hash_password
}


  func CheckPassword(userPasswordfromDB,providedPassword string) error {
	  err := bcrypt.CompareHashAndPassword([]byte(userPasswordfromDB), []byte(providedPassword))
	  if err != nil {
		  return err
	  }

	  return nil
 }

 func InitializeOauth2Configuration() {
	os.Setenv(constants.SendGridAPI, connections.SendGridAPI)
	// Setup Google's example test keys
	globalvariables.OauthStateString = RandStringBytes(14)
	os.Setenv(constants.CLIENT_ID, connections.GoogleClientId)
	os.Setenv(constants.SECRET_KEY, connections.GoogleSecretKey)
	globalvariables.GoogleOauthConfig = &oauth2.Config{
	   RedirectURL:  connections.GoogleRedirectURL,
	   ClientID:     os.Getenv(constants.CLIENT_ID),
	   ClientSecret: os.Getenv(constants.SECRET_KEY),
	   Scopes:       []string{connections.GoogleScopes},
	   Endpoint:     google.Endpoint,
   }
}

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = constants.LettersWithNumbers[rand.Intn(len(constants.LettersWithNumbers))]
    }
    return string(b)
}