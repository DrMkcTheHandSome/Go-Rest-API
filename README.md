![GitHub Logo](Golang.png)

### Go Rest API N-tier Architecture

> for example, Local module model contains only model.go which has the following content

```
package models

type Example struct {
    Name string
}

func (e *Example) Foo() string {
    return e.Name
}

```
> For this local module generate module in the path of package models

 ```
 go mod init models
 ```

 > Then you will get this file, go.mod that contains

 ```
module models
go 1.13
 ```

> After that, generate go mod again in the path of package main. then add the following

```
require (
    models v1.0.0
)

replace models v1.0.0 => ./models
```

> So the go.mod main file would be like this

```
module main

go 1.15

require (
    models v1.0.0
)

replace models v1.0.0 => ./models
```

###### Reference
* https://stackoverflow.com/questions/60919877/go-import-local-package 

###### Note: First letter of the functions or variables must be capital if you want to re-use it from another packages
* https://stackoverflow.com/questions/24487943/invoking-struct-function-gives-cannot-refer-to-unexported-field-or-method  
