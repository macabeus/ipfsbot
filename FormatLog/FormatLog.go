package FormatLog

import (
	"fmt"
	"time"
)

func Print(prefix string, message ...interface{}) {
	t := time.Now()
	textTime := t.Format(time.Stamp)

	fmt.Print("[", textTime, "] [", prefix, "] ")
	fmt.Println(message)
}
