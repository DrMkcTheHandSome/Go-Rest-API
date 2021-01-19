package repositories

import(
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/sqlserver"
	connections "connections"
	 entities "entities"
	)

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
		fmt.Println("services GetAllProducts")
		db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		 }
		
		var products []entities.Product
		// Get all records
		db.Exec("SELECT * FROM products").Scan(&products)
		return products
	}