package main

import "fmt"

type tabel [127]rune

func isiArray(t *tabel, n *int) {
	var token string
	*n = 0
	for {
		fmt.Scan(&token)
		if token == "." || *n >= 127 {
			break
		}
		t[*n] = rune(token[0])
		(*n)++
	}
}

func cetakArray(t tabel, n int) {
	var i int = 0
	for i = 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	var (
		i int = 0
		j int = n - 1
	)
	for i < j {
		t[i], t[j] = t[j], t[i]
		i++
		j--
	}
}

func palindrom(t tabel, n int) bool {
	var (
		salinan tabel
		i       int = 0
	)
	for i = 0; i < n; i++ {
		salinan[i] = t[i]
	}
	balikanArray(&salinan, n)
	for i = 0; i < n; i++ {
		if t[i] != salinan[i] {
			return false
		}
	}
	return true
}

func main() {
	var (
		tab     tabel
		m       int = 0
		isPalin bool
	)

	fmt.Print("Teks\t\t: ")
	isiArray(&tab, &m)

	isPalin = palindrom(tab, m)

	balikanArray(&tab, m)

	fmt.Print("Reverse teks\t: ")
	cetakArray(tab, m)

	fmt.Printf("Palindrom\t? %v\n", isPalin)
}