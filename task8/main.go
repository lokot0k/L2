package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

var ntpServer = "time.google.com"

func main() {

	currTime, err := ntp.Time(ntpServer)
	if err != nil {
		_, err = os.Stderr.WriteString(err.Error() + "\n")
		if err != nil {
			panic("Couldn't write error to STDERR")
		}
	}

	fmt.Printf("Текущее время %s: %s\n", ntpServer, currTime.Format(time.RFC822Z))
}
