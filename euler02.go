package main

import (
	"fmt"
)

var (
		a = 0
		b = 1
)
		
func main() {
/*
 * "By considering the terms in the Fibonacci sequenzuence whose values do not exceed four million, find the sum of the even-valued terms."
*/
    summe := 0
    for x := fib(); x < 4000000; x = fib()  {
        if (x % 2) == 0 {
            summe += x
		}
	}
    fmt.
    Println("Summe:", summe)
}
// Summe: 4613732

func fib() int {
	a, b = b, a+b
	return a
}
