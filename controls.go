package main

import "fmt"

/*1. if-else statement
if(condition){
	//statements
}
if(condition){
	//statements
}else{
	//statements
}else if{
	//statements
}else{
	//statements
}

2. switch(data){
case var1:
	statement
case var2:
	statement

default:
}

*/

func main() {
	//data := 10
	//fmt.Print(data)
	// fmt.Print("Enter a Number:")
	// var input int
	// fmt.Scanln(&input)

	// if input%2 == 0 {
	// 	fmt.Println("Hey, You are even !!")
	// } else {
	// 	fmt.Println("Hey, You are odd !!")
	// }

	// if x := 10; x%2 == 0 {
	// 	fmt.Println("Hey, You are even !!")
	// } else {
	// 	fmt.Println("Hey, You are odd !!")
	// }

	data := 10
	switch data {
	case 10:
		data = 100
		fmt.Println(data)
		fallthrough // execute the next case also

	case 100:
		data = 103
		fmt.Println(data)
		fallthrough
	case 11:
		data = 104
		fmt.Println(data)
		fallthrough

	case 15:
		data = 1002
		fmt.Println(data)
		fallthrough

	default:
		data = 10909
		fmt.Println(data)
	}
}
