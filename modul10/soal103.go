package main
import "fmt"

type arrBalita [100]float64
var n int

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	var i int

	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i = 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita) float64 {
	var (
		i int
		total float64
	)
	for i = 0; i < n; i++ {
		total = total + arrBerat[i]
	}

	return total / float64(n)
}

func main() {
	var (
		i int
		minVal, maxVal, avg float64
		data arrBalita
	)

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	hitungMinMax(data, &minVal, &maxVal)
	avg = rerata(data)

	fmt.Printf("Berat balita minimum: %.2f kg\n", minVal)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", maxVal)
	fmt.Printf("Rerata berat balita: %.2f kg\n", avg)
}