package main

import (
	"math"
)

// Here are two solutions
// The first one (disabled) is the more beautiful,
// but the second (enabled) solution is the much faster one
// "Calculate the sum of all the primes below two million."
// Solution: 142913828922
func euler10() (summe int64) {
	/*
			erzeugePrimzahl := primeNumber()
		for primZahl := erzeugePrimzahl(); primZahl < 2000000; primZahl = erzeugePrimzahl() {
			summe += int64(primZahl)
		}
		return summe 
	*/
	maximum := 1999999
	var sieb [2000000]bool
	for i, _ := range sieb {
		sieb[i] = true
	}
	sieb[0] = false
	sieb[1] = false
	for i, max := 2, int(math.Sqrt(float64(maximum)))+1; i <= max; i++ {
		if sieb[i] {
			for j := 2 * i; j <= maximum; j += i {
				sieb[j] = false
			}
		}
	}
	for i, v := range sieb {
		if v {
			summe += int64(i)
		}
	}
	return summe
}
