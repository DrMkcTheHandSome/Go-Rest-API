package controllers

import(
	"github.com/gorilla/mux"
	"log"
	"net/http"
    "github.com/gorilla/handlers"
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
		myRouter.HandleFunc("/scene", services.CreateScene).Methods("POST")
		myRouter.HandleFunc("/scene/{label}",services.GetSceneByLabel).Methods("GET")
		myRouter.HandleFunc("/scene/{id}", services.UpdateScene).Methods("PUT")
		myRouter.HandleFunc("/googlecallback", services.HandleGoogleCallback).Methods("GET")
		myRouter.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
		headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
        originsOk := handlers.AllowedOrigins([]string{"*","http://127.0.0.1:8000","http://localhost:8000","http://localhost:9000","http://localhost:9000/products"})
        methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
		log.Fatal(http.ListenAndServe(":9000", handlers.CORS(originsOk, headersOk, methodsOk)(myRouter)))
		
	 }
	 