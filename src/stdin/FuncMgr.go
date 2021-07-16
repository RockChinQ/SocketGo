package stdin

import (
	"SocketGo/src/model"
	"SocketGo/src/stdin/fun"
	"SocketGo/src/util"
	"strings"
)

type Run func(info *model.ExecInfo)

var funcList map[string]Run

//Process a provided command
func Process(info *model.ExecInfo) {
	if len(info.Args) == 0 { //empty cmd
		return
	}
	found := false
	for k, v := range funcList {
		if k == strings.TrimLeft(info.Args[0], "@") {
			found = true
			info.Cmd = strings.TrimLeft(info.Cmd, "@")
			info.Args[0] = strings.TrimLeft(info.Args[0], "@")
			v(info)
			info.Data["output"] = strings.TrimRight(info.Data["output"], "\n")
			return
		}
	}
	if !found {
		util.SaySub("FuncMgr", "err:No such function:"+info.Args[0])
		info.Error("err:No such function:" + info.Args[0])
	}
}

//Register all supported cmd
func RegisterFuns() {
	util.DebugMsg("FuncMgr", "Initially registering all functions.")
	funcList = make(map[string]Run)

	funcList["help"] = fun.FuncHelp
	funcList["exit"] = fun.FuncExit
	funcList["list"] = fun.FuncList
	funcList["client"] = fun.FuncClient
	funcList["kill"] = fun.FuncKill
	funcList["server"] = fun.FuncServer
	funcList["echo"] = fun.FuncEcho
}
