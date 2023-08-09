package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
	age   int
}

func main() {
	mbuki := user{
		name:  "Mbuki",
		email: "mbuki@ngong.com",
		age:   23,
	}

	fmt.Println("Name", mbuki.name)
	fmt.Println("Email", mbuki.email)
	fmt.Println("Age", mbuki.age)

	// Declare a variable of an anonymous struct
	creed := struct {
		name  string
		email string
		age   int
	}{
		name:  "Creed",
		email: "creed@tracy.com",
		age:   5,
	}
	//Display the values
	fmt.Println("Name", creed.name)
	fmt.Println("Email", creed.email)
	fmt.Println("Age", creed.age)

}
