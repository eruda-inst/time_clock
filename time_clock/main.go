package main

import (
	"fmt"
	"os"
	"time"
	"time_clock/time_clock/menu"
	"time_clock/time_clock/osactions"
	"time_clock/time_clock/watch"
)

func main() {
	filename := "log/log.csv"
	// fmt.Println("Iniciando")
	// menu.ShowIntro()
	for {
		menu.ShowMenu()
		command := menu.ReadCommand()

		switch command {
		case 1:
			name := watch.RegisterName()
			dateAndTime := watch.RegisterTime()
			if osactions.FileExists(filename) {
				file, err := os.OpenFile("log/log.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
				if err != nil {
					fmt.Println("An error has ocurred", err)
				}
				file.WriteString("name,date/time\n")
				fmt.Println("Arquivo não existe")
			}
			watch.RegisterLog(name, dateAndTime)
			osactions.ClearScreen()
			fmt.Println("Presença registrada com sucesso")
			time.Sleep(2 * time.Second)

		case 0:
			osactions.ClearScreen()
			fmt.Println("Até logo...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

	}
}
