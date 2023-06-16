package main

import "fmt"


func main() {
	var result int
	result = add(5, 2)
	fmt.Println("your result is: ", result)
}

func add(num1,num2 int) int {
	res:= num1 + num2
	return res
}