package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	result := add(2, 4)
	fmt.Printf("Result : %v\n", result)
}
