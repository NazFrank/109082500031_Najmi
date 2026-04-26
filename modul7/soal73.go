package main
import "fmt"

func main() {
	var (
		hasil          [100]string
		jumlah, skorA, skorB, pertandinganKe,i int = 0,0,0,1,0
		klubA, klubB string
	)

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	for {
		fmt.Printf("Pertandingan %d : ", pertandinganKe)
		fmt.Scan(&skorA, &skorB)
		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[jumlah] = klubA
		} else if skorB > skorA {
			hasil[jumlah] = klubB
		} else {
			hasil[jumlah] = "Draw"
		}
		jumlah++
		pertandinganKe++
	}

	for i = 0; i < jumlah; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil[i])
	}

	fmt.Println("Pertandingan selesai")
}