package watch

import "time"

func RegisterTime() string {
	CurrentDateTime := time.Now().Format("01/02/2006 15:04:05")
	return CurrentDateTime
}
