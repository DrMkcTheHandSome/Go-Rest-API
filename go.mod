module main

go 1.15

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/swaggo/http-swagger v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/oauth2 v0.0.0-20210113205817-d3ed898aa8a3 // indirect
	gorm.io/driver/sqlserver v1.0.5 // indirect
	gorm.io/gorm v1.20.11 // indirect
	docs v1.2.3
)

replace (
	docs v1.2.3 => ./docs
)
