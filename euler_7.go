package main

// Euler 7
// "Find the 10001st prime."
// Lösung: 104743
func euler7() int {
	erzeugePrimzahl := primeNumber()
	primzahl := 0
	for i := 1; i <= 10001; i++ {
		primzahl = erzeugePrimzahl()
	}
	return primzahl
}
