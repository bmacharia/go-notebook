package main

import "fmt"

func main() {

	var age int
	var name string
	var legal bool

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(legal)

	month := 10
	dayOfWeek := "Wednesday"
	happy := false

	fmt.Println(month)
	fmt.Println(dayOfWeek)
	fmt.Println(happy)

	//Perform a type conversion
	pi := float32(3.14)

	fmt.Printf("%T [%v]\n", pi, pi)

}
