package main

import "fmt"

func main() {
	//deklarasi variable dan tipe data
	var i, nilai int
	//proses input
	fmt.Println("=====<PORGRAM FAKTOR BILGANAN>=====")
	fmt.Printf("masukan angka : ")
	fmt.Scanf("%d\n", &nilai)
	//output dengan decision
	fmt.Printf("faktor dari bilangan %d\n", nilai)
	for i = nilai; i > 0; i-- {
		if nilai%i == 0 {
			fmt.Println(i)
		}
	}
}
