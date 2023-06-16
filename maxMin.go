package main

import "fmt"

func main() {
	var result int

	// calling function 
	result = minMax(10, 40)

	// printing
	fmt.Println("Max number is : ", result)
}

// creating function

func minMax(num1, num2 int) int {
	var max int
	if(num1> num2) {
		 max = num1
	} else  {
		 max = num2
	}
	return max
	
}