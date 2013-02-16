package main

// Euler 3
// "Find the largest prime factor of a composite number."
// Liste der Primfaktoren: [71 839 1471 6857]
// Größter Primfaktor: 6857
func euler3(maximum int64) []int64 {
	var liste []int64
	for i := int64(2); i <= maximum; i++ {
		for (maximum % i) == 0 {
			maximum /= i
			liste = append(liste, i)
		}
	}
	return liste
}
