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
	fmt.Println("Iniciando")
	menu.ShowIntro()
	for {
		menu.ShowMenu()
		command := menu.ReadCommand()

		switch command {
		case 1:
			name := watch.RegisterName()
			dateAndTime := watch.RegisterTime()
			watch.RegisterLog(name, dateAndTime)
			osactions.ClearScreen()
			fmt.Println("Presença registrada com sucesso")
			time.Sleep(2 * time.Second)
		case 2:
			osactions.ClearScreen()
			fmt.Println("Função em desenvolvimento")
			time.Sleep(5 * time.Second)
			// exibeLog()
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
