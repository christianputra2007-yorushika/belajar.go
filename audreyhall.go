package main

import "fmt"

func main() {
	fmt.Println("Nama: Christian Paja | NIM: 42630013")

	N := 4
	x := 1
	total := 0
	for i := 1; i <= N; i++ {
		if i%2 == 0 {
			x = x * 2
		} else {
			x = x + 3
		}
		total += x
		fmt.Printf("%d ", x)
	}
	fmt.Println()
	fmt.Print(total)
}
