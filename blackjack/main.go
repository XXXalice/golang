package main

import (
	"fmt"
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

func main() {
	//fmt.Println("who are you?")
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//fmt.Println("Hi,", scanner.Text())
	deck := BuiltDeck()
	fmt.Println(deck)
}
