package stdin

import (
	"SocketGo/src/util"
	"os"

	"github.com/nsf/termbox-go"
)

var prompt = "> "
var x = 0
var buffer = ""
var spcBuffer string

func Loop() {
	ReputPrompt()
	for {
		event := termbox.PollEvent()
		switch int(event.Key) {
		case 0: //char input
			util.Say(string(event.Ch))
			buffer += string(event.Ch)
			spcBuffer += string(event.Ch)
			x++
		case 32: //space
			util.Say(" ")
			buffer += " "
			spcBuffer = ""
			x++
		case 3:
			os.Exit(0)
		case 8: //backspace
			if len(buffer) >= 1 {
				buffer = buffer[:len(buffer)-1]
				x--
				// util.DebugMsg("stdin", "x="+strconv.Itoa(x)+" h="+strconv.Itoa(GetTerminalHeight()))
				termbox.SetCursor(x, GetTerminalHeight()-1)
				util.Say(" ")
				termbox.SetCursor(x, GetTerminalHeight()-1)
			}
		case 13: //enter
			Process(buffer)
			//cmd finished
			ReputPrompt()
		}
	}
}

func ReputPrompt() {
	_, h := termbox.Size()
	termbox.SetCursor(0, h-1)
	for i := 0; i < GetTerminalWidth()-1; i++ {
		util.Say(" ")
	}
	termbox.SetCursor(0, h-1)
	x = 0
	util.Say(prompt)
	x += len(prompt)
	termbox.SetCursor(x, h-1)
	buffer = ""
	spcBuffer = ""
}
func GetTerminalHeight() int {
	_, h := termbox.Size()
	return h
}
func GetTerminalWidth() int {
	w, _ := termbox.Size()
	return w
}
