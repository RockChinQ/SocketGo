package stdin

import (
	"SocketGo/src/stdin/fun"
	"SocketGo/src/util"
	"bufio"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var Prompt = "SKG > "

func Loop() {
	for {
		cmd, args, err := FormatInput()
		if err != nil {
			if err.Error() == "EOF" {
				fun.FuncExit([]string{"!exit"}, "!exit")
			} else {
				util.SaySub("Stdin", err.Error())
			}
			continue
		}
		Process(cmd, args)
	}
}

//Get input data splited by ' ' as []string from stdin
func FormatInput() (string, []string, error) {

	input, err := reader.ReadString('\n')

	input = strings.TrimRight(input, "\n")

	formatInputFields := strings.Fields(input)

	return input, formatInputFields, err
}
func PutPrompt() {
	util.Say("\033[42;30m" + Prompt + "\033[0m")
}
