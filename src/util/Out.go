package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nsf/termbox-go"
)

func InitTermbox() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	termbox.SetCursor(0, 0)
}

func Say(msg string) {
	fmt.Print(msg)
}
func Sayln(msg string) {
	Say(msg + "\n")
}
func SaySub(sub string, msg string) {
	Sayln(getTimeStr() + "[" + sub + "]" + msg)
}
func Print() {

}

//Get timeStamp string as MM-DD,HH:mm:ss
func getTimeStr() string {
	t := time.Now()
	return strconv.Itoa(int(t.Month())) + "-" + strconv.Itoa(t.Day()) + "," + strconv.Itoa(t.Hour()) + ":" + strconv.Itoa(t.Minute()) + ":" + strconv.Itoa(t.Second())
}
