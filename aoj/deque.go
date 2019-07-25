package main

import (
	"fmt"
	"os"
	"bufio"
	//"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	var q_num int
	fmt.Scan(&q_num)
	order := []string{}
	for range(make([]int, q_num)) {
		order = append(order, nLine())
	}

}