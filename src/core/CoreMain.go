package main

import "SocketGo/src/util"

type Run func(args []string, cmd string)

var FuncList map[string]Run

func main() {
	util.DebugMsg("Main", "Booting.")
	FuncList = make(map[string]Run)
	RegisterFuns()
}
