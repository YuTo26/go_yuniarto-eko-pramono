package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {
	if number <= 0 {
		return 0
	}
	count := 0
	for i := 2; ; i++ {
		if isPrime(i) {
			count++
			if count == number {
				return i
			}
		}
	}
}

// fungsi yang untuk membantu prime checking
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i <= int(math.Sqrt(float64(n))); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
