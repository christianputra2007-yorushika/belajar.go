package main

import "fmt"

const TARIF_PER_KG = 6500
const BIAYA_ANTAR = 3000

func main() {
	fmt.Println("Christian paja| NIM: 42630013")

	var berat float64
	var uangDiterima int

	fmt.Print("Berat cucian (kg): ")
	fmt.Scan(&berat)

	if berat <= 0 {
		fmt.Println("Berat cucian tidak valid!")
		return
	}

	biayaCuci := int(berat*TARIF_PER_KG + 0.5)
	total := biayaCuci + BIAYA_ANTAR
	totalBayar := (total + 99) / 100 * 100
	fmt.Print("Uang diterima (Rp): ")
	fmt.Scan(&uangDiterima)

	if uangDiterima < 0 {
		fmt.Println("Uang yang dimasukkan tidak valid!")
		return
	}

	if uangDiterima < totalBayar {
		fmt.Println("Uang pembayaran tidak mencukupi!")
		return
	}

	kembalian := uangDiterima - totalBayar

	fmt.Println("===== STRUK LAUNDRY =====")
	fmt.Printf("Berat cucian    : %.1f kg\n", berat)
	fmt.Printf("Tarif per kg    : Rp%d\n", TARIF_PER_KG)
	fmt.Printf("Biaya cuci      : Rp%d\n", biayaCuci)
	fmt.Printf("Biaya antar     : Rp%d\n", BIAYA_ANTAR)
	fmt.Printf("Total           : Rp%d\n", total)
	fmt.Printf("Total bayar     : Rp%d\n", totalBayar)
	fmt.Printf("Uang diterima   : Rp%d\n", uangDiterima)
	fmt.Printf("Kembalian       : Rp%d\n", kembalian)
}
