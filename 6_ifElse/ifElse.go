package main
import "fmt"


func main() {
age := 20
if age < 18 {
	fmt.Println("You are a minor.")
} else {
	fmt.Println("You are an adult.")
}
// if else if else
age = 69
if age < 18 {
	fmt.Println("You are a minor.")
} else if age < 65 {
	fmt.Println("You are an adult.")
} else {
	fmt.Println("You are a senior citizen.")
}
var role = "admin"
var hasPermission = true

if role == "admin" || hasPermission {
	fmt.Println("You have admin access.")
} else {
	fmt.Println("You do not have admin access.")
}

if age1:= 20; age1>= 18 {
	fmt.Println("person is an adult", age1)
} else if age1>=12 {
	fmt.Println("person is a teenager", age1 )
}
}

// gO DOES NOT HAVE TERNARY operator we have to use if else 