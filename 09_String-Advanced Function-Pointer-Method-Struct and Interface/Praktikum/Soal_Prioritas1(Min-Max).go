package main

import "fmt"

func getMinMax(numbers ...*int) (min int, max int) {
	if len(numbers) > 0 {
		min = *numbers[0]
		max = *numbers[0]
		for _, num := range numbers {
			if *num < min {
				min = *num
			}
			if *num > max {
				max = *num
			}
		}
	}
	return
}

// visualgo.net
// https://opendsa-server.cs.vt.edu/embed/mergesortAV
func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Println("=====>PENENTUAN MIN-MAX LIST NUMBER<=====")
	fmt.Println("Silahkan maskan angka : ")
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Printf("nilai terkecil yang di dapat adalah [ %d ]\n", min)
	fmt.Printf("nilai terbesarg yang di dapat adalah [ %d ]\n", max)
}
