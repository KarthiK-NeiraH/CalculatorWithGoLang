package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Enter first number:")
	fmt.Scanln(&num1)

	fmt.Println("Enter operator (+, -, *, /, ^):")
	fmt.Scanln(&operator)

	if operator == "sqrt" {
		fmt.Printf("Square root of %v is %v", num1, math.Sqrt(num1))
		return
	}

	fmt.Println("Enter second number:")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		fmt.Printf("%v + %v = %v", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%v - %v = %v", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%v * %v = %v", num1, num2, num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}
		fmt.Printf("%v / %v = %v", num1, num2, num1/num2)
	case "^":
		fmt.Printf("%v ^ %v = %v", num1, num2, math.Pow(num1, num2))
	default:
		fmt.Println("Invalid operator")
	}
}
