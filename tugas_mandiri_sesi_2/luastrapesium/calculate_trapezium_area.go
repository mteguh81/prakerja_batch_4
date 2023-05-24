package main

import (
	"fmt"
)

func areaOfTrapezium(a, b, d float64) float64 {
	// returning the area by applying the formula
	return (1.0 / 2) * ((a + b) * d)
}

func main() {
	var a, b, d, Area float64
	fmt.Println("Program to find the Area of a Trapezium.")

	fmt.Print("Enter number of lower base: ")
	fmt.Scanf("%f\n", &a)
	fmt.Print("Enter number of upper base: ")
	fmt.Scanf("%f\n", &b)
	fmt.Print("Enter number of height: ")
	fmt.Scanf("%f\n", &d)

	// finding the Area of a Trapezium by calling the function with the respective parameter
	Area = areaOfTrapezium(a, b, d)

	// printing the result
	fmt.Print("The Area of a Trapezium whose length of the parallel sides and distance between the sides ", a, ", ", b, ", ", d, " is ", Area, " cm.")
}
