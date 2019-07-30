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
var arr = []int{}

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

func executeOrder(order []int) (out int, err error) {
	// 破壊: arr
	var output int
	var output_judge bool
	output_judge = false
	switch order[0] {
	case 0:
		// push
		if order[1] == 0 || len(arr) == 0{
			arr = append([]int{order[2]}, arr[0:]...)
		}else {
			arr = append(arr, order[2])
		}
	case 1:
		// randomAccess
		output_judge = true
		output =  arr[order[1]]
	case 2:
		// pop
		if order[1] == 0 {
			// 前
			arr = append(arr[:0], arr[0+1:]...)
		}else if order[1] == 1 {
			// 末尾
			arr = arr[:len(arr)-1]
		}

	}

	if output_judge {
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
	result := []int{}
	for _, op := range(orders) {
		out, err := executeOrder(op)

		fmt.Println(arr)

		if err == nil {
			result = append(result, out)
		}
	}
	for _, num := range(result) {
		fmt.Println(num)
	}
}