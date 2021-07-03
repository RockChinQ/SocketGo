package stdin

import (
	"SocketGo/src/stdin/fun"
	"SocketGo/src/util"
)

type Run func(args []string, cmd string) map[string]string

var FuncList map[string]Run

//Process a provided command
func Process(cmd string) {

}

//Register all supported cmd
func RegisterFuns() {
	util.DebugMsg("FuncMgr", "Initially registering all functions.")
	FuncList = make(map[string]Run)
	FuncList["!help"] = fun.FuncHelp
}
