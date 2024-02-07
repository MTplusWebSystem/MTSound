package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

var (
	musicIndex   int
	musicPaths   []string
	playerCmd    *exec.Cmd
	playerPaused bool
)

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
		if playerCmd == nil || playerCmd.Process == nil {
			Start(musicIndex)
		} else {
			Stop()
		}
	case "stop":
		fmt.Println("Desligar")
		Shutdown()
	case "break":
		fmt.Println("Comando de Pausar")
		Stop()
	case "next-":
		fmt.Println("Comando de voltar para música anterior")
		Back()
	case "next+":
		fmt.Println("Comando de ir para próxima música")
		Next()
	default:
		fmt.Println("Comando inválido! Use 'start', 'stop', 'break', 'next+' ou 'next-'")
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

func Start(index int) {
	Stop()
	fmt.Println("Inicia a reprodução da música atual.")
	playerCmd = exec.Command("ffplay", "-nodisp", "-autoexit", musicPaths[index])
	playerCmd.Stdout = nil
	playerCmd.Stderr = nil

	if err := playerCmd.Start(); err != nil {
		fmt.Println("Erro ao iniciar a reprodução:", err)
		return
	}
}

func Stop() {
	if playerCmd == nil || playerCmd.Process == nil {
		fmt.Println("Nenhuma música em reprodução.")
		return
	}

	err := playerCmd.Process.Signal(syscall.SIGINT)
	if err != nil {
		fmt.Println("Erro ao pausar a reprodução:", err)
		return
	}

	playerPaused = true
	fmt.Println("Reprodução pausada.")
}

func Play() {
	if playerCmd == nil || playerCmd.Process == nil || !playerPaused {
		fmt.Println("Nenhuma música pausada.")
		return
	}
	playerCmd.Process.Signal(syscall.SIGCONT)
	fmt.Println("Reprodução retomada.")
	playerPaused = false
}

func Shutdown() {
	fmt.Println("Encerrando a reprodução.")
	cmd := exec.Command("pkill", "ffplay")
	cmd.Run()
}

func Next() {
    fmt.Println("Muda para a próxima música da reprodução.")
    if playerPaused {
        Play()
    }
    if musicIndex < len(musicPaths)-1 {
        musicIndex++
    } else {
        musicIndex = 0
    }
    Start(musicIndex)
}

func Back() {
    fmt.Println("Volta a reprodução da música anterior.")
    if playerPaused {
        Play()
    }
    if musicIndex > 0 {
        musicIndex--
    } else {
        musicIndex = len(musicPaths) - 1
    }
    Start(musicIndex)
}

