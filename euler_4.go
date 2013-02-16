package main

import (
    "strconv"
)

// Euler 4
// "Find the largest palindrome made from the produkt of two 3-digit numbers."
// Lösung: 906609
func euler4() int {
	var palindrom int
	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			zahl := i * j
			if palindrome(strconv.Itoa(zahl)) && zahl > palindrom {
				palindrom = zahl
			}
		}
	}
	return palindrom
}
