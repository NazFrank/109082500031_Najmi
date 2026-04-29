package main
import "fmt"

func main() {
	var (
    x, y, i, jumlahWadah, wadah, ikanDiWadah int
    berat, total [1000]float64
	)
	fmt.Print("Masukkan banyak ikan yang akan dijual (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan banyak ikan per wadah (y): ")
	fmt.Scan(&y)

	for i = 0; i < x; i++ {
		fmt.Printf("Masukkan berat ikan ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	jumlahWadah = x / y
	if x%y != 0 {
		jumlahWadah = jumlahWadah + 1
	}

	for i = 0; i < x; i++ {
		wadah = i / y
		total[wadah] = total[wadah] + berat[i]
	}

	fmt.Print("Total berat per wadah    : ")
	for i = 0; i < jumlahWadah; i++ {
		if i > 0 {
			fmt.Print("  ")
		}
		fmt.Printf("%.2f", total[i])
	}
	fmt.Println()

	fmt.Print("Rata-rata berat per wadah: ")
	for i = 0; i < jumlahWadah; i++ {
		if (i+1)*y <= x {
			ikanDiWadah = y
		} else {
			ikanDiWadah = x - i*y
		}
		if i > 0 {
			fmt.Print("  ")
		}
		fmt.Printf("%.2f", total[i]/float64(ikanDiWadah))
	}
	fmt.Println()
}