package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type GoogleAuthResponse struct {
	Id string `json:"id"` 
	Email string `json:"email"` 
	IsEmailVerified bool `json:"verified_email"` 
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