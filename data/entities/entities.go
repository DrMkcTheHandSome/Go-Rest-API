package entities

import(
	"gorm.io/gorm"
	)

  type Product struct {
	gorm.Model 
	Code string `gorm:"column:code"`
	Price uint  `gorm:"column:price"`
	}
	
	type User struct{
	 gorm.Model 
	Email string    `json:"email" gorm:"unique"` 
	Password string `json:"password"`
	IsEmailVerified bool `json:"verified_email" gorm:"column:is_email_verified"` 
	AuthCode string `json:"auth_code" gorm:"column:auth_code"` 
 }