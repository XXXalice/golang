package main

import (
	"fmt"
	"./card"
)


func main() {
	deck := card.BuiltDeck()
	card.Shuffle(deck)
	c, deck := card.Draw(deck)
	fmt.Println(c)
}
