package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	sli := make([]int, 0, 3)

	println("Please type some numbers")

	for {
		var input string
		fmt.Scan(&input)
		if input == "X" {
			break
		}
		number, err := strconv.Atoi(input)
		if err != nil {
			println(err)
			continue
		}

		sli = append(sli, number)
		sort.Ints(sli)
		fmt.Println(sli)
	}

}
