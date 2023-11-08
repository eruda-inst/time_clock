package watch

import (
	"fmt"
	"os"
)

func RegisterLog(name string, dateAndTime string) {

	file, err := os.OpenFile("log/log.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("An error has ocurred", err)
	}

	file.WriteString(name + "," + dateAndTime + "\n")

	file.Close()
}
