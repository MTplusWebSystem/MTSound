package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var musicPaths   []string

func main() {
	pathAbsolute,err := filepath.Abs("~/Musicas")

	if err != nil{
		fmt.Println("Erro ao converter caminho absoluto: ", err)
		return 
	}
	
	ListAllSongs(pathAbsolute)

	fmt.Println("Músicas disponíveis:")
 	for i, path := range musicPaths {
  		fmt.Printf("%d. %s\n", i+1, filepath.Base(path))
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


func ListAllSongs(path string) {
	expandedPath := filepath.Join(os.Getenv("HOME"), "Musicas")

	fmt.Println(expandedPath)
	files, err := os.ReadDir(expandedPath)
	if err != nil {
		fmt.Println("Erro ao listar os arquivos:", err)
		return
	}

	var musicPaths []string

	for _, file := range files {
		if !file.IsDir() && isMusicFile(file.Name()) {
			musicPaths = append(musicPaths, filepath.Join(expandedPath, file.Name()))
		}
	}

	if len(musicPaths) == 0 {
		fmt.Println("Nenhuma música encontrada no diretório.")
		return
	}

}