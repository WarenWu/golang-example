package ex

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ExSdtin() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}
	switch input {
	case "waren\r\n": //不加\r\n也可以
		fallthrough
	case "weiming":
		fallthrough
	case "wufufu":
		fmt.Println("Welcome!, ", input)
	default:
		fmt.Println("You are not welcome here, GoodBye!")
	}
}

func ExFileOs() {
	inputFile, err := os.Open("./io/input.txt")
	if err != nil {
		fmt.Println("file opened fail!")
		return
	}
	defer inputFile.Close()
	

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, err := inputReader.ReadString('\n')
		inputRune := []rune(inputString)
		fmt.Println(inputString,len(inputRune))
		if err == io.EOF {
			return
		}
	}
}

func ExFileIoUtil() {
	inputFile := "./io/input.txt"
	outputFile := "./io/output.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("file opened fail!")
		return
	}
	err = ioutil.WriteFile(outputFile,buf,0644)
	if err != nil {
        panic(err.Error())
    }
}
