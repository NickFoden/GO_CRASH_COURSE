package main

import "fmt"

func main() {

	// var name = "Nick"
	var age int32 = 37
	const isCool = true

	// name := "Nick"
	size := 1.3
	// email := "nick@gmail.com"

	name, email := "Brad", "brad@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
