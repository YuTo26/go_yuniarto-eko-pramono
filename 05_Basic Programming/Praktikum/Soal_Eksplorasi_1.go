package main

import (
	"fmt"
)

func main() {
	//deklarasi
	var kata,nama string
	var i, lengt int
	var flag = false
	//proeses
	fmt.Println("=====<PROGRAM PALINDROM>=====")
	fmt.Printf("masuan namamu : ")
	fmt.Scanf("%s\n",&nama)
	fmt.Printf("masukan kata : ")
	fmt.Scanf("%s\n", &kata)
	fmt.Printf("susunan kata = %s\n", kata)
	lengt = len(kata)
	//looping
	for i = 0; i < lengt; i++ {
		if kata[i] != kata[lengt-i-1] {
			flag = true
			break
		}
	}
	//decision
	if flag {
		fmt.Println("=====OUTPUT=====")
		fmt.Printf("hallo %s\n",nama)
		fmt.Printf("kata [%s] bukan palindrome",kata)
	} else {
		fmt.Println("=====OUTPUT=====")
		fmt.Printf("hallo %s\n",nama)
		fmt.Printf("kata [%s] adalah palindrome",kata)
	}
	return
}
