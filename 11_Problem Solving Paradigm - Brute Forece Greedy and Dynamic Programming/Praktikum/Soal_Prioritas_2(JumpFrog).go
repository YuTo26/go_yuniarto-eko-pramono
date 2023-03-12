package main

import "fmt"

func frog(loncat []int) int {
	n := len(loncat)
	strg := make([]int, n)
	strg[0] = 0

	for i := 1; i < n; i++ {
		minCost := strg[i-1] + abs(loncat[i]-loncat[i-1])

		if i > 1 {
			minCost2 := strg[i-2] + abs(loncat[i]-loncat[i-2])
			if minCost2 < minCost {
				minCost = minCost2
			}
		}

		strg[i] = minCost
	}

	return strg[n-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
