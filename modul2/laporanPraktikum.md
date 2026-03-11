# <h1 align="center">Laporan Praktikum Modul 2 - REVIEW ALGORITMA DAN PEMROGRAMAN 1 </h1>
<p align="center">Muhammad Najmi - 109082500031</p>


## 1. Soal Latihan Modul 2A
### soal2a.go

```go
package main
import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```

### Output 
![Screenshot Output Unguided 1_1]()


## 2. Soal Latihan Modul 2B
### soal2b.go

```go
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
```
### Output:

![Screenshot Output Unguided 1_1]()
[penjelasan]

## 3. Soal Latihan Modul 2C
### soal2c.go

```go
package main
import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
## Output:
![Screenshot Output Unguided 1_1]()
[penjelasan]