package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ranks = [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
var suits = [4]string{"Spades", "Clubs", "Diamonds", "Hearts"}

func generateDeck() []string{
	deck := []string{}
	
	for i := range ranks{
		for j := range suits {
			deck = append(deck, ranks[i] + " of " + suits[j])
		}
	}

	return deck
}

func dealOneCard(deck []string){
	random := rand.Intn(len(deck))
	chosenCard := deck[random]
	fmt.Println("Picked the", chosenCard)
}

func shuffle(deck []string) []string{
	newOrder := rand.Perm(len(deck))

	shuffledDeck := []string{}
	for _, i := range newOrder{
		shuffledDeck = append(shuffledDeck, deck[i])
	}

	return shuffledDeck
}

func main() {
	rand.Seed(time.Now().UnixNano())

	deck := generateDeck()
	dealOneCard(deck)

	fmt.Println("\nTop 5 cards before shuffle")
	for i := 0; i < 5; i++{
		fmt.Println(deck[i])
	}

	fmt.Println()
	shuffledDeck := shuffle(deck)

	fmt.Println("Top 5 cards after shuffle")
	for i := 0; i < 5; i++{
		fmt.Println(shuffledDeck[i])
	}
}
