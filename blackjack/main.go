package main

import (
	"fmt"
	"./card"
)

type Dealer struct {
	Money int
}

func (d Dealer) CardJudge(p_card string, d_card string) {

}

//func (d Dealer) Split(str string, idf string) []string {
//	for _, char := range str {
//		complex(char, 1)
//	}
//}


func main() {
	deck := card.BuiltDeck()
	card.Shuffle(deck)
	c, deck := card.Draw(deck)
	fmt.Println(c)
}
