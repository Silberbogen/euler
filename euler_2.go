package main

// Euler 2
// "By considering the terms in the Fibonacci sequenzuence whose values do not exceed four million, find the sum of the even-valued terms."
// Summe: 4613732
func euler2() int {
	var summe int
	fib := fibonacci()
	for x := fib(); x < 4000000; x = fib() {
		if (x % 2) == 0 {
			summe += x
		}
	}
	return summe
}
