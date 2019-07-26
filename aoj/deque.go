package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nLine() string {
	sc.Scan()
	return sc.Text()
}

func strToIntArray(s string) []int {
	str_arr := strings.Split(s, " ")
	int_arr := []int{}
	for _, val := range str_arr {
		int_val, err := strconv.Atoi(val)
		if err == nil {
			int_arr = append(int_arr, int_val)
		}
	}
	return int_arr
}

func main() {
	var q_num int
	fmt.Scan(&q_num)
	order := [][]int{}
	for range(make([]int, q_num)) {
		one_order := nLine()
		order = append(order, strToIntArray(one_order))
	}
	fmt.Println(order)
}