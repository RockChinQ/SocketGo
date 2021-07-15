package util

import (
	"fmt"
	"strconv"
	"time"
)

func Say(msg string) {
	fmt.Print(msg)
}
func Sayln(msg string) {
	Say(msg + "\n")
}
func SaySub(sub string, msg string) {
	Sayln(getNowTimeStr() + "[" + sub + "]" + msg)
}
func Print() {

}

//Get timeStamp string as MM-DD,HH:mm:ss
func getNowTimeStr() string {
	t := time.Now()
	return strconv.Itoa(int(t.Month())) + "-" +
		strconv.Itoa(t.Day()) + "," +
		strconv.Itoa(t.Hour()) + ":" +
		strconv.Itoa(t.Minute()) + ":" +
		strconv.Itoa(t.Second())
}
func GetTimeStr(t time.Time) string {
	return strconv.Itoa(int(t.Month())) + "-" +
		strconv.Itoa(t.Day()) + "," +
		strconv.Itoa(t.Hour()) + ":" +
		strconv.Itoa(t.Minute()) + ":" +
		strconv.Itoa(t.Second())
}
