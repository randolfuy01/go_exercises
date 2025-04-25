package main

import (
	"fmt"
)

func main() {
	var string1 string = "Hello world!\n"
	var string2 string = "Goodbye world!\n"
	fmt.Print(string1)
	fmt.Print(string2)

	// Using type inference
	string3 := "Hello mars!\n"
	fmt.Print(string3)

	// Testing function calls for addition and subtract_numbers
	var num1 int = 7
	var num2 int = 4

	var result1 int = add_numbers(num1, num2)
	var result2 int = subtract_numbers(num1, num2)

	addition_result_string := fmt.Sprintf("Addition result: %d\n", result1)
	subtraction_result_string := fmt.Sprintf("Subtraction result: %d\n", result2)
	fmt.Print(addition_result_string)
	fmt.Print(subtraction_result_string)
}

func add_numbers(a int, b int) int {
	result := a + b
	return result
}

func subtract_numbers(a int, b int) (result int) {
	result = a - b
	return
}
