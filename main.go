package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
)

var (
	musicIndex   int
	musicMap     = make(map[int]string)
	musicMapLock sync.Mutex
	playerCmd    *exec.Cmd
	playerPaused bool
)

func main() {
	expandedPath := filepath.Join(os.Getenv("HOME"), "Musicas")

	ListAllSongs(expandedPath)

	fmt.Println("Músicas disponíveis:")
	for i, path := range musicMap {
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
			Start(0)
		} else {
			Stop()
		}
	case "stop":
		fmt.Println("Desligar")
		Shutdown()
	case "break":
		fmt.Println("Comando de Pausar")
		Pause()
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
			musicMap[len(musicMap)] = filepath.Join(path, file.Name())
		}
	}

	if len(musicMap) == 0 {
		fmt.Println("Nenhuma música encontrada no diretório.")
		return
	}
}

func Start(index int) {
	Stop()
	musicMapLock.Lock()
	defer musicMapLock.Unlock()
	fmt.Println("Inicia a reprodução da música atual:", filepath.Base(musicMap[index]))
	playerCmd = exec.Command("ffplay", "-nodisp", "-autoexit", musicMap[index])
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

	err := playerCmd.Process.Signal(syscall.SIGTERM)
	if err != nil {
		fmt.Println("Erro ao parar a reprodução:", err)
		return
	}

	playerPaused = false
}

func Pause() {
	if playerCmd == nil || playerCmd.Process == nil || playerPaused {
		fmt.Println("Nenhuma música reproduzindo ou já pausada.")
		return
	}

	err := playerCmd.Process.Signal(syscall.SIGSTOP)
	if err != nil {
		fmt.Println("Erro ao pausar a reprodução:", err)
		return
	}

	playerPaused = true
	fmt.Println("Reprodução pausada.")
}

func Shutdown() {
	fmt.Println("Encerrando a reprodução.")
	cmd := exec.Command("pkill", "ffplay")
	cmd.Run()
}

func Next() {
	if playerCmd != nil && playerCmd.Process != nil {
		Stop()
	}
	musicMapLock.Lock()
	defer musicMapLock.Unlock()
	Start((getCurrentIndex() + 1) % len(musicMap))
}

func Back() {
	if playerCmd != nil && playerCmd.Process != nil {
		Stop()
	}
	musicMapLock.Lock()
	defer musicMapLock.Unlock()
	index := getCurrentIndex() - 1
	if index < 0 {
		index = len(musicMap) - 1
	}
	Start(index)
}

func getCurrentIndex() int {
	musicMapLock.Lock()
	defer musicMapLock.Unlock()
	return musicIndex
}
