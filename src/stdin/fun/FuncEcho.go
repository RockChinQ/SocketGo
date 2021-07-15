package fun

import "SocketGo/src/util"

func FuncEcho(args []string, cmd string) map[string]string {
	if len(args) > 1 {
		util.Sayln(cmd[6:])
	}
	m := make(map[string]string)
	m["msg"] = cmd[6:]
	return m
}
