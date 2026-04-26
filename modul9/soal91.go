package main

import "fmt"

type titik struct {
	x, y int
}

type lingkaran struct {
	pusat  titik
	radius int
}

func jarak(p, q titik) float64 {
	var (
		dx, dy, val, z float64 = float64(p.x - q.x), float64(p.y - q.y), 0 , 0
		i int
	)
	val = dx*dx + dy*dy
	if val == 0 {
		return 0
	}
	z = val
	for i = 0; i < 100; i++ {
		z = (z + val/z) / 2
	}
	return z
}

func didalam(c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {
	var (
		l1     lingkaran
		l2     lingkaran
		p      titik
		dalam1, dalam2 bool
	)
	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
	fmt.Scan(&p.x, &p.y)
	dalam1 = didalam(l1, p)
	dalam2 = didalam(l2, p)
	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}