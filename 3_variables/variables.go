package main
import "fmt"


func main() {
	//  we need to declare a variable type at time of deceleration if we do not assign a value to it at the time of deceleration
var name string = "Prateek" // string
// var name = "John Doe" // string
var age int = 30 // int

// shorthand syntax

// name := "John Doe" // string 

fmt.Println("Name:", name)
fmt.Println("Age:", age)
}