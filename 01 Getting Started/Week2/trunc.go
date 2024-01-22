package main

import "fmt"

func executeTrunc() {

	var input float64

	fmt.Println("Please enter a float number")

	fmt.Scan(&input)

	var truncatedInput int = int(input)

	fmt.Println(truncatedInput)
}
