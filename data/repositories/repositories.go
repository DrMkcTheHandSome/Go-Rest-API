package repositories

import(
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/sqlserver"
	connections "connections"
	 entities "entities"
	 "time"
	)

// TO DO: Refactor

	func SchemaMigration(){
		fmt.Println("services SchemaMigration")

		db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
        if err != nil {
            panic("failed to connect database")
        }
     
        //Migrate the schema
		db.Migrator().CreateTable(&entities.Product{})
		db.Migrator().CreateTable(&entities.User{})
		db.Migrator().CreateTable(&entities.Scene{})	
	}

	func GetAllProducts() []entities.Product  {
		fmt.Println("repository GetAllProducts")
		db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
		if err != nil {
           panic("failed to connect database")
		   }
		   
	var products []entities.Product
	db.Raw("select * from products").Scan(&products)
      return products
	}

	func CreateNewProduct(product entities.Product) entities.Product {
		fmt.Println("repository CreateNewProduct")
		db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
		if err != nil {
           panic("failed to connect database")
		   }
	   db.Exec("INSERT INTO products (created_at,code,price) VALUES (?,?,?)",time.Now(), product.Code,product.Price)
	   return product
	}

	func GetSingleProduct(productId string) entities.Product {
		fmt.Println("repository GetSingleProduct")
		db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
		if err != nil {
           panic("failed to connect database")
		   }
		 var product entities.Product
		 db.Raw("select * from products where id = ?",productId).Scan(&product)  
		 return product
	}

	func UpdateProduct(productId string, product entities.Product) {
		fmt.Println("repository UpdateProduct")
		db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
		if err != nil {
           panic("failed to connect database")
		   }
		 db.Exec("UPDATE products SET code=?, price = ? WHERE id = ?", product.Code, product.Price, productId)
	}

	func DeleteProduct(productId string){
		fmt.Println("repository DeleteProduct")
		db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
		if err != nil {
           panic("failed to connect database")
		   }

		db.Exec("DELETE FROM products WHERE id = ?", productId)
	}

	func CreateNewUser(user entities.User, hash_password string,is_email_verified bool) entities.User {
		fmt.Println("repository CreateNewUser")

		db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
		if err != nil {
           panic("failed to connect database")
		   }

		db.Exec("INSERT INTO users (created_at,email,password,is_email_verified) VALUES (?,?,?,?)",time.Now(), user.Email,hash_password,is_email_verified)
		return user
	}

	
	func GetAllUsers() []entities.User  {
		fmt.Println("repository GetAllUsers")
		db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
		if err != nil {
           panic("failed to connect database")
		   }

	var users []entities.User
	db.Raw("select * from users").Scan(&users)
      return users
	}

	func GetUserByEmail(email string) entities.User {
		fmt.Println("repository GetUserByEmail")
		db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
		if err != nil {
           panic("failed to connect database")
		   }
		 var user entities.User
		 db.Raw("select * from users where email = ?",email).Scan(&user)  
		 return user
	}

	func UpdateUserEmailVerification(userId string) {
		fmt.Println("repository UpdateUserEmailVerification")
		db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
		if err != nil {
           panic("failed to connect database")
		   }
		 db.Exec("UPDATE users SET is_email_verified = 1 WHERE id = ?", userId)
	}

	func UpdateUserAuthCode(email string, auth_code string) {
		fmt.Println("repository UpdateUserAuthCode")
		db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
		if err != nil {
           panic("failed to connect database")
		   }
		 db.Exec("UPDATE users SET auth_code = ? WHERE email = ?", auth_code,email)
	}

	func GetUserByAuthCode(code string) entities.User {
		fmt.Println("repository GetUserByAuthCode")
		db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
		if err != nil {
           panic("failed to connect database")
		   }
		 var user entities.User
		 db.Raw("select * from users where auth_code = ?",code).Scan(&user)  
		 return user
	}

	func CreateScene(scene entities.Scene) entities.Scene {
		fmt.Println("repository CreateScene")
		db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
		if err != nil {
           panic("failed to connect database")
		   }
	   db.Exec("INSERT INTO scenes (created_at,label,value) VALUES (?,?,?)",time.Now(), scene.Label,scene.Value)
	   return scene
	}

