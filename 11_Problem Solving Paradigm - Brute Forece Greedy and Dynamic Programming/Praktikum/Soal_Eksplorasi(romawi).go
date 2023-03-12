package main

import "fmt"

func main() {
	var num int
	roman := ""

	fmt.Println("=====>PROGRAM KONVERSI ANGKA ROMAWI<=====")
	fmt.Printf("masukan angka : ")
	fmt.Scanf("%d", &num)

	// Mendefinisikan array untuk mengkonversi angka ke Romawi
	symbols := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	// Melakukan loop sebanyak simbol yang tersedia
	for _, s := range symbols {
		// Melakukan loop sampai s.value dapat dimasukkan ke dalam num
		for num >= s.value {
			roman += s.symbol // Menambahkan simbol Romawi ke hasil
			num -= s.value    // Mengurangi s.value dari num
		}
	}

	fmt.Println(roman)
}
