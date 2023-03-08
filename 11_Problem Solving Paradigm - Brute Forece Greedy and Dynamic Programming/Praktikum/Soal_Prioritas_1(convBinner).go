package main

import (
	"fmt"
)

// deklarasi konversi binner dengan looping
// penggunaan quote %b = pengubah nilai dari integer menjadi binner
func konversiBinner(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%b ", i)
	}
	fmt.Printf("%b", n)
}

func main() {
	konversiBinner(3) // 0 1 10 11
}
