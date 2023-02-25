package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	var mergedArray []string

	for _, element := range arrayA {
		mergedArray = append(mergedArray, element)
	}

	for _, element := range arrayB {
		mergedArray = append(mergedArray, element)
	}

	return mergedArray
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"serge1", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"deveil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"deveil jin", "serge1"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
