package main

import (
	"fmt"
)

func Compare(a, b string) string {
	if len(a) < len(b) {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Printf("nilai yang akan di kompare adalah\nAKA\nAKASHI\n")
	fmt.Println("hasilnya adalah = " + Compare("AKA", "AKASHI"))
	fmt.Println("<---------->")
	fmt.Printf("nilai yang akan di kompare adalah\nKANGOORO\nKANG\n")
	fmt.Println("hasilnya adalah = " + Compare("KANGOORO", "KANG"))
	fmt.Println("<---------->")
	fmt.Printf("nilai yang akan di kompare adalah\nKI\nKIJANG\n")
	fmt.Println("hasilnya adalah = " + Compare("KI", "KIJANG"))
	fmt.Println("<---------->")
	fmt.Printf("nilai yang akan di kompare adalah\nKUPU-KUPU\nKUPU\n")
	fmt.Println("hasilnya adalah = " + Compare("KUPU-KUPU", "KUPU"))
	fmt.Println("<---------->")
	fmt.Printf("nilai yang akan di kompare adalah\nILALANG\nILA\n")
	fmt.Println("hasilnya adalah = " + Compare("ILALANG", "ILA"))
}
