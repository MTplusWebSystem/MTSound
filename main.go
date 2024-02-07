package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var musicPaths []string

func main() {
	expandedPath := filepath.Join(os.Getenv("HOME"), "Musicas")

	ListAllSongs(expandedPath)

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

func isMusicFile(fileName string) bool {
	extensions := []string{".mp3", ".mp4", ".m4a"}
	for _, ext := range extensions {
		if strings.HasSuffix(fileName, ext) {
			return true
		}
	}
	return false
}

func ListAllSongs(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Erro ao listar os arquivos:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() && isMusicFile(file.Name()) {
			musicPaths = append(musicPaths, filepath.Join(path, file.Name()))
		}
	}

	if len(musicPaths) == 0 {
		fmt.Println("Nenhuma música encontrada no diretório.")
		return
	}
}
