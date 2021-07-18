package stdin

import (
	"SocketGo/src/model"
	"SocketGo/src/stdin/fun"
	"SocketGo/src/util"
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var Prompt = "SKG > "

/*
>Get a command
>Pack strings
>Replace escape characters
>Parse channel relationship
>Launch functions
*/

func Loop() {
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				fun.FuncExit(model.InitExecInfo("!exit", []string{"!exit"}, false))
			} else {
				util.SaySub("Stdin", err.Error())
			}
		}
		if len(input) > 0 {
			processInput(input)
		}
		PutPrompt()
	}
}

//Handle raw input from stdin,process channel relationship,launch each cmd
func processInput(input string) {
	input = strings.TrimRight(input, "\n")
	cmds := strings.Split(input, "|")
	leng := len(cmds)
	dataSet := fun.NoErrMap()
	var err error
	for i := 0; i < leng; i++ {
		if len(strings.TrimSpace(cmds[i])) == 0 {
			continue
		}
		dataSet, err = processCmd(strings.TrimSpace(cmds[i]), dataSet)
		if err != nil {
			util.SaySub("Process", "error occurred while handle:"+cmds[i])
		}
	}
}

//Process single cmd:Replace ds quote in raw cmd,launch processed cmd
func processCmd(cmd string, ds map[string]string) (map[string]string, error) {
	//check if thia is a simplified cmd
	cmd = RealCmd(cmd)

	//check data requirement
	//$fieldName$
	sptRaw := strings.Fields(cmd)
	sptProcessed := make([]string, len(sptRaw))
	reg := regexp.MustCompile(`[$](.*?)[$]`)
	for index, v := range sptRaw { //scan each ele of splited-by-space-arr
		matchs := reg.FindAllStringSubmatch(v, -1)
		tempStr := v //process str of this ele
		//anlz each match
		for _, v1 := range matchs {
			if existKey(ds, v1[1]) {
				tempStr = strings.Replace(tempStr, v1[0], ds[v1[1]], 1)
				cmd = strings.Replace(cmd, v1[0], ds[v1[1]], 1)
			}
		}
		sptProcessed[index] = tempStr
	}
	//Launch func,and cat data from *ExecInfo
	ei := model.InitExecInfo(cmd, sptProcessed, cmd[0] == '@')
	ei.LsData = ds
	Process(ei)
	data := ei.Data

	//check result
	if data["error"] == "NULL" || !existKey(data, "error") {
		return data, nil
	} else {
		return data, errors.New(data["error"])
	}
}

//If this is a simplified form of a specific cmd
//change it to the real cmd.
//If not,return original cmd.
func RealCmd(raw string) string {
	stWithAt := raw[0] == '@'
	noat := strings.TrimLeft(raw, "@")

	if noat[0] == '"' { //echo
		if stWithAt {
			return "@echo " + noat[1:]
		} else {
			return "echo " + noat[1:]
		}
	} else {
		return raw
	}
}

func PutPrompt() {
	util.Say("\033[42;30m" + Prompt + "\033[0m")
}
func existKey(m map[string]string, key string) bool {
	_, ok := m[key]
	return ok
}
