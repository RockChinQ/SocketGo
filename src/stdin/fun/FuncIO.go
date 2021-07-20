package fun

import (
	"SocketGo/src/conn"
	"SocketGo/src/model"
	"bufio"
	"strconv"
	"strings"
)

//Read from a conn or write to a conn
func FuncIO(info *model.ExecInfo) {
	if len(info.Args) < 3 {
		info.SaySub("FuncIO", "err:Syntax error,type help to get help.")
		info.Error("err:Syntax error,type \"help\" to get help.")
		return
	}
	handler, ok := GetHandler(info.Args[2])
	if !ok {
		info.SaySub("FuncIO", "err:No such conn:"+info.Args[2])
		info.Error("err:No such conn:" + info.Args[2])
		return
	}
	switch info.Args[1] {
	case "r": //read from a conn
		msg := *handler.Buf
		info.Set("result", msg)
		*handler.Buf = ""
		return
	case "w": //write to a conn,message:string as args[3] required.
		if len(info.Args) < 4 {
			info.Set("result", "FAILED")
			info.SaySub("FuncIO", "err:Message(args[3]) is null,type \"help\" to get help.")
			info.Error("err:Message(args[3]) is null,type \"help\" to get help.")
			return
		}
		//Replace escapes
		msg := info.Cmd[(6 + len(info.Args[2])):]
		msg = strings.ReplaceAll(msg, "\\\\", "\\")
		msg = strings.ReplaceAll(msg, "\\n", "\n")
		msg = strings.ReplaceAll(msg, "\\\"", "\"")
		msg = strings.ReplaceAll(msg, "\\t", "\t")
		msg = strings.ReplaceAll(msg, "\\r", "\r")
		//try to write
		err := WriteString(handler, msg)
		if err != nil {
			info.Set("result", "FAILED")
			// info.SaySub("FuncIO", "err:Failed to write msg to conn SID="+strconv.Itoa(handler.SID))
			info.Error("err:Failed to write msg to conn SID=" + strconv.Itoa(handler.SID) + " err=" + err.Error())
			return
		}
		info.Set("result", "SUCCEEDED")
		return
	default:
		info.Set("result", "")
		info.SaySub("FuncIO", "err:No such operation:"+info.Args[1])
		info.Error("err:No such operation:" + info.Args[1])
		return
	}
}

func GetHandler(str string) (*conn.Handler, bool) {
	sid := strings.Split(str, ".")[0]
	conn.Lock.Lock()
	for _, v := range conn.SocketPool {
		if strconv.Itoa(v.SID) == sid {
			conn.Lock.Unlock()
			return v, true
		}
	}
	conn.Lock.Unlock()
	return nil, false
}

func WriteString(h *conn.Handler, msg string) (err error) {
	writer := bufio.NewWriter(h.Conn)
	l, err := writer.Write([]byte(msg))
	writer.Flush()
	if err == nil {
		h.UpD += uint32(l)
	}
	return err
}
