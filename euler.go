// euler
package main

import (
	"fmt"
	"strconv" // Euler 4, 8
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
	fmt.Print("Lösung: ")
	for i := 0; i <= 4; i++ {
		fmt.Printf("%c * ", list8[i])
	}
	fmt.Println("=", max8)
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

// Euler 1
// "Add all the natural numbers below one thousand that are multiples of 3 or 5."
// Liste aller Zahlen: [3, 5, 6, 9, 10, 12, 15, 18, 20, 21, 24, 25, 27, 30, 33, 35, 36, 39, 40, 42, 45, 48, 50, 51, 54, 55, 57, 60, 63, 65, 66, 69, 70, 72, 75, 78, 80, 81, 84, 85, 87, 90, 93, 95, 96, 99, 100, 102, 105, 108, 110, 111, 114, 115, 117, 120, 123, 125, 126, 129, 130, 132, 135, 138, 140, 141, 144, 145, 147, 150, 153, 155, 156, 159, 160, 162, 165, 168, 170, 171, 174, 175, 177, 180, 183, 185, 186, 189, 190, 192, 195, 198, 200, 201, 204, 205, 207, 210, 213, 215, 216, 219, 220, 222, 225, 228, 230, 231, 234, 235, 237, 240, 243, 245, 246, 249, 250, 252, 255, 258, 260, 261, 264, 265, 267, 270, 273, 275, 276, 279, 280, 282, 285, 288, 290, 291, 294, 295, 297, 300, 303, 305, 306, 309, 310, 312, 315, 318, 320, 321, 324, 325, 327, 330, 333, 335, 336, 339, 340, 342, 345, 348, 350, 351, 354, 355, 357, 360, 363, 365, 366, 369, 370, 372, 375, 378, 380, 381, 384, 385, 387, 390, 393, 395, 396, 399, 400, 402, 405, 408, 410, 411, 414, 415, 417, 420, 423, 425, 426, 429, 430, 432, 435, 438, 440, 441, 444, 445, 447, 450, 453, 455, 456, 459, 460, 462, 465, 468, 470, 471, 474, 475, 477, 480, 483, 485, 486, 489, 490, 492, 495, 498, 500, 501, 504, 505, 507, 510, 513, 515, 516, 519, 520, 522, 525, 528, 530, 531, 534, 535, 537, 540, 543, 545, 546, 549, 550, 552, 555, 558, 560, 561, 564, 565, 567, 570, 573, 575, 576, 579, 580, 582, 585, 588, 590, 591, 594, 595, 597, 600, 603, 605, 606, 609, 610, 612, 615, 618, 620, 621, 624, 625, 627, 630, 633, 635, 636, 639, 640, 642, 645, 648, 650, 651, 654, 655, 657, 660, 663, 665, 666, 669, 670, 672, 675, 678, 680, 681, 684, 685, 687, 690, 693, 695, 696, 699, 700, 702, 705, 708, 710, 711, 714, 715, 717, 720, 723, 725, 726, 729, 730, 732, 735, 738, 740, 741, 744, 745, 747, 750, 753, 755, 756, 759, 760, 762, 765, 768, 770, 771, 774, 775, 777, 780, 783, 785, 786, 789, 790, 792, 795, 798, 800, 801, 804, 805, 807, 810, 813, 815, 816, 819, 820, 822, 825, 828, 830, 831, 834, 835, 837, 840, 843, 845, 846, 849, 850, 852, 855, 858, 860, 861, 864, 865, 867, 870, 873, 875, 876, 879, 880, 882, 885, 888, 890, 891, 894, 895, 897, 900, 903, 905, 906, 909, 910, 912, 915, 918, 920, 921, 924, 925, 927, 930, 933, 935, 936, 939, 940, 942, 945, 948, 950, 951, 954, 955, 957, 960, 963, 965, 966, 969, 970, 972, 975, 978, 980, 981, 984, 985, 987, 990, 993, 995, 996, 999]
// Summe: 233168
func euler1() (int, []int) {
	var (
		summe int
		liste []int
	)
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			summe += i
			liste = append(liste, i)
		}
	}
	return summe, liste
}

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

// Euler 3
// "Find the largest prime factor of a composite number."
// Liste der Primfaktoren: [71 839 1471 6857]
// Größter Primfaktor: 6857
func euler3(maximum int64) []int64 {
	var liste []int64
	for i := int64(2); i <= maximum; i++ {
		for (maximum % i) == 0 {
			maximum /= i
			liste = append(liste, i)
		}
	}
	return liste
}

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

// Euler 7
// "Find the 10001st prime."
// Lösung: 104743
func euler7() int {
	erzeugePrimzahl := primeNumber()
	primzahl := 0
	for i := 1; i <= 10001; i++ {
		primzahl = erzeugePrimzahl()
	}
	return primzahl
}

// Euler 8
// "Discover the largest produkt of five consecutive digits in the 1000-digit number."
// Lösung: 9 * 9 * 8 * 7 * 9 = 40824
func euler8() (int, string) {
	const zahl = "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"
	größteZahl := 0
	faktorenListe := ""
	for i := range zahl {
		// Bei weniger als 5 Zeichen, Schleife beenden
		if i > len(zahl)-5 {
			break
		}
		// Umwandeln der nächsten 5 Zeichen in eine Zahl
		teilZahl, _ := strconv.Atoi(zahl[i : i+5])
		neueZahl := 1
		for j := 1; j <= 5; j++ {
			neueZahl *= teilZahl % 10
			teilZahl /= 10
		}
		// Ist die neue Zahl größer als die bisherige größte Zahl?
		if neueZahl > größteZahl {
			größteZahl = neueZahl
			faktorenListe = zahl[i : i+5]
		}
	}
	return größteZahl, faktorenListe
}
