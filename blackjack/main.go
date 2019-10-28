package main

import (
	"./card"
	"strings"
	"strconv"
)

type BlackjackDealer struct {
	Money int
	Dealer_sum int
	Player_sum int
	Dealer_hands []string
	Player_hands []string
	Deck []string
}

func (d BlackjackDealer) Init() {
	d.Deck = card.BuiltDeck()
	card.Shuffle(d.Deck)
	var c string
	for range(make([]int, 2)) {
		c, d.Deck = card.Draw(d.Deck)
		d.Dealer_hands = append(d.Dealer_hands, c)
		c, d.Deck = card.Draw(d.Deck)
		d.Player_hands = append(d.Player_hands, c)
	}
}

func (d BlackjackDealer) CardJudge(p_card string, d_card string) bool {
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

func (d BlackjackDealer) GameProgress(choice string) {
	if choice == "hit" {

	} else if choice == "stand" {

	}
}

//func (d Dealer) Split(str string, idf string) []string {
//	for _, char := range str {
//		complex(char, 1)
//	}
//}


func main() {
	dealer := BlackjackDealer{
		Money: 100,
		Dealer_sum: 0,
		Player_sum: 0,
		Dealer_hands: make([]string, 0),
		Player_hands: make([]string, 0),
		Deck: make([]string, 0),
	}
	dealer.Init()
}
