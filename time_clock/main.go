package main

import (
	"fmt"
	"os"
	"time_clock/time_clock/menu"
)

func main() {
	fmt.Println("Iniciando")
	menu.ShowIntro()
	for {
		menu.ShowMenu()
		command := menu.ReadCommand()
		fmt.Println(command)

		switch command {
		case 1:
			// iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			// exibeLog()
		case 0:
			fmt.Println("Até logo...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

	}
}
