package main

import "fmt"

type employee struct {
	Name   string
	Age    int
	Salary int
}

func main() {

	name := "dileep"

	for index, char := range name {
		fmt.Println(index, string(char))
	}

}
