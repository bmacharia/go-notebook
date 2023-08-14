package main

func main() {

	// declare a variable of type int with the value of 10
	count := 10

	//display the "value of" and "address of" count
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	//Pass thge value of count
	increment(count)

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

}

func increment(inc int) {
	// increment the "value of" inc
	inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}
