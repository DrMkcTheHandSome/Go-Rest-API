# GoLang

### SETUP ENVIRONMENT
Download and Install [Go](https://golang.org/doc/install)


### HOW TO RUN GO? Navigate to the path of file in command prompt then type this snippet

```
go run filename.go
```

### BASIC STRUCTURE of HTTP Server
```
import (
    "fmt"
    "log"
    "net/http"
)
 
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}
 
func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8000", nil))
}
 
func main() {
    handleRequests()
}
```


### GO STRUCT JSON Response

```
type Article struct {
    ID     string `json:"Id"`
    Title  string `json:"Title"`
    Author string `json:"author"`
    Link   string `json:"link"`
}
```

> The call to **JSON.NewEncoder(w).Encode("object/array")** does the job of encoding our  array into a JSON string and then writing as part of our response.


### GoLang GORILLA MUX Router
> The name mux stands for **HTTP request multiplexer**. Like the standard http.ServeMux, mux.Router matches incoming requests against a list of registered routes
> and calls a handler for the route that matches the URL or other conditions.

##### To include our gorilla/mux package. Install it in cmd using this snippet.

```
go get -u github.com/gorilla/mux
```

##### Thereafter, import gorilla mux package. 
> Example Snippet:

```
package main;
import(
"github.com/gorilla/mux"
)

func handleRequests(){
   myRouter := mux.NewRouter().StrictSlash(true)
   myRouter.HandleFunc("/", homePage)
   myRouter.HandleFunc("/articles", returnAllArticles).Methods("GET")
   myRouter.HandleFunc("/article/{id}",returnSingleArticle).Methods("GET")
   myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
   myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
   log.Fatal(http.ListenAndServe(":9000", myRouter))
}

```

### GoLang G-ORM Package
>The GORM is fantastic **ORM library** for Golang, aims to be developer friendly. It is an ORM library for dealing with relational databases

##### Installing G-ORM Snippet
```
go get gorm.io/gorm
```
##### Create Db file in windows command prompt run this snippet
```
fsutil file createnew filename.db 0
```
##### Mind that gorm uses **SQLite**. You need to install **GCC** on your computer using a **ming-w installer** 
##### Installing SQLITE with GCC
   1. Install http://msys2.github.io/ **Read the instruction carefully**
   2. launch mingw64_shell.bat or mingw32_shell.bat
   4. pacman -S mingw-w64-x86_64-gcc
   5. Go to Program Files and click on "MinGW Command Prompt". This will open a console with the correct environment for using MinGW with GCC.
   6. Within this console, navigate to your GOPATH.
   7. Enter the following commands:
      * go get -u github.com/mattn/go-sqlite3
      * go install github.com/mattn/go-sqlite3
   
##### Fixed GORM Errors using this References:
* https://github.com/mattn/go-sqlite3/issues/435
* https://github.com/mattn/go-sqlite3/issues/212 

### GORM CRUD Operation Example Snippet
```
package main;

import(
"gorm.io/gorm"
"gorm.io/driver/sqlite"
)

func main(){
   db, err := gorm.Open(sqlite.Open("practice.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    
    var product Product 
    var products []Product

    // Migrate the schema
    db.AutoMigrate(&Product{})
    
    // Create
    db.Create(&Product{Code: "P1", Price: 100})
    //Read 
      // Return Single row
    db.First(&product,"code = ?",key)
      // Return Multiple rows
    db.Find(&products)
      //Update multiple columns
      db.Model(&product).Where("id = ?", 1).Updates(Product{Code: "p2", Price: 120})
       // Delete Single Row
          db.Delete(&product,1)

}

```

##### GORM Queries for more info read [GORM Document](https://gorm.io/docs/query.html) 


