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
		myRouter.HandleFunc("/migration", services.CreateDatabaseSchema).Methods("POST")
		myRouter.HandleFunc("/products", services.ReturnAllProducts).Methods("GET")
		myRouter.HandleFunc("/product", services.CreateNewProduct).Methods("POST")
		// myRouter.HandleFunc("/product/{id}", updateProduct).Methods("PUT")
		// myRouter.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")
		myRouter.HandleFunc("/product/{id}",services.ReturnSingleProduct).Methods("GET")
		myRouter.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
		log.Fatal(http.ListenAndServe(":9000", myRouter))

		/*
		myRouter.HandleFunc("/", homePage).Methods("GET")
	
		myRouter.HandleFunc("/user", createNewUser).Methods("POST")
		myRouter.HandleFunc("/user/loginViaGoogle", loginUserViaGoogle).Methods("GET")
		myRouter.HandleFunc("/user/login", loginUserWithPassword).Methods("POST")
		myRouter.HandleFunc("/users", returnAllUsers).Methods("GET")
		myRouter.HandleFunc("/googlecallback", handleGoogleCallback).Methods("GET")
		*/
	 }
	 