// euler
package main

import (
	"fmt"
)

func main() {
	fmt.Println("1. Euler Task")
	fmt.Println("Add all the natural numbers below one thousand that are multiples of 3 or 5.")
	sum1, list1 := euler1()
	fmt.Println("List of all numbers:", list1)
	fmt.Println("Sum:", sum1)
	fmt.Println()
	fmt.Println("2. Euler Task")
	fmt.Println("By considering the terms in the Fibonacci sequenzuence whose values do not exceed four million, find the sum of the even-valued terms.")
	fmt.Println("Summe:", euler2())
	fmt.Println()
	fmt.Println("3. Euler Task")
	fmt.Println("Find the largest prime factor of a composite number: 600851475143")
	list3 := euler3(600851475143)
	fmt.Println("Prime factors:", list3)
	fmt.Println("Largest prime factor:", list3[len(list3)-1])
	fmt.Println()
	fmt.Println("4. Euler Task")
	fmt.Println("Find the largest palindrome made from the produkt of two 3-digit numbers.")
	fmt.Println("Largest palindrome:", euler4())
	fmt.Println()
	fmt.Println("5. Euler Task")
	fmt.Println("What is the smallest number divisible by each of the numbers 1 to 20?")
	fmt.Println("Smallest number:", euler5())
	fmt.Println()
	fmt.Println("6. Euler Task")
	sum6_1, sum6_2, diff6 := euler6()
	fmt.Printf("(1+2+...+100)² = %d\n1²+2²+...+100² = %d\nDifferenz = %d\n", sum6_1, sum6_2, diff6)
	fmt.Println()
	fmt.Println("7. Euler Task")
	fmt.Println("Find the 10001st prime.")
	fmt.Println("The 10001st prime is", euler7())
	fmt.Println()
	fmt.Println("8. Euler Task")
	fmt.Println("Discover the largest produkt of five consecutive digits in the 1000-digit number:")
	fmt.Println("7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450")
	max8, list8 := euler8()
	fmt.Print("Solution: ")
	for i := 0; i <= 4; i++ {
		fmt.Printf("%c * ", list8[i])
	}
	fmt.Println("=", max8)
	fmt.Println()
	fmt.Println("9. Euler Task")
	fmt.Println("Find the only Pythagorean triplet, {a, b, c}, for which a + b + c = 1000.")
	a9, b9, c9 := euler9()
	fmt.Printf("Pythagoras: %d²+%d²=%d² entspricht %d+%d=%d\nLösung: %d+%d+%d=%d\n",
		a9, b9, c9, a9*a9, b9*b9, c9*c9, a9, b9, c9, a9+b9+c9)
	fmt.Println()
	fmt.Println("10. Euler Task")
	fmt.Println("Calculate the sum of all the primes below two million.")
	fmt.Println("Solution:", euler10())
	fmt.Println()
	fmt.Println("11. Euler Task")
	fmt.Println("What is the greatest product of four adjacent numbers on the same straight line in the 20 by 20 grid?")
	liste11, maximum11 := euler11()
	fmt.Println("Solution:", maximum11, liste11)
	fmt.Println()

}

// Fibonacci
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// Palindrom
func palindrome(folge string) bool {
	// convert number to string for positional tests
	half := len(folge) / 2
	for i := 0; i < half; i++ {
		if folge[i] != folge[len(folge)-1-i] {
			return false
		}
	}
	return true
}

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

// Creates prime numbers
func primeNumber() func() int {
	var (
		primz           = 2
		oldprimz        int
		listePrimzahlen []int
	)
	return func() int {
		if primz == 2 {
			listePrimzahlen = append(listePrimzahlen, 2)
			oldprimz, primz = primz, 3
		} else if primz == 3 {
			listePrimzahlen = append(listePrimzahlen, 3)
			oldprimz, primz = primz, 5
		} else {
			for istPrimzahl := false; istPrimzahl != true; {
				istPrimzahl = true
				for _, p := range listePrimzahlen {
					if (primz % p) == 0 {
						istPrimzahl = false
						break
					}
				}
				if istPrimzahl {
					listePrimzahlen = append(listePrimzahlen, primz)
					oldprimz = primz
				}
				primz += 2
			}
		}
		return oldprimz
	}
}
