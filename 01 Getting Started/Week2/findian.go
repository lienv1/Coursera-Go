package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func executeFindian() {

	fmt.Println("Enter a string")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Not Found!")
		return
	}

	input = input[:len(input)-2]
	input = strings.ToLower(input)

	if !strings.HasPrefix(input, "i") {
		fmt.Println("Not Found!")
		return
	}
	if !strings.HasSuffix(input, "n") {
		fmt.Println("Not Found!")
		return
	}
	if !strings.Contains(input, "a") {
		fmt.Println("Not Found!")
		return
	}
	fmt.Println("Found!")

}
