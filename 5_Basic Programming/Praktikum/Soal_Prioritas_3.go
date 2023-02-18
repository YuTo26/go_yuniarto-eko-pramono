package main

import "fmt"

func main() {
	//deklrasai variable
	var nama string
	var tugas, uts, uas, nilai float64
	//input proses
	fmt.Printf("masukan nama : ")
	fmt.Scanf("%s\n", &nama)
	fmt.Printf("masukan nilai tugas : ")
	fmt.Scanf("%f\n", &tugas)
	fmt.Printf("masukan nilai uts : ")
	fmt.Scanf("%f\n", &uts)
	fmt.Printf("masukan nilai uas : ")
	fmt.Scanf("%f\n", &uas)
	//statement
	nilai = (tugas + uts + uas) / 2.5
	//decision proses
	if nilai >= 80 && nilai <= 100 {
		fmt.Printf("===== GRADE NILAI MAHASISWA =====\n")
		fmt.Printf("nama : %s\n ", nama)
		fmt.Printf("ugas = %v\n UTS = %v\n UAS = %v\n", tugas, uts, uas)
		fmt.Printf("rata-rata nilai = %v\n", nilai)
		fmt.Printf("mendapat nilai A , dan di nyatakan LULUS")
	} else if nilai >= 65 && nilai <= 79 {
		fmt.Printf("===== GRADE NILAI MAHASISWA =====\n")
		fmt.Printf("nama : %s\n ", nama)
		fmt.Printf("Tugas = %v\n UTS = %v\n UAS = %v\n", tugas, uts, uas)
		fmt.Printf("rata-rata nilai = %v\n", nilai)
		fmt.Printf("mendapat nilai B , dan di nyatakan LULUS")
	} else if nilai >= 50 && nilai <= 64 {
		fmt.Printf("===== GRADE NILAI MAHASISWA =====\n")
		fmt.Printf("nama : %s\n ", nama)
		fmt.Printf("Tugas = %v\n UTS = %v\n UAS = %v\n", tugas, uts, uas)
		fmt.Printf("rata-rata nilai = %v\n", nilai)
		fmt.Printf("mendapat nilai C , tolong TINGKATKAN LAGI")
	} else if nilai >= 35 && nilai <= 49 {
		fmt.Printf("===== GRADE NILAI MAHASISWA =====\n")
		fmt.Printf("nama : %s\n ", nama)
		fmt.Printf("Tugas = %v\n UTS = %v\n UAS = %v\n", tugas, uts, uas)
		fmt.Printf("rata-rata nilai = %v\n", nilai)
		fmt.Printf("mendapat nilai D , sangat di wajibkan BELAJAR DEENGAN GIAT")
	} else if nilai == 0 && nilai >= 100 {
		fmt.Printf("nilai INVALID")
	}

}
