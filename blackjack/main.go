package main

import (
	"fmt"
	"math/rand"
)

func BuiltDeck() []string {
	number := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	mark := []string{"spade","heart","clover","diamond"}
	deck := make([]string, 0)
	for _, n := range number {
		for _, m := range mark {
			card := m + "_" + n
			deck = append(deck, card)
		}
	}
	return deck
}

func Shuffle(d []string) {
	n := len(d)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}

func Draw(d []string) (string, []string) {
	drawCard := d[0]
	d = append(d[:0], d[1:]...)
	return drawCard, d
}

func main() {
	deck := BuiltDeck()
	Shuffle(deck)
	c, deck := Draw(deck)
	fmt.Println(c)
}
