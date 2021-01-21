module main

go 1.15

require (
	connections v1.2.3
	constants v1.2.3
	controllers v1.2.3
	docs v1.2.3
	entities v1.2.3
	github.com/codegangsta/negroni v1.0.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/pjebs/jsonerror v0.0.0-20190614034432-63ef9a8df848 // indirect
	github.com/pjebs/restgate v0.0.0-20200504001537-fd9a58a4fe75 // indirect
	github.com/swaggo/http-swagger v1.0.0
	globalvariables v1.2.3
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/oauth2 v0.0.0-20210113205817-d3ed898aa8a3
	gopkg.in/unrolled/render.v1 v1.0.0 // indirect
	gorm.io/driver/sqlserver v1.0.5
	gorm.io/gorm v1.20.11
	helpers v1.2.3
	models v1.2.3
	repositories v1.2.3
	services v1.2.3
)

replace docs v1.2.3 => ./docs

replace entities v1.2.3 => ./data/entities

replace controllers v1.2.3 => ./presentation/controllers

replace services v1.2.3 => ./business/services

replace repositories v1.2.3 => ./data/repositories

replace connections v1.2.3 => ./data/connections

replace constants v1.2.3 => ./business/constants

replace helpers v1.2.3 => ./business/helpers

replace globalvariables v1.2.3 => ./business/globals/globalvariables

replace models v1.2.3 => ./business/models
