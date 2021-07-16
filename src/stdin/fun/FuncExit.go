package fun

import (
	"SocketGo/src/model"
	"SocketGo/src/util"
	"os"
)

//Dispose resources and exit:all sockets,opend ports
func FuncExit(info *model.ExecInfo) {
	DisposeAll()
	util.SaySub("FuncExit", "Exiting.")
	os.Exit(0)
}

func DisposeAll() {
	//kill all conns
	FuncKill(model.InitExecInfo("!kill all", []string{"!kill", "all"}, false))
	//close all ports
	FuncServer(model.InitExecInfo("!server close all", []string{"!server", "close", "all"}, false))
}
