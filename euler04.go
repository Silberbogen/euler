package main

import (
	"fmt"
    "strconv"
)

func main() {
/*
 *  "Find the largest palindrome made from the produkt of two 3-digit numbers."
*/
    palindrom := 0
    for i := 999; i >= 100; i-- {
        for j := i; j >= 100; j--{
            zahl := i * j
            // Die Integer zahl in einen String konvertieren
            folge := strconv.Itoa(zahl)
            // Vergleichen aller Positionen des Strings auf Gleichheit
            ist_palindrom := true
            for x := 0; x < (len(folge) / 2); x++ {
				if folge[x] != folge[len(folge) -1 -x] {
					ist_palindrom = false
				}
			}
			if ist_palindrom && zahl > palindrom {
				palindrom = zahl
			}
		}
    }
    fmt.Println("Lösung:", palindrom)
}
// Lösung: 906609
