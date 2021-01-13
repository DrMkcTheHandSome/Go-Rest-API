package main;

import(
"encoding/json"
"fmt"
"log"
"net/http"
 "io/ioutil"
"github.com/gorilla/mux"
)
// Global Variables
// Article ...
type Article struct {
    ID     string `json:"Id"`
    Title  string `json:"Title"`
    Author string `json:"author"`
    Link   string `json:"link"`
}
 
// Articles ...
var Articles []Article

// Trigger Functions at start  
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}
 
func handleRequests() {
   initializeRoutes()
   fmt.Println("Hello Go!") 
}
 
func main() { // like ngOnInit in Angular
    initializeValues()
    handleRequests()
}

func initializeRoutes(){
   initRoutesByGorillaMux()
}

func initRoutesByGorillaMux(){
   myRouter := mux.NewRouter().StrictSlash(true)
   myRouter.HandleFunc("/", homePage)
   myRouter.HandleFunc("/articles", returnAllArticles).Methods("GET")
   myRouter.HandleFunc("/article/{id}",returnSingleArticle).Methods("GET")
   myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
   myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
   log.Fatal(http.ListenAndServe(":9000", myRouter))
}

/* 
func initRouteByNetHttp(){
  http.HandleFunc("/", homePage)
  http.HandleFunc("/articles", returnAllArticles)
  log.Fatal(http.ListenAndServe(":9000", nil))
}
*/ 

func initializeValues(){
  Articles = []Article{
        Article{
            ID: "1",
            Title: "Python Intermediate and Advanced 101",
            Author: "Arkaprabha Majumdar",
            Link:   "https://www.amazon.com/dp/B089KVK23P"},
        Article{
            ID: "2",
            Title: "R programming Advanced",
            Author: "Arkaprabha Majumdar",
            Link:   "https://www.amazon.com/dp/B089WH12CR"},
        Article{
            ID: "3",
            Title: "R programming Fundamentals",
            Author: "Arkaprabha Majumdar",
            Link:   "https://www.amazon.com/dp/B089S58WWG"},
    }
}

// LOGIC
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]
 
    for _, article := range Articles {
        if article.ID == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
    json.Unmarshal(reqBody, &article)
    Articles = append(Articles, article)
 
    json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
 
    for index, article := range Articles {
        if article.ID == id {
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }
}