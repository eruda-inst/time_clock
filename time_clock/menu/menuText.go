package menu

import (
	"fmt"
	"time_clock/time_clock/osactions"
)

func ShowMenu() {
	version := "1.0"
	osactions.ClearScreen()
	fmt.Println("==============================")
	fmt.Println("REGISTRO DE PRESENÇA -", version)
	fmt.Println("==============================\n")
	fmt.Println("1 - Registrar um ponto")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}
