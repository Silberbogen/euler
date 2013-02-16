package main

// Euler 9
// "Find the only Pythagorean triplet, {a, b, c}, for which a + b + c = 1000."
// Pythagoras: 200²+375²=425² entspricht 40000+140625=180625
// Produkt von a*b*c=31875000
// Lösung: 200+375+425=1000
func euler9() (int, int, int) {
	for a := 1; a <= 500; a++ {
		for b := 1; b <= 500; b++ {
			c := 1000 - a - b
			if c > b && a*a+b*b == c*c {
				return a, b, c
			}
		}
	}
	return 0, 0, 0
}
