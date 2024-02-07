package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Uso: mtsound <comando>")
		return
	}

	switch args[1] {
	case "start":
		fmt.Println("Comando de Iniciar")
	case "break":
		fmt.Println("Comando de Pausar")
	case "next+":
		fmt.Println("Comando de ir para próxima música")
	case "next-":
		fmt.Println("Comando de voltar para música anterior")
	default:
		fmt.Println("Comando inválido! Use 'start', 'break', 'next+' ou 'next-'")
	}
}
