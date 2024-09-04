package main

import "fmt"

func main() {
	fmt.Println("HELLO Pointers")

	var ptr *int
	value := 20
	ptr = &value

	fmt.Println("Value of ptr:", &ptr)
}
