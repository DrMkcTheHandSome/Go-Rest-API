package repositories

import(
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/sqlserver"
	connections "connections"
	 entities "entities"
	 "time"
	)

	type Product struct {
		gorm.Model 
		Code string `gorm:"column:code"`
		Price uint  `gorm:"column:price"`
		}

	func SchemaMigration(){
		fmt.Println("services SchemaMigration")

		db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
        if err != nil {
            panic("failed to connect database")
        }
     
        //Migrate the schema
        db.Migrator().CreateTable(&entities.Product{})
        db.Migrator().CreateTable(&entities.User{})	
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

	func CreateNewUser(user entities.User, hash_password string) entities.User {
		fmt.Println("repository CreateNewUser")

		db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
		if err != nil {
           panic("failed to connect database")
		   }

	    db.Exec("INSERT INTO users (created_at,email,password,is_email_verified) VALUES (?,?,?,?)",time.Now(), user.Email,hash_password,false)
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