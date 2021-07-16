package fun

import (
	"SocketGo/src/model"
	"SocketGo/src/util"
)

func FuncEcho(info *model.ExecInfo) {
	if len(info.Args) > 1 {
		util.Sayln(info.Cmd[6:])
	}
	info.Set("echo", info.Cmd[6:])
}
