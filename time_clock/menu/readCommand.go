package menu

import "fmt"

func ReadCommand() int {
	var readCommand int
	_, err := fmt.Scan(&readCommand)
	if err != nil {
		fmt.Println("Erro ao ler o comando:", err)
		return -1
	}
	return readCommand
}
