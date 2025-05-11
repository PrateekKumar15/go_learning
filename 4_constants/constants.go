package main
import "fmt"


const timezone = "UTC" // string
// name := "John Doe" //  This will not work because we cannot use shorthand syntax =outside main function


func main() {
	// constants	
	// constants are immutable values that are known at compile time and do not change for the life of the program
	// constants can be character, string, boolean, or numeric values
	// constants can be declared using the const keyword
	const pi = 3.14
	const name = "John Doe"
	const age = 30
	const isMarried = false
	const isEmployed = true
	const isStudent = false
	const isRetired = false
	const isAlive = true
	const isDead = false
	const isMale = true
	const isFemale = false
	const isOther = false
	fmt.Println("Pi:", pi)
	fmt.Println("Name:", name)	
	fmt.Println("Age:", age)
	fmt.Println("Is Married:", isMarried)
	fmt.Println("Is Employed:", isEmployed)
	fmt.Println("timezone:", timezone)


const(
	port = 8080 // int
	host = "localhost" // string

)

fmt.Println(host, port)

}