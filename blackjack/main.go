package main

import (
	"fmt"
	"./card"
	"strings"
	"strconv"
	"os"
)

type BlackjackDealer struct {
	Money int
	Dealer_sum int
	Player_sum int
	Dealer_hands []string
	Player_hands []string
	Deck []string
}

func (d *BlackjackDealer) Init() {
	d.Deck = card.Shuffle(card.BuiltDeck())
	var c string
	for range make([]int, 2) {
		c, d.Deck = card.Draw(d.Deck)
		d.Dealer_hands = append(d.Dealer_hands, c)
		c, d.Deck = card.Draw(d.Deck)
		d.Player_hands = append(d.Player_hands, c)
	}
}

func (d *BlackjackDealer) CardJudge(p_card string, d_card string) bool {
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

func (d *BlackjackDealer) GameProgress(choice string) bool {
	var c string
	var judge bool
	if choice == "hit" {
		d.Player_sum = 0
		c, d.Deck = card.Draw(d.Deck)
		d.Player_hands = append(d.Player_hands, c)
		for _, p_card := range d.Player_hands {
			num, _ := strconv.Atoi(strings.Split(p_card, "_")[1])
			d.Player_sum += num
		}
		if d.Player_sum > 21 {
			fmt.Println(d.Player_hands)
			fmt.Println("you bust!")
			os.Exit(0)
		}
		judge = false
	} else if choice == "stand" {
		d.Dealer_sum = 0
		for _, d_card := range d.Dealer_hands {
			num, _ := strconv.Atoi(strings.Split(d_card, "_")[1])
			d.Dealer_sum += num
		}
		judge = true
	}
	return judge
}

func (d *BlackjackDealer) Judge() bool {
	var win bool
	if d.Player_sum > d.Dealer_sum {
		win = true
	} else {
		win = false
	}
	fmt.Println("your sum:", d.Player_sum)
	fmt.Println("dealer sum:", d.Dealer_sum)
	return win
}


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
	var command string
	var win bool
	for {
		fmt.Println("---your hands---")
		fmt.Println(dealer.Player_hands)
		fmt.Println("---dealer hands---")
		var hidden []string
		for idx, c := range dealer.Dealer_hands {
			if idx == 0 {
				hidden = append(hidden, c)
			} else {
				hidden = append(hidden, "???")
			}
		}
		fmt.Println(hidden)
		fmt.Print("Plz input command hit or stand. :")
		fmt.Scan(&command)
		if dealer.GameProgress(command) {
			win = dealer.Judge()
			if win {
				fmt.Println("you win!")
			} else {
				fmt.Println("you lose!")
			}
			os.Exit(0)
		}
	}

}
