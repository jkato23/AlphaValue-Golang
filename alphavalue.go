package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var statement string
	fmt.Println("Hello")
	fmt.Print("Please enter a statement to convert: ")
	reader := bufio.NewReader(os.Stdin)
	statement, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred, please try again.")
		return
	}
	statement = strings.TrimSuffix(statement, "\n")
	for i := 0; i < len(statement); i++ {
		if unicode.IsLetter(rune(statement[i])) == true {
			fmt.Println((statement[i]) - 96)
		} else {
			fmt.Println(string(statement[i]) + " ")
		}
	}
}
