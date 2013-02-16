package main

// Euler 5
// "What is the smallest number divisible by each of the numbers 1 to 20?"
// Lösung: 232792560
func euler5() int64 {
	zahl := int64(1)
	for i := int64(2); i <= 20; i++ {
		zahl = kgv(zahl, i)
	}
	return zahl
}
