package main

import (

"fmt"
"strconv"
"time"

)

const constantTest string = "New Zealand"
// const constantTest := "New Zealand" // assignation operator := not working in constant 

const (
	PRODUCT  = "Mobile"
	QUANTITY = 50
	PRICE    = 50.50
	STOCK  = true
) // Multiple constant declaration


func getGoLangBasic() {
fmt.Println("Hello, World")

var firstname string = "Marc Kenneth"
    middlename := " C. " // other type of declaration
var lastname string = "Lomio"


 fmt.Println("===== variable declaration with assigned values =====")
 
 fmt.Println(firstname + middlename + lastname)

}



func getIfElseExample(){
 var number int = 11
var numberToString = strconv.Itoa(number) // for conversion need to import strconv

 fmt.Println("===== condition statement with data type conversion =====")
 // Note: be aware to the formatting in conditional statement because it was case sensitive
 if number == 10 {
  fmt.Println(numberToString + " number is equal to 10")
 } else if number < 10 {
   fmt.Println(numberToString + " number is less than to 10")
 } else { 
      fmt.Println(numberToString + " number is greater than to 10")
 }
}

func getArraysExample(){

  fmt.Println("===== Arrays =====")
  
    fmt.Println("===== Array Declaration =====")
  var arrayA [5]int
  fmt.Println(arrayA)
  
   fmt.Println("===== assigned value in array using index =====")
  arrayA[2] = 1
  fmt.Println(arrayA)
  
  fmt.Println("===== Array Declaration with assigned values =====")
  
  //arrayB := [5]int{1,2,3,4,5} // other type of declaration
 // var arrayB = [5]int{1,2,3,4,5} // inputting size is optional
   var arrayB = []int{1,2,3,4,5} 

  fmt.Println(arrayB)
  
}

func getMapsExample(){
  fmt.Println("===== Maps =====")
  
   vertices := make(map[string]int) // this is like a dictionary that have key value pairs
   
   vertices["iverson"] = 3
   vertices["bryant"] = 24
   vertices["james"] = 23
   
   fmt.Println("===== Print All values =====")
   fmt.Println(vertices)
  
     fmt.Println("===== Print length of map vertices =====")
     fmt.Println(len(vertices))              

   fmt.Println("===== Print Specific value =====")
   fmt.Println(vertices["bryant"])
    
  fmt.Println("===== Delete Specific value (Iverson) in the map key value pairs =====")
  delete(vertices,"iverson")  
  //Issue: Operation did not complete successfully because the file contains a virus or potentially unwanted software.
  // if the above issue occured there's a window security will pop up just click it then allow the file in this device
  
  fmt.Println(vertices)

  newVertices := map[string]float64{           
    "e":  2.71828,
    "pi": 3.1416,
}

//newVertices["pi"] = 3.14             // Add a new key-value pair
//newVertices["pi"] = 3.1416           // Update value

  fmt.Println("===== Map declaration with assigned values =====")
   fmt.Println(newVertices)
}

func getConstantExample(){
  fmt.Println("===== Constant =====")
  fmt.Println(constantTest)
}

func getDateTimeExample(){
  fmt.Println("===== Date Time =====")
  today := time.Now() // Need to import time 
  fmt.Println(today)
}

func getSwitchCaseExample(){
 
  fmt.Println("===== Switch Case =====")
  switchCaseNumber := 10
  
  switch switchCaseNumber {
  case 5:
		fmt.Println("Number 5th. Clean your house.")
  case 10:
		fmt.Println("Number 10th. Buy some wine.")
  default:
		fmt.Println("No information available for that day.")
  }
  
  fmt.Println("===== Multiple Switch Case =====")
 switch switchCaseNumber {
  case 5, 10:
		fmt.Println("Number 5th & 10th. Clean your house.")
  case 15, 20:
		fmt.Println("Number 15th & 20th.  Buy some wine.")
  default:
		fmt.Println("No information available for that day.")
  }
  
   fmt.Println("=====  Switch Case Conditional Statement =====")
   
 switch  {
  case switchCaseNumber == 10: {
  		fmt.Println("switchCaseNumber equal to 10th")
  }
  case switchCaseNumber > 10:
		fmt.Println("switchCaseNumber is greater than to 10th")
  default:
		fmt.Println("No information available for that day.")
  }
}

func getForLoopExample(){
     fmt.Println("=====  For Loop Statement =====")
	for incrementY := 1; incrementY <= 5; incrementY++ {
		fmt.Println(strconv.Itoa(incrementY) + " loop result")
		
		if incrementY == 4 {
		fmt.Println("break the loop. i don't want to reach the incrementY = 5; ")
		break
		}
	}
     
	 fmt.Println("===== Golang - for range Statement =====")

	// Example 1
	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}
 
	// Example 2
	for key := range strDict {
		fmt.Println(key)
	}
 
	// Example 3
	for _, value := range strDict { // _, need this to occur the key values
		fmt.Println(value)
	}
}

func getBasicOperation(num1 int, num2 int){
fmt.Println("===== Function with parameters =====")
 sum := num1 + num2
fmt.Println("Sum result: " + strconv.Itoa(sum) )
}

func getFunctionReturnText() string {
fmt.Println("===== Function with return type =====")
return "Hello, Marc Kenneth!"
}

func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}

func getReturnFunctionMultipleReturn(){
   var a, p int
	a, p = rectangle(20, 30)
	fmt.Println("===== Function with multiple return =====")
	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)
}

func main(){

getGoLangBasic()
getIfElseExample()
getArraysExample()
getMapsExample()
getConstantExample()
getDateTimeExample()
getSwitchCaseExample()
getForLoopExample()
getBasicOperation(10,20)
fmt.Println(getFunctionReturnText())
getReturnFunctionMultipleReturn()
// Reference: https://www.golangprograms.com/go-language.html
}



