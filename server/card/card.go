package card

import "math/rand"

const HAND_SIZE = 8

type Card struct {
	Suit  string
	Value string
}

var Suits = []string{
	"Club", "Diamond", "Spade", "Heart",
}

var Values = []string{
	"Ace", "King", "Queen", "10", "9",
}

func DealHand() []Card {
	cards := []Card{}
	for _, suit := range Suits {
		for _, val := range Values {
			cards = append(cards, Card{suit, val})
		}
	}
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return cards[:HAND_SIZE]
}
