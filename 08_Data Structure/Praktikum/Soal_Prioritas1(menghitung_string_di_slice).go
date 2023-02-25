package main

import "fmt"

func Mapping(slice []string) map[string]int {
	m := make(map[string]int)
	for _, v := range slice {
		m[v]++
	}
	return m
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
}
