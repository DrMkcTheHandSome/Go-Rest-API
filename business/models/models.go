package models

import(
	"encoding/json"
	)

	type Product struct {
		Code string `json:"code"`
		Price uint  `json:"price"`
		}
		
type User struct{
	Email string    `json:"email"` 
	Password string `json:"password"`
	IsEmailVerified bool `json:"verified_email"`
	}
	
	type GoogleAuthResponse struct {
		Id string `json:"id"` 
		Email string `json:"email"` 
		IsEmailVerified bool `json:"verified_email"` 
	}


	