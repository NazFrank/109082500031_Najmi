package main

import "fmt"

func hitungSkor(soal, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0
	for i := 1; i <= 8; i++ {
		fmt.Scan(&waktu)
		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var soal, skor, maxSoal, minSkor int

	maxSoal = -1
	minSkor = 0

	fmt.Scan(&nama)
	for nama != "Selesai" {
		hitungSkor(&soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			pemenang = nama
			maxSoal = soal
			minSkor = skor
		}

		fmt.Scan(&nama)
	}

	fmt.Println(pemenang, maxSoal, minSkor)
}