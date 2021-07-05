package fun

import (
	"SocketGo/src/conn"
	"SocketGo/src/util"
	"strconv"
)

func FuncList(args []string, cmd string) map[string]string {
	if len(args) == 1 { //all connection
		util.SaySub("FuncList", "SID\tcreator\tconnTime\tspeed(up/down)\tdata(up/down)")
		for k, v := range conn.SocketPool {
			util.SaySub("FuncList", itoa(k)+"\t"+v.As+"\t"+util.GetTimeStr(v.ConnT)+"\t"+itoa(v.UpV)+"/"+itoa(v.DownV)+"\t"+itoa32(v.UpD)+"/"+itoa32(v.DownD))
		}
	}
	return EmptyMap()
}

func itoa(i int) string {
	return strconv.Itoa(i)
}
func itoa32(i uint32) string {
	return strconv.FormatUint(uint64(i), 10)
}

func EmptyMap() map[string]string {
	return make(map[string]string)
}
