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