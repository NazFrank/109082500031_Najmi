package main

import "fmt"

func t(a []int, m, x int) {
	var i int
	for i = 0; i < len(a); i++ {
		if m == 0 || (m == 1 && i%2 != 0) || (m == 2 && i%2 == 0) || (m == 3 && i%x == 0) {
			fmt.Print(a[i], " ")
		}
	}
	fmt.Println()
}

func h(a []int, idx int) []int {
	var i int
	for i = idx; i < len(a)-1; i++ {
		a[i] = a[i+1]
	}
	return a[:len(a)-1]
}

func c(a []int) (float64, float64) {
	var i, total int
	var r, j, s float64

	for i = 0; i < len(a); i++ {
		total = total + a[i]
	}
	r = float64(total) / float64(len(a))

	for i = 0; i < len(a); i++ {
		s = float64(a[i]) - r
		j = j + s*s
	}
	return r, j / float64(len(a))
}

func main() {
	var N, i, x, idx int

	fmt.Print("Input jumlah data: ")
	fmt.Scan(&N)

	var a []int
	a = make([]int, N)

	fmt.Print("Input isi array: ")
	for i = 0; i < N; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Print("Input nilai x: ")
	fmt.Scan(&x)

	fmt.Print("Input index yang dihapus: ")
	fmt.Scan(&idx)

	// proses
	t(a, 0, 0)
	t(a, 1, 0)
	t(a, 2, 0)
	t(a, 3, x)

	a = h(a, idx)
	fmt.Println("Setelah dihapus:")
	t(a, 0, 0)

	var r, s float64
	r, s = c(a)
	fmt.Println("Rata-rata:", r)
	fmt.Println("Standar deviasi:", s)
}