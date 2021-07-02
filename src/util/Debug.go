package util

func DebugMsg(sub string, msg string) {
	if DebugMode {
		SaySub("Debug-"+sub, msg)
	}
}
