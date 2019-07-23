package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"reflect"
	"strconv"
)

func oneLineInt(s string) []int {
	temp := strings.Split(s, " ")
	fmt.Println(len(temp))
	ints := []int{}
	for _, elem := range(temp) {
		int_elem, _ := strconv.Atoi(elem)
		ints = append(ints, int_elem)
	}
	return ints
}

func coinCount(ints []int) int {
	cons := ints[0]
	bonus := ints[1]
	target := ints[2]
	count := 0
	day := 0
	var dc int
	for day_count, _ := range(make([]int, target)) {
		day++
		if day % 7 == 0 {
			count += bonus
		}
		count += cons
		if count >= target {
			dc = day_count + 1
			break
		}
	}
	return dc
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var txt string
	if scanner.Scan() {
		txt = scanner.Text()
	}
	ints := oneLineInt(txt)
	fmt.Println(ints[1])
	fmt.Println(reflect.TypeOf(ints[0]))
	result := coinCount(ints)
	fmt.Println(result)
}