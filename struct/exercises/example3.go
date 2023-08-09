// Sample program to show hot to declare and initialize anonymous
// struct types.
package main

import "fmt"

func main() {
	//declare a variable of an anonymous type set to its zero value
	var el struct {
		flag    bool
		counter int16
		pi      float32
	}

	//Display the value
	fmt.Printf("%+v\n", el)

	//Declare a variable of an anonymous type and init using a struct literal
	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	//Display the values
	fmt.Printf("%+v\n", e2)
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}
