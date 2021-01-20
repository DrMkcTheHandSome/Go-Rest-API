package models

type GoogleAuthResponse struct {
	Id string `json:"id"` 
	Email string `json:"email"` 
	IsEmailVerified bool `json:"verified_email"` 
}