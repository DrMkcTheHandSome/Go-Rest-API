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

	func GetAllProducts()  {
		fmt.Println("services GetAllProducts")
		// db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
		// if err != nil {
		// 	panic("failed to connect database")
		//  }
		
		// var products []Product
		// // Get all records
		// db.Exec("select * from dbo.products").Scan(&products)
		// fmt.Println(products)
		connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
    db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		 fmt.Println("failed to connect database") 
        panic("failed to connect database")
     }
	
    // Get all records


	var product Product
    db.Raw("select * from products where id = 1").Scan(&product)
	fmt.Println(product)
	fmt.Println("=========")

	db.Exec("INSERT INTO products (created_at,code,price) VALUES (?,?,?)",time.Now(), "TEST",9000)

	var products1 []Product
	db.Raw("select * from products").Scan(&products1)
	}