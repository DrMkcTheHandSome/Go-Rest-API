package controllers

import(
	"github.com/gorilla/mux"
	"log"
	"net/http"
	services "services"
	httpSwagger "github.com/swaggo/http-swagger"
	)


	func HandleRequests() {
		InitializeRoutes()
	 }
	 
	 func InitializeRoutes(){
		InitRoutesByGorillaMux()
	 }
	 
	 func InitRoutesByGorillaMux(){
		myRouter := mux.NewRouter().StrictSlash(true)
		myRouter.HandleFunc("/", services.HomePage).Methods("GET")
		myRouter.HandleFunc("/migration", services.CreateDatabaseSchema).Methods("POST")
		myRouter.HandleFunc("/products", services.ReturnAllProducts).Methods("GET")
		myRouter.HandleFunc("/product", services.CreateNewProduct).Methods("POST")
		myRouter.HandleFunc("/product/{id}", services.UpdateProduct).Methods("PUT")
		myRouter.HandleFunc("/product/{id}", services.DeleteProduct).Methods("DELETE")
		myRouter.HandleFunc("/product/{id}",services.ReturnSingleProduct).Methods("GET")
		myRouter.HandleFunc("/user", services.CreateNewUser).Methods("POST")
		myRouter.HandleFunc("/users", services.ReturnAllUsers).Methods("GET")
		myRouter.HandleFunc("/user/verification/{id}",services.VerifyUserEmail).Methods("GET")
		myRouter.HandleFunc("/user/login", services.LoginUserWithPassword).Methods("POST")
		myRouter.HandleFunc("/user/loginViaGoogle", services.LoginUserViaGoogle).Methods("GET")
		myRouter.HandleFunc("/user/{authcode}", services.GetUserByAuthCode).Methods("GET")
		myRouter.HandleFunc("/googlecallback", services.HandleGoogleCallback).Methods("GET")
		myRouter.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
		log.Fatal(http.ListenAndServe(":9000", myRouter))
	 }
	 