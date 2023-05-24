package main

import (
	"fmt"
)

func checkMultipleOfSeven(num int) {
	if num%7 == 0 {
		fmt.Println("Number is multiple of 7.")
		return
	}

	fmt.Println("Number is not multiple of 7.")
	return
}

func main() {
	var num int = 0

	fmt.Println("Program to check multiple of 7.")

	fmt.Print("Enter Number: ")
	fmt.Scanf("%d\n", &num)

	checkMultipleOfSeven(num)
}
