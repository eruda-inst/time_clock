package menu

import (
	"fmt"
	"time_clock/time_clock/osactions"
)

func ShowIntro() {
	var nome string
	versao := "0.1"
	fmt.Println("Olá, qual seu nome?")
	fmt.Scan(&nome)
	osactions.ClearScreen()
	println("Olá", nome)
	println("Este programa está na versão", versao)
}
