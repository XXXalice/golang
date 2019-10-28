package main

import (
	"fmt"
	"./card"
	"strings"
	"strconv"
)

type BlackjackDealer struct {
	Money int
}

func (d Dealer) CardJudge(p_card string, d_card string) bool {
	p_num, _ := strconv.Atoi(strings.Split(p_card, "_")[1])
	d_num, _ := strconv.Atoi(strings.Split(d_card, "_")[1])
	var res bool
	if p_num > d_num {
		res = true
	} else {
		res = false
	}
	return res
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
