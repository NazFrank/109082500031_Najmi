package main

import "fmt"

func cetakBaris(n int) {
	if n == 0 {
		return
	}
	fmt.Print("*")
	cetakBaris(n - 1)
}

func polaBintang(baris, n int) {
	if baris > n {
		return
	}
	cetakBaris(baris)
	fmt.Println()
	polaBintang(baris+1, n)
}

func main() {
	var n int
	fmt.Scan(&n)
	polaBintang(1, n)
}