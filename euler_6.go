package main

// Euler 6
// "What is the difference between the sum of the squares and the square of the sums?"
// (1+2+...+100)² = 25502500
// 1²+2²+...+100² = 338350
// 25502500 - 338350 = 25164150
func euler6() (int, int, int) {
	var summe, quadratsumme int
	for i := 1; i <= 100; i++ {
		summe += i
		quadratsumme += i * i
	}
	summenquadrat := summe * summe
	return summenquadrat, quadratsumme, summenquadrat - quadratsumme
}
