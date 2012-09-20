package main

import (
	"fmt"
)

func main() {
/*
 *   "What is the difference between the sum of the squares and the square of the sums?"
*/
    summe := 0
    quadratsumme := 0
    for i := 1; i <= 100; i++ {
        summe += i
        quadratsumme += i * i
	}
    summenquadrat := summe * summe
    fmt.Println("(1+2+...+100)² =", summenquadrat)
    fmt.Println("1²+2²+...+100² =", quadratsumme)
    fmt.Println(summenquadrat, "-", quadratsumme, "=", summenquadrat-quadratsumme)
}
// (1+2+...+100)² = 25502500
// 1²+2²+...+100² = 338350
// 25502500 - 338350 = 25164150
