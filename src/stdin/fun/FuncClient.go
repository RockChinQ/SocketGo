package fun

import (
	"SocketGo/src/conn"
	"SocketGo/src/conn/client"
	"SocketGo/src/util"
	"strconv"
)

func FuncClient(args []string, cmd string) map[string]string {
	if len(args) < 2 {
		util.SaySub("Client", "err:Params not enough.")
		return EmptyMap()
	}
	util.DebugMsg("Client", args[1])
	switch args[1] {
	case "conn":
		if len(args) < 3 { //Check if addr is provided
			util.SaySub("Client", "err:Please provide addr:port.")
			return EmptyMap()
		}
		c, err := client.Make(args[2]) //default timeout is 10s
		if err != nil {                //check if there is a err while making conn
			util.SaySub("Client", "err:Making conn:"+err.Error())
			return EmptyMap()
		}
		h := conn.AddHandler(c, "client")
		util.SaySub("Client", "Successfully established:SID="+strconv.Itoa(h.SID)+" info:"+h.Conn.RemoteAddr().String()+" "+util.GetTimeStr(h.ConnT))
	case "timeout": //set timeout
		if len(args) < 3 {
			util.SaySub("Client", "Timeout period="+strconv.Itoa(client.Timeout))
			return EmptyMap()
		}
		v, err := strconv.ParseInt(args[2], 10, strconv.IntSize)
		if err != nil {
			util.SaySub("Client", "err:args[2] is not a number.")
			return EmptyMap()
		}
		client.Timeout = int(v)
		util.SaySub("Client", "Successfully set timeout period="+args[2])
	default: //Provided invalid opertion
		util.SaySub("Client", "err:No such operation:"+args[1])
	}
	return EmptyMap()
}
