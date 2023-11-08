package menu

import "fmt"

func ReadCommand() int {
	var ReadCommand int
	_, err := fmt.Scan(&ReadCommand)
	if err != nil {
		fmt.Println("Erro ao ler o comando:", err)
		return -1
	}
	return ReadCommand
}
