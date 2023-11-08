package menu

import (
	"fmt"
	"time_clock/time_clock/osactions"
)

func ShowMenu() {
	osactions.ClearScreen()
	fmt.Println("1 - Registrar um ponto")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}
