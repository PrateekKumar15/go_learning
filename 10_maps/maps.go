package main

import (
	"fmt"
	"maps"
)

// maps -> hash, objects, key-value pairs
func main() {
//   crate a map
myMap := make(map[string]string)

myMap["name"] = "John"
myMap["age"] = "Thirty"

fmt.Println(myMap["name"]) // Output: John
fmt.Println(myMap)
fmt.Println(myMap["age"])  // Output: 30
// if the key does not exist, it returns the zero value for the value type
fmt.Println(myMap["phone"])

mymap2 := make(map[string]int)
mymap2["one"] = 1
mymap2["two"] = 2	
fmt.Println(mymap2["one"]) // Output: 1
fmt.Println(mymap2["two"]) // Output: 2
fmt.Println(len(mymap2)) // Output: 2

clear(myMap) // clear the map
fmt.Println(myMap) // Output: map[] (empty map)
//  other way to initiate map
myMap3 := map[string]int{"one": 1, "two": 2}
fmt.Println(myMap3) // Output: map[one:1 two:2]

k, ok := myMap3["one"]  // the k represents the vlue corresponding to the key "one"
// ok is a boolean that indicates if the key exists in the map or not
// if the key exists, ok is true, otherwise it is false

if ok {
	fmt.Println(("OK"))
	}else {
		fmt.Println(("Not OK"))
	}
	fmt.Println(k, ok) // Output: 1 true
	
	myMap4 := map[string]int{"one": 1, "two": 2}
	
	fmt.Println(maps.Equal(myMap3, myMap4)) // Output: true)
	
	
}