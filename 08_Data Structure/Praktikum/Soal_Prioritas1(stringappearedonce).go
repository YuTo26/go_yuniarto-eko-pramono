package main

import "fmt"

func munculSekali(angka string) []int {
	charAngka := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	notSame := true
	var newAngka []int

	//Iterasi untuk menemukan angka yang tidak berulang
	for i := 0; i < len(angka); i++ {
		notSame = true
		for j := 0; j < len(angka); j++ {
			if angka[i] == angka[j] && i != j {
				notSame = false
				break
			}
		}

		if notSame == true {
			newAngka = append(newAngka, charAngka[string(angka[i])])
		}
	}

	return newAngka
}

func main() {
	//variable & tipe data declaration
	var inputAngka string
	//process declaration
	fmt.Printf("masukan angka : ")
	fmt.Scanf("%s\n", &inputAngka)
	//output with function
	fmt.Println("Output : ", munculSekali(inputAngka))

}
