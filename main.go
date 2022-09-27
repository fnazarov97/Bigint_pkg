package main

import (
	"fmt"
	bigint "less08/bigint"
)

func main() {

	var str1, str2 string
	fmt.Println("Enter numbers for ADD:")
	fmt.Println("number1:")
	fmt.Scan(&str1)
	fmt.Println("number2:")
	fmt.Scan(&str2)
	num1, _ := bigint.NewInt(str1)
	num2, _ := bigint.NewInt(str2)

	resolve := bigint.Mod(num1, num2)
	fmt.Println(resolve)

	//ADD ...
	// result := bigint.Add(num1, num2)
	// fmt.Printf("Result: %s\n", result.Value)
	// //SUBTRACTION ...
	// result = bigint.Sub(num1, num2)
	// fmt.Printf("Result: %s\n", result.Value)
	//MULTIPLY ...

}
