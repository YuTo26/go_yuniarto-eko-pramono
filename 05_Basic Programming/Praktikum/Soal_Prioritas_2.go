package main

import "fmt"

func main() {
	//deklarasi variable
	var bilangan int
	//input proses
	fmt.Println("=====>PROGRAM GANIL - GENAP<=====")
	fmt.Printf("masukan bilangan : ")
	fmt.Scanf("%d\n", &bilangan)
	//decision
	if bilangan%2 == 0 {
		fmt.Printf("bilangan %d adalah bilangan genap\n", bilangan)
	} else {
		fmt.Printf("bilangan %d adalah bilangan ganjil", bilangan)
	}

}
