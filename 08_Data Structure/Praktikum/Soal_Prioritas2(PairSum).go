package main

import "fmt"

func PairSum(arr []int, target int) []int {
	pairs := []int{}
	for index, i := 0; i < len(arr); i++ {
		for index_2, j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				pairs = append(pairs, arr[index], arr[index_2])
			}
		}
	}
	return pairs
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))
}
