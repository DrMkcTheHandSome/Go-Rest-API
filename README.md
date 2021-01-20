![GitHub Logo](Golang.png)

### Go Rest API Swagger

> Swaggo to be simple and hassle-free and can be a good starting point for documenting APIs in Go.  

### Swaggo setup
> install the libraries we are dependent on. Run the following commands from the commandline:

```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/http-swagger
go get -u github.com/alecthomas/template
```

### Generate swagger.json
```
# In your project dir (~/GOPATH/src/swaggo-orders-api normally)
swag init -g orders.go
```

###### The request body is described by the @Param annotation, which has the following syntax:

```
 @Param [param_name] [param_type] [data_type] [required/mandatory] [description]

``` 
###### The param_type can be one of the following values:

1. query (indicates a query param)
2. path (indicates a path param)
3. header (indicates a header param)
4. body
5. formData

###### Example values for model attributes
```
// Order represents the model for an order
type Order struct {
	OrderID      string    `json:"orderId" example:"1"`
	CustomerName string    `json:"customerName" example:"Leo Messi"`
	OrderedAt    time.Time `json:"orderedAt" example:"2019-11-09T21:21:46+00:00"`
	Items        []Item    `json:"items"`
}

// Item represents the model for an item in the order
type Item struct {
	ItemID      string `json:"itemId" example:"A1B2C3"`
	Description string `json:"description" example:"A random description"`
	Quantity    int    `json:"quantity" example:"1"`
}
```

##### References
* https://www.soberkoder.com/swagger-go-api-swaggo/ 
* https://github.com/soberkoder/swaggo-orders-api 
* https://github.com/swaggo/swag  

##### Fixes
* https://gitmemory.com/issue/swaggo/swag/810/717284737 
* https://github.com/swaggo/swag/issues/810  