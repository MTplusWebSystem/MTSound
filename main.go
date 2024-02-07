package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	pathAbsolute,err := filepath.Abs("~/Musicas")

	if err != nil{
		fmt.Println("Erro ao converter caminho absoluto: ", err)
		return 
	}
	

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

func isMusicFile(fileName string) bool{
	extensions := []string{".mp3", ".mp4", ".m4a"}
 	for _, ext := range extensions {
  	if strings.HasSuffix(fileName, ext) {
   		return true
  	}
 }
 return false
}