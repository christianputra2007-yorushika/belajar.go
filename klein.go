package main

import "fmt"

func main() {
	fmt.Println("Nama: Christian Paja | NIM: 42630013")

	var kategori, n int

	fmt.Print("Kode kategori (1/2/3): ")
	fmt.Scan(&kategori)
	fmt.Print("Jumlah pengunjung: ")
	fmt.Scan(&n)

	var namaKategori string
	var tarif int

	switch kategori {
	case 1:
		namaKategori = "Dewasa"
		tarif = 20000
	case 2:
		namaKategori = "Anak"
		tarif = 10000
	case 3:
		namaKategori = "Turis Asing"
		tarif = 100000
	default:
		fmt.Println("Kategori tidak valid")
		return
	}

	if n < 1 || n > 100 {
		fmt.Println("Jumlah tidak valid")
		return
	}

	diskon := 0
	if n >= 2*8 {
		diskon = 20
	} else if n >= 8 {
		diskon = 10
	}

	if kategori == 3 && diskon > 10 {
		diskon = 10
	}

	totalTiket := tarif * n
	totalBayar := totalTiket - totalTiket*diskon/100

	fmt.Println("Kategori      :", namaKategori)
	fmt.Println("Tarif/orang   : Rp", tarif)
	fmt.Println("Jumlah (n)    :", n)
	fmt.Println("Diskon        :", diskon, "%")
	fmt.Println("Total bayar   : Rp", totalBayar)
}
