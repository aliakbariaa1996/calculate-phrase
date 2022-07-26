package main

import (
	"bufio"
	"fmt"
	"go/token"
	"go/types"
	"os"
	"strconv"
	"strings"
)

var (
	counter    = 0
	userPhrase = ""
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the specify lines of calculation: ")
	numStr, err := reader.ReadString('\n')
	numStr = strings.Replace(numStr, "\n", "", -1)
	if err != nil {
		panic(err)
	}
	strings.ReplaceAll(numStr, " ", "")
	num, err := convertStringToInt(numStr)
	if err != nil {
		panic(err)
	}
	getUserInput(num)
	fs := token.NewFileSet()
	tv, err := types.Eval(fs, nil, token.NoPos, userPhrase)
	if err != nil {
		panic(err)
	}
	println("\nExpected result: ", tv.Value.String())
}

// get phrase from user
func getUserInput(num int) {
	if num != counter {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Phrase Number " + string(rune(counter+1)) + ":")
		numStr, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		numStr = strings.ReplaceAll(numStr, "\n", "")
		numStr = strings.ReplaceAll(numStr, ",", "")
		numStr = strings.ReplaceAll(numStr, " ", "")
		if counter > 0 {
			userPhrase = userPhrase + "+" + numStr
		} else {
			userPhrase = numStr
		}
		counter++
		getUserInput(num)
	}
}

// convert string to number
func convertStringToInt(str string) (int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return num, nil
}
