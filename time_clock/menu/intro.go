package menu

import (
	"fmt"
	"os"
	"os/exec"
)

func ShowIntro() {
	var nome string
	versao := "0.1"
	fmt.Println("Olá, qual seu nome?")
	fmt.Scan(&nome)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	println("Olá", nome)
	println("Este programa está na versão", versao)
}
