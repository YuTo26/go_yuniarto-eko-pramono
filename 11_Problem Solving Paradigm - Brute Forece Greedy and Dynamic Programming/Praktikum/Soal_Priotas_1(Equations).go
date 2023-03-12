package main

import "fmt"

func SimpleEquations(a, b, c int) {
	for x := 1; x <= 100; x++ {
		for y := x + 1; y <= 100; y++ {
			z := a - x - y
			if z <= y {
				break
			}
			if x*y*z == b && x*x+y*y+z*z == c {
				fmt.Println(x, y, z)
				return
			}
		}
	}
	fmt.Println("no solution")
}

func main() {
	var a, b, c int

	fmt.Printf("masukan nilai a : ")
	fmt.Scanf("%d\n", &a)
	fmt.Printf("masukan nilai b : ")
	fmt.Scanf("%d\n", &b)
	fmt.Printf("masukan nilai c : ")
	fmt.Scanf("%d\n", &c)

	SimpleEquations(a, b, c)
}
