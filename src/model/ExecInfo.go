package model

import (
	"SocketGo/src/util"
	"strconv"
	"time"
)

/*
Pass a ExecInfo instance pointer to a Func and return with the info of the Func
*/
type ExecInfo struct {
	Cmd  string
	Args []string

	Data map[string]string

	Mute bool
}

func InitExecInfo(cmd string, args []string, mute bool) *ExecInfo {
	var ei = &ExecInfo{
		Cmd:  cmd,
		Mute: mute,
	}
	ei.Args = args
	util.DebugMsg("InitExecInfo", "lenof args:"+strconv.Itoa(len(args)))
	ei.Data = make(map[string]string)
	ei.Data["error"] = "NULL"
	ei.Data["output"] = ""
	return ei
}

func (e *ExecInfo) Set(key string, value string) {
	e.Data[key] = value
}
func (e *ExecInfo) Append(key string, value string) {
	pV, ok := e.Data[key]
	if !ok {
		e.Data[key] = value
	} else {
		e.Data[key] = pV + value
	}
}
func (e *ExecInfo) Error(msg string) {
	e.Data["error"] = msg
}
func (e *ExecInfo) Say(msg string) {
	e.Append("output", msg)
	if !e.Mute {
		util.Say(msg)
	}
}
func (e *ExecInfo) Sayln(msg string) {
	e.Say(msg + "\n")
}
func (e *ExecInfo) SaySub(sub string, msg string) {
	e.Sayln(util.GetTimeStr(time.Now()) + "[" + sub + "]" + msg)
}
