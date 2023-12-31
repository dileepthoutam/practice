package main

import "fmt"

func GetAllKeys(sample map[string]string) []string {
	var keys []string

	for key := range sample {
		keys = append(keys, key)
	}

	return keys
}

func IterateOverAMap() {
	sample := map[string]string{
		"a": "x",
		"b": "y",
		"c": "z",
		"d": "p",
	}

	for key, value := range sample {
		fmt.Println(key, value)
	}

	fmt.Println("----------------------")

	for key := range sample {
		fmt.Println(key)
	}

	fmt.Println("----------------------")

	for _, value := range sample {
		fmt.Println(value)
	}

	fmt.Println(GetAllKeys(sample))
}

func main() {
	IterateOverAMap()
}
