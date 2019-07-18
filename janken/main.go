package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	txt := lineParse(sc)
	fmt.Println(txt)
}

func lineParse(sc bufio.Scanner) []string {
	line := []string
	read_judge := sc.Scan()
	if read_judge {
		txt := sc.Text()
		for _, s := range txt{
			line = append(line, s)
		}
	}
	return txt
}

func janken(cpu_hand, you_hand int) string {
	return "test"
}