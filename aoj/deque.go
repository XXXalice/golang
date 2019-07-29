package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"errors"
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

func executeOrder(order []int, arr []int) (out int, err error) {
	// 破壊: arr
	output := -1
	switch order[0] {
	case 0:
		// push
		if order[1] == 0 {
			arr, arr[0] = append(arr[:1], arr[0:]...), order[2]
		}else if order[1] == 1{
			arr = append(arr, order[2])
		}
	case 1:
		// randomAccess
		output =  arr[order[1]]
	case 2:
		// pop
		var pop_pos int
		if order[1] == 0 {
			pop_pos = 0
		}else if order[1] == 1 {
			pop_pos = -1
		}
		arr = append(arr[:pop_pos], arr[pop_pos+1:]...)
	}
	if output >= 0 {
		return output, nil
	}else {
		return -1, errors.New("not randomAccess.")
	}
}

func main() {
	var q_num int
	fmt.Scan(&q_num)
	orders := [][]int{}
	for range(make([]int, q_num)) {
		one_order := nLine()
		orders = append(orders, strToIntArray(one_order))
	}
	arr := []int{}
	result := []int{}
	for _, op := range(orders) {
		out, err := executeOrder(op, arr)
		if err == nil {
			result = append(result, out)
		}
	}
	for _, num := range(result) {
		fmt.Println(num)
	}
}