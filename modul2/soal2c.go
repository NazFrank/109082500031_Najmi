package main
import "fmt"

func main()  {
	var berat_parsel, berat_kg, berat_gr, biaya_kg, biaya_gr, total_biaya int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&berat_parsel)
	berat_kg = berat_parsel / 1000
	berat_gr = berat_parsel % 1000
	biaya_kg = berat_kg * 10000
	if berat_gr >= 500{
		biaya_gr = berat_gr * 5
	} else {
		biaya_gr = berat_gr * 15
 	} 
	if berat_kg > 10 {
		total_biaya = biaya_kg 
	} else {
		total_biaya = biaya_kg + biaya_gr
	}
	fmt.Printf("Detail berat: %d kg + %d gr\nDetail biaya: Rp. %d + Rp. %d\nTotal biaya: Rp. %d", berat_kg, berat_gr, biaya_kg, biaya_gr, total_biaya )
}