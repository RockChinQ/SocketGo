package fun

import (
	"SocketGo/src/model"
	"strings"
)

func FuncEcho(info *model.ExecInfo) {
	msg := info.Cmd[5:]
	msg = strings.ReplaceAll(msg, "\\\\", "\\")
	msg = strings.ReplaceAll(msg, "\\n", "\n")
	msg = strings.ReplaceAll(msg, "\\\"", "\"")
	msg = strings.ReplaceAll(msg, "\\t", "\t")
	msg = strings.ReplaceAll(msg, "\\r", "\r")
	info.Sayln(msg)
	info.Set("echo", msg)
}
