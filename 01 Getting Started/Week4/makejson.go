package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Addr string `json:"address"`
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	println("Enter your name")
	name, err := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	println("Enter your address")
	addr, err := reader.ReadString('\n')
	addr = strings.TrimSpace(addr)

	person := new(Person)

	person.Name = name
	person.Addr = addr

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile("person.json", jsonData, 0777)
	if err != nil {
		fmt.Println(err)
	}
}
