package main

import (
    "fmt"
)

func main() {
/*
 *   "What is the smallest number divisible by each of the numbers 1 to 20?"
*/
    zahl := int64(1)
    for i := int64(2); i <= 21; i++ {
        zahl = kgv(zahl, i)
    }
    fmt.Println("Lösung:", zahl)
}
// Lösung: 232792560

// Ermittelt den größten gemeinsamen Teiler zweier Zahlen

func ggt(x, y int64) int64 {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}

// Ermittelt das kleinste gemeinsame Vielfacher zweier Zahlen
func kgv(a, b int64) int64 {
    return a * b / ggt(a, b)
}
