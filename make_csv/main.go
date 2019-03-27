package main

import (
	"./io"
)

func main() {
	io.MakeSrcdir2("ddd")
	io.OperateFile("ddd/test.txt", "aaaa")
}