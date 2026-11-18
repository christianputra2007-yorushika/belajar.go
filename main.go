package main

import "fmt"

func main() {
	fmt.Println("Nama: (Christian Paja) | NIM: (42630013)")

	a := 23
	b := 4
	c := 2.5

	p := a / b
	q := a % b
	r := float64(a / b)
	s := float64(a) / float64(b)
	u := int(c*float64(b)) + a%3
	t := (a%2 == 0) && (float64(b) > c)

	fmt.Println("p =", p)
	fmt.Println("q =", q)
	fmt.Println("r =", r)
	fmt.Println("s =", s)
	fmt.Println("u =", u)
	fmt.Println("t =", t)
}
