package main
import "fmt"
func main() {


	// while loop
	// while loop is not a keyword in golang
	i:=1
	for i<=10 {
		fmt.Println(i)
		i++
	}
	// for loop
	for i:=1; i<=10; i++ {
		if i%2==0 {
			continue // skip even numbers
			
		}
		fmt.Println(i)
	}

// range loop
	// range loop is used to iterate over a collection
	// range loop is not a keyword in golang
	for i := range 10 {
		fmt.Println(i)
	}

}