// Sample program to show how to declare and intialize struct types
package main

import "fmt"

// example represents a type with different fields
type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	//declaare a variable of type example set to its zero value
	var e1 example

	//Display the value
	fmt.Printf("%+v\n", e1)

	//Declare a variable of type example and init using a struct literal
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.14159,
	}

	//Display the field values
	// the dot operator is used to access a filed on a struct
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)

}
