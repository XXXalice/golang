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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var txt string
	if scanner.Scan() {
		txt = scanner.Text()
	}
	ints := oneLineInt(txt)
	fmt.Println(ints[1])
	fmt.Println(reflect.TypeOf(ints[0]))
}