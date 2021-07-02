package main

import "SocketGo/src/fun"

//Process a provided command
func Process(cmd string) {

}

//Register all supported cmd
func RegisterFuns() {
	FuncList["!help"] = fun.FuncHelp
}
