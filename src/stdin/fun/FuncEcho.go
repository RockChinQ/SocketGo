package fun

import (
	"SocketGo/src/model"
)

func FuncEcho(info *model.ExecInfo) {
	if len(info.Args) > 1 {
		info.Sayln(info.Cmd[5:])
	}
	info.Set("echo", info.Cmd[5:])
}
