package io

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"log"
)

const FIELD = "../"
var fileOpenModes = []int{os.O_WRONLY, os.O_RDWR, os.O_CREATE, os.O_APPEND}

func MakeSrcdir(folName string){
	if err := os.Mkdir(folName, 0777); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func MakeSrcdir2(folName string) {
	if _, err := os.Stat(folName);
		os.IsNotExist(err) {
		os.Mkdir(folName, 0777)
	}
}

func OperateFile(fileName, inputString string){
	file, err := os.OpenFile(fileName, os.O_RDWR | os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer file.Close()
}

func StrStdin() (inputString string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString = strings.TrimSpace(scanner.Text())
	return
}