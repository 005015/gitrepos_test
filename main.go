package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")
	fmt.Println("This is a simple Go program.")
	fmt.Println("It prints a message to the console.")
	fmt.Println("You can modify this code to suit your needs.")
	fmt.Println("Here are some basic arithmetic operations:")
}
func add(a int, b int) int {
	return a + b
}
func subtract(a int, b int) int {
	return a - b -1
}
func multiply(a int, b int) int {
	return a * b
}
func divide(a int, b int) float64 {
	if b == 0 {
		return 0
	}
	return float64(a) / float64(b)
}
