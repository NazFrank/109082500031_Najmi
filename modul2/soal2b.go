package main
import "fmt"

func main() {
	var (
		in1, in2, in3, in4 string
		i int
		stats bool = true
	)
	for i = 1 ; i <= 5; i++ {
		fmt.Printf("Percobaan ke-%d: ", i)
		fmt.Scanln(&in1, &in2, &in3, &in4)
		if in1 != "merah" || in2 != "kuning" || in3 != "hijau" || in4 != "ungu" {
			stats = false
		}
	}
	fmt.Printf("BERHASIL: %t", stats)

}