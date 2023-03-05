package main

import "fmt"

func caesar(offset int, input string) string {
	var output string
	//97 adalah nilai desimal dari ASCII a
	//dimana nilai dari integer di ubah menjadi char
	for _, char := range input {
		shifted := (int(char) - 97 + offset) % 26
		output += string(rune(shifted + 97))
	}
	return output
}

func main() {
	fmt.Println("=====>ALGORITMA CAESAR CHIPER<=====")
	fmt.Println("nilai asli ( abc ) => " + caesar(3, "abc"))                                                  // def
	fmt.Println("nilai asli ( alta ) => " + caesar(2, "alta"))                                                // cnvc
	fmt.Println("nilai asli ( alterraacademy ) => " + caesar(10, "alterraacademy"))                           // kvdobbkkmknowi
	fmt.Println("nilai asli ( abcdefghijklmnopqrstuvwxyz ) => " + caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println("nilai asli ( abcdefghijklmnopqrstuvwxyz ) => " + caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl

}
