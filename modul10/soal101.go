package main
import "fmt"

func main() {
	var (
		n, i int
		berat [1000]float64
		minVal, maxVal float64
	)

	fmt.Print("Masukkan banyak anak kelinci: ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	minVal = berat[0]
	maxVal = berat[0]

	for i = 1; i < n; i++ {
		if berat[i] < minVal {
			minVal = berat[i]
		}
		if berat[i] > maxVal {
			maxVal = berat[i]
		}
	}

	fmt.Printf("Berat kelinci terkecil: %.2f\n", minVal)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", maxVal)
}