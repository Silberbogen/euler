package main

import (
	"fmt"
)

var (
	primz = 2
	listePrimzahlen []int
	primzahl int
)

func main() {
/*
 *   "Find the 10001st prime."
 */
	for i := 1; i <= 10001; i++ {
		primzahl = erzeugePrimzahl()
	}
    fmt.Println("Lösung:", primzahl)
}

// Lösung: 104743

func erzeugePrimzahl() int {
	if primz == 2 {
		listePrimzahlen = append(listePrimzahlen, 2)
		primz = 3
		return primz
	}
	if primz == 3 {
		listePrimzahlen = append(listePrimzahlen, 3)
		primz = 5
		return primz
	}
	for istPrimzahl := false; istPrimzahl != true; {
		istPrimzahl = true
		for _,p := range listePrimzahlen {
			if (primz % p) == 0 {
				istPrimzahl = false
				break
			}
		}
		if istPrimzahl {
			listePrimzahlen = append(listePrimzahlen, primz)
		}
		primz += 2
	}
	return primz
}
