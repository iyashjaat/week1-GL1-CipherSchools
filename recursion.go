package main

import "fmt"

func main() {
	rec(4)
}
func rec(num int) {
	if num <= 0 {
		return
	}

	rec(num - 1)
	rec(num - 2)
	fmt.Println(num - 1)
	// if num%2 == 0 {
	// 	// fmt.Println(num + 1)
	// 	rec(num - 1)
	// 	fmt.Println(num - 1)
	// } else {
	// 	// fmt.Println(num + 2)
	// 	rec(num - 1)
	// 	fmt.Println(num - 1)
	// }
	// // fmt.Println(num - 1)
}

func fib(number int) int {
	if number == 0 || number == 1 {
		return number
	}
	result := fib(number-1) + fib(number-2)
	return result
}

func fact(number int) int {
	if number == 1 || number == 0 {
		return 1
	}
	if number < 0 {
		fmt.Println("Invalid number")
		return -1
	}
	result := number * fact(number-1)
	return result
}
