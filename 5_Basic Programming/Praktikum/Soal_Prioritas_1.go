package main

import "fmt"

func main() {
	//deklarasi variable dan tipe data
	var luas, sisiA, sisiB, tinggi float64
	//input proses
	fmt.Println("=====>PROGRAM LUAS TRAPESIUM<=====")
	fmt.Printf("masukan nilai sisi atas : ")
	fmt.Scanf("%f\n", &sisiA)
	fmt.Printf("masukan nilai sisi bawah : ")
	fmt.Scanf("%f\n", &sisiB)
	fmt.Printf("masukan nilai tinggi : ")
	fmt.Scanf("%f\n", &tinggi)
	//statement
	luas = (0.5 * (sisiA + sisiB) * tinggi)
	//output
	fmt.Println("=====>OUTPUT PROGRAM<=====")
	fmt.Printf("Rumus Luas Trapesium =\n(L = 1/2x(a+b)xt)\nDiketahui :\n")
	fmt.Printf("sisi atas = %v cm\nsisi bawah = %v cm\ndan tinggi = %v cm\n",sisiA,sisiB,tinggi)
	fmt.Printf("jadi luasnya adalah L = 1/2 x (%v+%v)x%v = %v cm\n",sisiA,sisiB,tinggi)

}
