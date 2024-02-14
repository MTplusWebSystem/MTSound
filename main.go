package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"github.com/MTplusWebSystem/MTSound/variables"
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
    case "-help":
        fmt.Println(variables.Help)
    case "start":
        fmt.Println("Comando de Iniciar")
        cmd := exec.Command("vlc", "--intf", "dummy", musicMap[0])
        cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
        cmd.Start()
    case "stop":
        fmt.Println("Desligar")
        cmd := exec.Command("pkill", "vlc")
        cmd.Run()
    case "pause":
        fmt.Println("Comando de Pausar")
        cmd := exec.Command("pkill", "-STOP", "vlc")
        cmd.Run()
    case "touch":
        fmt.Println("Comando de Continuar")
        cmd := exec.Command("pkill", "-CONT", "vlc")
        cmd.Run()
    case "next-":
		kill := exec.Command("pkill", "vlc")
        kill.Run()
        fmt.Println("Comando de voltar para música anterior")
		if variables.Indicator == 0 { 
			variables.Indicator = len(musicMap) - 1
		}
		cmd := exec.Command("vlc", "--intf", "dummy", musicMap[0])
        cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
        cmd.Start()
    case "next+":
		kill := exec.Command("pkill", "vlc")
        kill.Run()
        fmt.Println("Comando de ir para próxima música")
		if variables.Indicator == len(musicMap) - 1 { 
			variables.Indicator = 0
		}else{
			variables.Indicator++
		}
		cmd := exec.Command("vlc", "--intf", "dummy", musicMap[0])
        cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
        cmd.Start()
    default:
        fmt.Println("Comando inválido! Use 'start', 'stop', 'pause', 'next+' ou 'next-'")
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
