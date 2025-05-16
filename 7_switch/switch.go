package main

import (
	"fmt"
	// "time"
)

func main() {
	// simple switch
// 	i := 6 

// 	switch i {
// 	case 1:{
// 		fmt.Println("One")
// 	}
// 	case 2:{
// 		fmt.Println("Two")
// 	}
// 	case 3:{
// 		fmt.Println("Three")
// 	}
// 	case 4:{
// 		fmt.Println("FOur")
// 	}
// 	case 5:{
// 		fmt.Println("Five")
// 	}
// 	default:{
// 		fmt.Println("enter a valid no.")
// 	}
// }

//  multiple condition switch

// switch time.Now().Weekday(){
// case time.Saturday, time.Sunday:
// 	fmt.Println("It's a weekend")
// default:
// 	fmt.Println("It's a weekday")
// }

// type switch

whoAmI := func (i any)  {
	switch t := i.(type){
	case int:
		fmt.Println("I am an int", i)
	case string:
		fmt.Println("I am a string", i)
	case bool:
		fmt.Println("I am a bool", i)
	case nil:
		fmt.Println("I am a nil", i)
	default:
		fmt.Println("I am a unknown type", t) 

	}
}

whoAmI(1)
whoAmI("hello")
whoAmI(true)
whoAmI(nil)
whoAmI(3.14)	


}