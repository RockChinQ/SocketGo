package util

import (
	"fmt"
	"strconv"
	"time"
)

func Say(msg string) {
	fmt.Print(getTimeStr() + msg)
}
func Sayln(msg string) {
	Say(msg + "\n")
}
func SaySub(sub string, msg string) {
	Sayln("[" + sub + "]" + msg)
}

//Get timeStamp string as MM-DD,HH:mm:ss
func getTimeStr() string {
	t := time.Now()
	return strconv.Itoa(int(t.Month())) + "-" + strconv.Itoa(t.Day()) + "," + strconv.Itoa(t.Hour()) + ":" + strconv.Itoa(t.Minute()) + ":" + strconv.Itoa(t.Second())
}
