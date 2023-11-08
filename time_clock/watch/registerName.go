package watch

import (
	"fmt"
	"os"
	"time_clock/time_clock/osactions"
)

func RegisterName() string {
	osactions.ClearScreen()
	fmt.Println("Indique o nome para registro: ")
	var NameRagistration string
	_, err := fmt.Scan(&NameRagistration)
	if err != nil {
		fmt.Println("Erro ao ler o nome:", err)
		os.Exit(1)
	}
	return NameRagistration
}
