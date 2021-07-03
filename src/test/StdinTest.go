package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader

func StdinTest() {
	//get user's input
	reader = bufio.NewReader(os.Stdin)

	for {
		sA, err := formatInput()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			for i := 0; i < len(sA); i++ {
				fmt.Print(sA[i] + " ")
			}
			fmt.Println()
		}

	}
}

func formatInput() ([]string, error) {

	input, err := reader.ReadString('\n')

	input = strings.TrimRight(input, "\n")

	formatInputFields := strings.Fields(input)

	return formatInputFields, err
}
