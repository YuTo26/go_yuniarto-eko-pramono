package main

import "fmt"

func main() {
	//deklrasai variable
	var nama string
	var absen, tugas, uts, uas, nilai float32
	//input proses
	fmt.Println("=====>PRORGRAM NILAI AKHIR MAHASISWA<======")
	fmt.Println("FORM INPUT DATA =>")
	fmt.Printf("masukan nama : ")
	fmt.Scanf("%s\n", &nama)
	fmt.Printf("masukan nilai absensi : ")
	fmt.Scanf("%f\n", &absen)
	fmt.Printf("masukan nilai tugas : ")
	fmt.Scanf("%f\n", &tugas)
	fmt.Printf("masukan nilai uts : ")
	fmt.Scanf("%f\n", &uts)
	fmt.Printf("masukan nilai uas : ")
	fmt.Scanf("%f\n", &uas)
	//statement
	nilai = ((absen + tugas + uts + uas) / 4)
	//decision proses
	if nilai >= 80 && nilai <= 100 {
		fmt.Printf("===== GRADE NILAI MAHASISWA =====\n")
		fmt.Printf("nama : %s\n ", nama)
		fmt.Printf("Absensi = %v\n Tugas = %v\n UTS = %v\n UAS = %v\n", absen, tugas, uts, uas)
		fmt.Printf("rata-rata nilai = %v\n", nilai)
		fmt.Printf("mendapat nilai A , dan di nyatakan LULUS")
	} else if nilai >= 65 && nilai <= 79 {
		fmt.Printf("===== GRADE NILAI MAHASISWA =====\n")
		fmt.Printf("nama : %s\n ", nama)
		fmt.Printf("Absensi = %v\n Tugas = %v\n UTS = %v\n UAS = %v\n", absen, tugas, uts, uas)
		fmt.Printf("rata-rata nilai = %v\n", nilai)
		fmt.Printf("mendapat nilai B , dan di nyatakan LULUS")
	} else if nilai >= 50 && nilai <= 64 {
		fmt.Printf("===== GRADE NILAI MAHASISWA =====\n")
		fmt.Printf("nama : %s\n ", nama)
		fmt.Printf("Absensi = %v\n Tugas = %v\n UTS = %v\n UAS = %v\n", absen, tugas, uts, uas)
		fmt.Printf("rata-rata nilai = %v\n", nilai)
		fmt.Printf("mendapat nilai C , tolong TINGKATKAN LAGI")
	} else if nilai > 35 && nilai <= 49 {
		fmt.Printf("===== GRADE NILAI MAHASISWA =====\n")
		fmt.Printf("nama : %s\n ", nama)
		fmt.Printf("Absensi = %v\n Tugas = %v\n UTS = %v\n UAS = %v\n", absen, tugas, uts, uas)
		fmt.Printf("rata-rata nilai = %v\n", nilai)
		fmt.Printf("mendapat nilai D , sangat di wajibkan BELAJAR DEENGAN GIAT")
	} else if nilai < 34 {
		fmt.Printf("===== GRADE NILAI MAHASISWA =====\n")
		fmt.Printf("nama : %s\n ", nama)
		fmt.Printf("Absensi = %v\n Tugas = %v\n UTS = %v\n UAS = %v\n", absen, tugas, uts, uas)
		fmt.Printf("rata-rata nilai = %v\n", nilai)
		fmt.Printf("mendapat nilai E , TIDAK PATUT DI TIRU")
	} else {
		fmt.Println("nilai INVALID")
	}

}
