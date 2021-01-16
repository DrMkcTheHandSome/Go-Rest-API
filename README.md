![GitHub Logo](Golang.png)

### [GOLANG CONNECT TO SQL SERVER](https://gorm.io/docs/connecting_to_the_database.html)

```
import (
  "gorm.io/driver/sqlserver"
  "gorm.io/gorm"
)

// github.com/denisenkom/go-mssqldb
dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
```

### Connection string format
```
// for windows authentication
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
// for SQL Server authentication
connectionString := "sqlserver://username:password@host:port?database=nameDB"
```
> For more info kindly check this [denisenkom/go-mssqldb](https://github.com/denisenkom/go-mssqldb) 

### [Table Migration](https://gorm.io/docs/migration.html)
```
type Product struct {
 gorm.Model
 Code string `gorm:"column:code"`
 Price uint  `gorm:"column:price"`
}

type User struct{
 gorm.Model
 Email string    `json:"email" gorm:"unique"` 
 Password string `json:"password"`
}

connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
 db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
 
    // Migrate the schema
	db.Migrator().CreateTable(&Product{})
	db.Migrator().CreateTable(&User{})	
	
```

### CRUD OPERATION IN GO to SQL SERVER
```
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
// CREATE 
db.Exec("INSERT INTO products (created_at,code,price) VALUES (?,?,?)",time.Now(), "code",100)
// READ ALL 
	var products []Product
    db.Exec("select * from products").Scan(&products)
// READ SPECIFIC DATA
    var product Product
    db.Exec("select * from products where id = ?",1).Scan(&product)
// UPDATE
db.Exec("UPDATE products SET code=?, price = ? WHERE id = ?", "code", 200, 1)
// DELETE
 db.Exec("DELETE FROM products WHERE id = ?", 1)
```
> For more info kindly check the following
* [SQL injection Methods](https://gorm.io/docs/security.html)  
* [GORM CRUD Interface](https://gorm.io/docs/)
