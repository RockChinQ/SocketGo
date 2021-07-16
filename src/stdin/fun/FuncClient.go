package fun

import (
	"SocketGo/src/conn"
	"SocketGo/src/conn/client"
	"SocketGo/src/model"
	"SocketGo/src/util"
	"strconv"
)

func FuncClient(info *model.ExecInfo) {
	if len(info.Args) < 2 {
		info.SaySub("FuncClient", "err:Params not enough.")
		info.Error("err:Params not enough.")
		return
	}
	switch info.Args[1] {
	case "conn":
		if len(info.Args) < 3 { //Check if addr is provided
			info.SaySub("FuncClient", "err:Please provide addr:port.")
			info.Error("err:Please provide addr:port.")
			return
		}
		util.DebugMsg("FuncClient", info.Args[2])
		c, err := client.Make(info.Args[2]) //default timeout is 10s
		if err != nil {                     //check if there is a err while making conn
			info.SaySub("FuncClient", "err:Making conn:"+err.Error())
			info.Error("err:Making conn:" + err.Error())
			return
		}
		h := conn.AddHandler(c, "client")
		info.SaySub("FuncClient", "Successfully established:SID="+
			strconv.Itoa(h.SID)+" info:"+
			h.Conn.RemoteAddr().String()+" "+
			util.GetTimeStr(h.ConnT))
	case "timeout": //set timeout
		if len(info.Args) < 3 {
			info.SaySub("FuncClient", "Timeout period="+strconv.Itoa(client.Timeout))
			info.Set("timeout", strconv.Itoa(client.Timeout))
			return
		}
		v, err := strconv.ParseInt(info.Args[2], 10, strconv.IntSize)
		if err != nil {
			info.SaySub("FuncClient", "err:args[2] is not a number.")
			info.Error("err:args[2] is not a number.")
			return
		}
		client.Timeout = int(v)
		info.SaySub("FuncClient", "Successfully set timeout period="+info.Args[2])
	default: //Provided invalid opertion
		info.SaySub("FuncClient", "err:No such operation:"+info.Args[1])
		info.Error("err:No such operation:" + info.Args[1])
		return
	}
}
