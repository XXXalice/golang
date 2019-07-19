package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	read_judge := sc.Scan()
	var txt string
	if read_judge {
		txt = sc.Text()
	}
	fmt.Println(txt)
}

func lineParse(sc bufio.Scanner) []string {
	line := []string{}
	read_judge := sc.Scan()
	if read_judge {
		txt := sc.Text()
		line = append(line, txt)
	}
	return line
}

func janken(cpu_hand, you_hand int) string {
	return "test"
}