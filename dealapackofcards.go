package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	for player := 0; player < 4; player++ {
		start := player * 3
		fmt.Printf("Player %d: %d, %d, %d\n", player+1, deck[start], deck[start+1], deck[start+2])
	}
}
