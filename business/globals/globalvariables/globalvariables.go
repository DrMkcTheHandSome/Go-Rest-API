package globalvariables

import (
"golang.org/x/oauth2"
)

var (
	GoogleOauthConfig *oauth2.Config 
	OauthStateString = ""
	JwtKey  = "my_secret_key"
)