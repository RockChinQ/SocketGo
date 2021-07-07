package fun

import (
	"SocketGo/src/conn"
	"SocketGo/src/util"
	"strconv"
)

func FuncList(args []string, cmd string) map[string]string {
	if len(args) == 1 { //all connection
		conn.Lock.Lock()
		util.SaySub("FuncList", "SID\tcreator\tconnTime\tspeed(up/down)\tdata(up/down)\tstatus")
		for k, v := range conn.SocketPool {
			util.SaySub("FuncList", itoa(k)+"\t"+v.As+"\t"+util.GetTimeStr(v.ConnT)+"\t"+itoa(v.UpV)+"/"+itoa(v.DownV)+"\t"+itoa32(v.UpD)+"/"+itoa32(v.DownD)+"\t"+v.Status)
		}
		util.SaySub("FuncList", "Done,count:"+strconv.Itoa(len(conn.SocketPool)))
		conn.Lock.Unlock()
	} else {
		if args[1] == "client" || args[1] == "server" {
			conn.Lock.Lock()
			count := 0
			util.SaySub("FuncList", "SID\tcreator\tconnTime\tspeed(up/down)\tdata(up/down)\tstatus")
			for k, v := range conn.SocketPool {
				if v.As == args[1] {
					count++
					util.SaySub("FuncList", itoa(k)+"\t"+v.As+"\t"+util.GetTimeStr(v.ConnT)+"\t"+itoa(v.UpV)+"/"+itoa(v.DownV)+"\t"+itoa32(v.UpD)+"/"+itoa32(v.DownD)+"\t"+v.Status)
				}
			}
			util.SaySub("FuncList", "Done,count:"+strconv.Itoa(count))
			conn.Lock.Unlock()
		} else {
			util.SaySub("FuncList", "err:args[1] should be \"client\" or \"server\".")
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
