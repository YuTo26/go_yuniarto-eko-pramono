package main

import "fmt"

func pascalTriangle(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	} else if numRows == 1 {
		return [][]int{[]int{1}}
	}

	prev := pascalTriangle(numRows - 1)
	lastRow := prev[len(prev)-1]
	newRow := make([]int, len(lastRow)+1)
	newRow[0], newRow[len(newRow)-1] = 1, 1
	for i := 1; i < len(newRow)-1; i++ {
		newRow[i] = lastRow[i-1] + lastRow[i]
	}
	return append(prev, newRow)
}

func main() {
	var input int

	fmt.Println("=====>PASCAL TRIANGLE<=====")
	fmt.Printf("masukan number : ")
	fmt.Scanf("%d", &input)
	hasil := pascalTriangle(input)
	fmt.Println(hasil)

}
