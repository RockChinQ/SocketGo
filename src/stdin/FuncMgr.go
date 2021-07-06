package stdin

import (
	"SocketGo/src/stdin/fun"
	"SocketGo/src/util"
)

type Run func(args []string, cmd string) map[string]string

var funcList map[string]Run

//Process a provided command
func Process(cmd string, args []string) {
	if len(args) == 0 { //empty cmd
		PutPrompt()
		return
	}
	found := false
	for k, v := range funcList {
		if k == args[0] {
			found = true
			//TODO add support for channel
			_ = v(args, cmd)
			break
		}
	}
	if !found {
		util.SaySub("FuncMgr", "err:No such function:"+args[0])
	}
	PutPrompt()
}

//Register all supported cmd
func RegisterFuns() {
	util.DebugMsg("FuncMgr", "Initially registering all functions.")
	funcList = make(map[string]Run)

	funcList["!help"] = fun.FuncHelp
	funcList["!exit"] = fun.FuncExit
	funcList["!list"] = fun.FuncList
	funcList["!client"] = fun.FuncClient
	funcList["!kill"] = fun.FuncKill
	funcList["!server"] = fun.FuncServer
}
