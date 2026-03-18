package main
import "fmt"

func jarak(a, b, c, d float64) float64 {
	var dx, dy float64
	dx = a - c
	dy = b - d
	return dx*dx + dy*dy
}

func didalam(cx, cy, r, x, y float64) bool {
	var hasil bool
	hasil = jarak(x, y, cx, cy) <= r*r
	return hasil
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64
	var in1, in2 bool

	// input
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	in1 = didalam(cx1, cy1, r1, x, y)
	in2 = didalam(cx2, cy2, r2, x, y)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}