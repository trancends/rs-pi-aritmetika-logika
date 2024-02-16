package main

import "fmt"

func main() {
	fmt.Println("========== Soal 1 ==========")
	soal1()
	fmt.Println()
	fmt.Println("========== Soal 2 ==========")
	soal2()
	fmt.Println()
	fmt.Println("========== Soal 3 ==========")
	soal3()
}

func soal1() {
	var allSwitch []int
	var lampOff []int

	for i := 1; i <= 30; i++ {
		switch {
		case i%2 == 0:
			allSwitch = append(allSwitch, i)
		case i%3 == 0:
			allSwitch = append(allSwitch, i)
		case i%5 == 0:
			allSwitch = append(allSwitch, i)
		case i%7 == 0:
			allSwitch = append(allSwitch, i)
		case i%11 == 0:
			allSwitch = append(allSwitch, i)
		default:
			lampOff = append(lampOff, i)
		}
	}

	fmt.Println("Semua Saklar:", allSwitch)
	fmt.Println("Lampu yang mati:", lampOff)
	fmt.Println("Jumlah lampu yang mati:", len(lampOff))
	fmt.Println("Jumlah maksimal lampu yang bisa nyala:", len(allSwitch))
}

func soal2() {
	candies := 12
	grandkids := 3
	minCandiesForFav := 7

	numDistributions := distributeCandies(candies, grandkids, minCandiesForFav)
	fmt.Println("Cara pak Zaki membagi:", numDistributions)
}

func distributeCandies(candies, grandkids, minCandiesForFav int) int {
	// Base cases:
	if candies == 0 || grandkids == 0 {
		return 0
	}
	if grandkids == 1 {
		return 1
	}

	// Recursive cases:
	count := 0
	for i := minCandiesForFav; i <= candies; i++ {
		// Give the favorite grandchild i candies.
		count += distributeCandies(candies-i, grandkids-1, 0)
	}

	return count
}

func soal3() {
	n := 5
	shuffle := 45
	fmt.Println("Kartu paling atas N=5 dan shuffle=45:", shuffleCards(n, shuffle))
	n = 7
	shuffle = 50
	fmt.Println("Kartu paling atas N=7 dan shuffle=50:", shuffleCards(n, shuffle))

	fmt.Println()
	fmt.Println("Jika N=3 berapa shuffle yang diperlukan agar kartu 2 berada diatas?")
	n = 3
	shuffle = 0
	var cards []int
	for i := 1; i <= 31; i++ {
		cards = append(cards, i)
	}

	for {
		cards = append(cards[len(cards)-n:], cards[:len(cards)-n]...)
		shuffle++
		if cards[0] == 2 {
			break
		}
	}
	fmt.Println(shuffle)
	fmt.Println("Kondisi kartu:", cards)
}

func shuffleCards(n int, shuffle int) int {
	var cards []int
	for i := 1; i <= 31; i++ {
		cards = append(cards, i)
	}

	// Swap the last 5 elements with the first 5 elements using array slicing
	for i := 0; i < shuffle; i++ {
		cards = append(cards[len(cards)-n:], cards[:len(cards)-n]...)
	}

	return cards[0]
}
