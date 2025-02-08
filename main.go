package main

import "fmt"

//start app
func main() {
	fmt.Println("Expected Output:")
	fmt.Println("Addition: 10+5 =", add(10, 5))
	fmt.Println("Subtraction: 10-5 =", subtract(10, 5))
	fmt.Println("Multipliction: 10*5 =", multiply(10, 5))
	fmt.Println("Divison: 10/5 =", divide(10, 5))
	fmt.Println("Modulo: 10%5 =", modulo(10, 5))
}

func init() {
	fmt.Println("...............Calculator Project............")
}

//Addition a,b
func add(a, b int) int {
	return a + b
}

//Subtraction a,b
func subtract(a, b int) int {
	return a - b
}

//Multipliction a,b
func multiply(a, b int) int {
	return a * b
}

//Divison a,b
func divide(a, b int) int {
	return a / b
}

//Modulo a,b
func modulo(a, b int) int {
	return a % b
}
