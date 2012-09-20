package main

/*
 * Euler 03
 * "Find the largest prime factor of a composite number."
*/
import (
	"fmt"
	)
	
func main() {
    i := int64(2)
    zahl := int64(600851475143)
    liste := []int64{} // empty slice

    for i < zahl {
        for (zahl % i) == 0 {
            zahl /= i
            liste = append(liste, i)
        }
        i += 1
    }
    liste  = append(liste, i)
    fmt.Println("Liste der Primfaktoren:", liste)
    fmt.Println("Größter Primfaktor:", liste[len(liste)-1])
}

// Liste der Primfaktoren: [71 839 1471 6857]
// Größter Primfaktor: 6857

