package fun

import (
	"SocketGo/src/conn"
	"SocketGo/src/conn/client"
	"SocketGo/src/util"
	"strconv"
)

func FuncClient(args []string, cmd string) map[string]string {
	if len(args) < 2 {
		util.SaySub("FuncClient", "err:Params not enough.")
		return EmptyMap()
	}
	switch args[1] {
	case "conn":
		if len(args) < 3 { //Check if addr is provided
			util.SaySub("FuncClient", "err:Please provide addr:port.")
			return EmptyMap()
		}
		util.DebugMsg("FuncClient", args[2])
		c, err := client.Make(args[2]) //default timeout is 10s
		if err != nil {                //check if there is a err while making conn
			util.SaySub("FuncClient", "err:Making conn:"+err.Error())
			return EmptyMap()
		}
		h := conn.AddHandler(c, "client")
		util.SaySub("FuncClient", "Successfully established:SID="+strconv.Itoa(h.SID)+" info:"+h.Conn.RemoteAddr().String()+" "+util.GetTimeStr(h.ConnT))
	case "timeout": //set timeout
		if len(args) < 3 {
			util.SaySub("FuncClient", "Timeout period="+strconv.Itoa(client.Timeout))
			return EmptyMap()
		}
		v, err := strconv.ParseInt(args[2], 10, strconv.IntSize)
		if err != nil {
			util.SaySub("FuncClient", "err:args[2] is not a number.")
			return EmptyMap()
		}
		client.Timeout = int(v)
		util.SaySub("FuncClient", "Successfully set timeout period="+args[2])
	default: //Provided invalid opertion
		util.SaySub("FuncClient", "err:No such operation:"+args[1])
	}
	return EmptyMap()
}